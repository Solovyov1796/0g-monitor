package blockchain

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/Conflux-Chain/go-conflux-util/health"
	"github.com/Conflux-Chain/go-conflux-util/metrics"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Validator struct {
	url     string
	name    string
	address string
	health  health.TimedCounter
}

func MustNewValidator(url *url.URL, name, address string) *Validator {
	validator, err := NewValidator(url, name, address)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"name":    name,
			"address": address,
		}).WithError(err).Info("Failed to create validator node")
		return nil
	}

	return validator
}

func NewValidator(url *url.URL, name, address string) (*Validator, error) {
	address = strings.TrimSpace(address)
	if len(address) == 0 {
		return nil, errors.New("empty address")
	}

	url.Path = "/cosmos/staking/v1beta1/validators/" + address

	metrics.GetOrRegisterGauge("monitor/blockchain/validator/jailed/%v", name).Update(0)

	return &Validator{
		url:     url.String(),
		name:    name,
		address: address,
	}, nil
}

func (validator Validator) String() string {
	if len(validator.name) == 0 {
		return validator.address
	}

	return validator.name
}

func (validator *Validator) Update(config health.TimedCounterConfig) {
	client := resty.New()
	var result map[string]interface{}
	resp, err := client.R().SetResult(&result).Get(validator.url)
	if err != nil || resp.StatusCode() != 200 {
		logrus.WithError(err).WithField("validator", validator.String()).Info("Failed to query validator info")
		return
	}
	jailed := result["validator"].(map[string]interface{})["jailed"].(bool)

	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		jsonStr, _ := json.Marshal(result)
		logrus.WithFields(logrus.Fields{
			"validator": validator.String(),
			"info":      fmt.Sprintf("%+v", jsonStr),
		}).Debug("Validator status report")
	}

	if jailed {
		unhealthy, unrecovered, elapsed := validator.health.OnFailure(config)

		if unhealthy {
			logrus.WithFields(logrus.Fields{
				"elapsed":   prettyElapsed(elapsed),
				"validator": validator.String(),
			}).Error("Validator jailed")

			metrics.GetOrRegisterGauge("monitor/blockchain/validator/jailed/%v", validator.name).Update(1)
		}

		if unrecovered {
			logrus.WithFields(logrus.Fields{
				"elapsed":   prettyElapsed(elapsed),
				"validator": validator.String(),
			}).Error("Validator jailed and not recovered yet")
		}
	} else if recovered, elapsed := validator.health.OnSuccess(config); recovered {
		logrus.WithFields(logrus.Fields{
			"elapsed":   prettyElapsed(elapsed),
			"validator": validator.String(),
		}).Warn("Validator unfailed now")

		metrics.GetOrRegisterGauge("monitor/blockchain/validator/jailed/%v", validator.name).Update(0)
	}
}
