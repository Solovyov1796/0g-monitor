package blockchain

import (
	"github.com/0glabs/0g-monitor/utils"
	"github.com/Conflux-Chain/go-conflux-util/health"
	"github.com/Conflux-Chain/go-conflux-util/metrics"
	"github.com/sirupsen/logrus"
)

type GrowChecker struct {
	lastHeight uint64
	health     health.TimedCounter
	cfg        health.TimedCounterConfig
}

func MustNewGrowChecker(healthCfg health.TimedCounterConfig) *GrowChecker {
	metrics.GetOrRegisterGauge(chainHeightHaltPattern).Update(0)
	return &GrowChecker{
		cfg: healthCfg,
	}
}

func (hc *GrowChecker) Check(height uint64) {
	if height > hc.lastHeight {
		if recovered, elapsed := hc.health.OnSuccess(hc.cfg); recovered {
			logrus.WithFields(logrus.Fields{
				"elapsed": utils.PrettyElapsed(elapsed),
				"old":     hc.lastHeight,
				"new":     height,
			}).Warn("Blockchain height is growing again")

			metrics.GetOrRegisterGauge(chainHeightHaltPattern).Update(0)
		}

		hc.lastHeight = height
	} else {
		unhealthy, unrecovered, elapsed := hc.health.OnFailure(hc.cfg)

		newHeight := height
		if newHeight == 0 {
			newHeight = hc.lastHeight
		}

		if unhealthy {
			logrus.WithFields(logrus.Fields{
				"elapsed": utils.PrettyElapsed(elapsed),
				"height":  newHeight,
			}).Error("Blockchain height stops growing")

			metrics.GetOrRegisterGauge(chainHeightHaltPattern).Update(1)
		}

		if unrecovered {
			logrus.WithFields(logrus.Fields{
				"elapsed": utils.PrettyElapsed(elapsed),
				"height":  newHeight,
			}).Error("Blockchain height stops growing for a long time and not recovered yet")
		}
	}
}
