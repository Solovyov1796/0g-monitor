package blockchain

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var defaultBlockchainHeightHealth BlockchainHeightHealth

type BlockchainHeightHealth struct {
	height uint64
	health Health
}

func (bhh *BlockchainHeightHealth) Update(config *ErrorTolerantReportConfig, height uint64) {
	if height > bhh.height {
		if recovered, elapsed := bhh.health.OnSuccess(config); recovered {
			logrus.WithFields(logrus.Fields{
				"elapsed": fmt.Sprint(elapsed),
				"old":     bhh.height,
				"new":     height,
			}).Warn("Blockchain height is growing again")
		}

		bhh.height = height
	} else {
		unhealthy, unrecovered, elapsed := bhh.health.OnFailure(config)

		newHeight := height
		if newHeight == 0 {
			newHeight = bhh.height
		}

		if unhealthy {
			logrus.WithFields(logrus.Fields{
				"elapsed": fmt.Sprint(elapsed),
				"height":  newHeight,
			}).Error("Blockchain height stops growing")
		}

		if unrecovered {
			logrus.WithFields(logrus.Fields{
				"elapsed": fmt.Sprint(elapsed),
				"height":  newHeight,
			}).Error("Blockchain height stops growing for a long time and not recovered yet")
		}
	}
}
