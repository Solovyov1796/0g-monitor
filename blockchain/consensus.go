package blockchain

import (
	"net/url"
	"time"

	"github.com/0glabs/0g-monitor/utils"
	"github.com/Conflux-Chain/go-conflux-util/health"
	"github.com/Conflux-Chain/go-conflux-util/metrics"
	"github.com/sirupsen/logrus"
)

type Consensus struct {
	url string

	cometbftRpcHealth health.TimedCounter
	cometbftRpcError  string // last rpc error message

	commonCfg health.TimedCounterConfig
}

func MustNewConsensus(urlstr string, commonCfg health.TimedCounterConfig) *Consensus {
	url, _ := url.Parse(urlstr)

	metrics.GetOrRegisterGauge(mempoolUncommitTxCntPattern).Update(0)
	metrics.GetOrRegisterGauge(mempoolHighLoadPattern).Update(0)

	metrics.GetOrRegisterHistogram(nodeCometbftRpcLatencyPattern, "consensus").Update(0)
	metrics.GetOrRegisterGauge(nodeCometbftRpcUnhealthPattern, "consensus").Update(0)

	return &Consensus{
		url:       url.String(),
		commonCfg: commonCfg,
	}
}

func (m *Consensus) UpdateUncommitTxCnt() int {
	var unconfirmedTxCnt int
	executeRequest(
		func() error {
			var err error
			unconfirmedTxCnt, err = rpcGetUncommitTxCnt(m.url)
			if err != nil {
				return err
			} else {
				return nil
			}
		},
		func(err error, unhealthy, unrecovered bool, elapsed time.Duration) {
			m.cometbftRpcError = err.Error()
			// report unhealthy
			if unhealthy {
				logrus.WithFields(logrus.Fields{
					"node":    "consensus",
					"elapsed": utils.PrettyElapsed(elapsed),
					"error":   err,
				}).Error("Node cometbft RPC became unhealthy")
			}

			// remind unhealthy
			if unrecovered {
				logrus.WithFields(logrus.Fields{
					"node":    "consensus",
					"elapsed": utils.PrettyElapsed(elapsed),
				}).Error("Node cometbft RPC not recovered yet")
			}
		},
		func(recovered bool, elapsed time.Duration) {
			m.cometbftRpcError = ""
			if recovered {
				logrus.WithFields(logrus.Fields{
					"node":    "consensus",
					"elapsed": utils.PrettyElapsed(elapsed),
				}).Warn("Node cometbft RPC is healthy now")
			}
		},
		nodeCometbftRpcLatencyPattern, nodeCometbftRpcUnhealthPattern, "consensus",
		&m.cometbftRpcHealth,
		m.commonCfg,
	)

	return unconfirmedTxCnt
}

func (m *Consensus) GetBlockValidatorCnt(height uint64) int {
	var validatorCnt int
	executeRequest(
		func() error {
			var err error
			validatorCnt, err = rpcGetBlockValidatorCnt(m.url, height)
			if err != nil {
				return err
			} else {
				return nil
			}
		},
		func(err error, unhealthy, unrecovered bool, elapsed time.Duration) {
			m.cometbftRpcError = err.Error()
			// report unhealthy
			if unhealthy {
				logrus.WithFields(logrus.Fields{
					"node":    "consensus",
					"elapsed": utils.PrettyElapsed(elapsed),
					"error":   err,
				}).Error("Node cometbft RPC became unhealthy")
			}

			// remind unhealthy
			if unrecovered {
				logrus.WithFields(logrus.Fields{
					"node":    "consensus",
					"elapsed": utils.PrettyElapsed(elapsed),
				}).Error("Node cometbft RPC not recovered yet")
			}
		},
		func(recovered bool, elapsed time.Duration) {
			m.cometbftRpcError = ""
			if recovered {
				logrus.WithFields(logrus.Fields{
					"node":    "consensus",
					"elapsed": utils.PrettyElapsed(elapsed),
				}).Warn("Node cometbft RPC is healthy now")
			}
		},
		nodeCometbftRpcLatencyPattern, nodeCometbftRpcUnhealthPattern, "consensus",
		&m.cometbftRpcHealth,
		m.commonCfg,
	)

	return validatorCnt
}
