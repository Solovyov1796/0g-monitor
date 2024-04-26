package blockchain

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Nodes              map[string]string
	Interval           time.Duration `default:"5s"`
	AvailabilityReport ErrorTolerantReportConfig
	HeightReport       HeightReportConfig
}

func MustMonitorFromViper() {
	var config Config
	viper.MustUnmarshalKey("blockchain", &config)
	Monitor(config)
}

func Monitor(config Config) {
	logrus.WithField("nodes", len(config.Nodes)).Info("Start to monitor blockchain")

	if len(config.Nodes) == 0 {
		return
	}

	// Connect to all fullnodes
	var nodes []*Node
	for name, url := range config.Nodes {
		logrus.WithField("name", name).WithField("url", url).Debug("Start to monitor fullnode")
		nodes = append(nodes, MustNewNode(name, url))
	}
	defer func() {
		for _, v := range nodes {
			defer v.Close()
		}
	}()

	// Monitor node status periodically
	ticker := time.NewTicker(config.Interval)
	defer ticker.Stop()

	for range ticker.C {
		monitorOnce(&config, nodes)
	}
}

func monitorOnce(config *Config, nodes []*Node) {
	for _, v := range nodes {
		v.UpdateHeight(&config.AvailabilityReport)
	}

	max := FindMaxBlockHeight(nodes)
	if max == 0 {
		return
	}

	logrus.WithField("height", max).Debug("Fullnode status report")

	for _, v := range nodes {
		v.CheckHeight(&config.HeightReport, max)
	}
}
