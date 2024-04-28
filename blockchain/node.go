package blockchain

import (
	"github.com/openweb3/web3go"
	"github.com/sirupsen/logrus"
)

type HeightReportConfig struct {
	ErrorTolerantReportConfig

	MaxGap uint64 `default:"30"`
}

type Node struct {
	*web3go.Client

	name string
	url  string

	height uint64 // fullnode block number

	rpcHealth Health
	rpcError  string // last rpc error message

	heightHealth Health
}

func MustNewNode(name, url string) *Node {
	return &Node{
		Client: web3go.MustNewClient(url),
		name:   name,
		url:    url,
	}
}

func (node *Node) UpdateHeight(config *ErrorTolerantReportConfig) {
	bn, err := node.Eth.BlockNumber()
	if err != nil {
		logrus.WithError(err).WithField("node", node.name).Debug("Failed to query block number")

		node.rpcError = err.Error()
		unhealthy, unrecovered, elapsed := node.rpcHealth.OnFailure(config)

		// report unhealthy
		if unhealthy {
			logrus.WithFields(logrus.Fields{
				"node":    node.name,
				"elapsed": elapsed,
				"error":   node.rpcError,
			}).Error("Node RPC became unhealthy")
		}

		// remind unhealthy
		if unrecovered {
			logrus.WithFields(logrus.Fields{
				"node":    node.name,
				"elapsed": elapsed,
				"error":   node.rpcError,
			}).Error("Node RPC not recovered yet")
		}
	} else {
		// check reorg
		if bn.Uint64() < node.height {
			logrus.WithFields(logrus.Fields{
				"node":     node.name,
				"old":      node.height,
				"new":      bn.Uint64(),
				"reverted": node.height - bn.Uint64(),
			}).Warn("Block reorg detected")
		}

		node.height = bn.Uint64()
		node.rpcError = ""

		// report on recovered
		if recovered, elapsed := node.rpcHealth.OnSuccess(config); recovered {
			logrus.WithFields(logrus.Fields{
				"node":    node.name,
				"elapsed": elapsed,
			}).Warn("Node RPC is healthy now")
		}
	}
}

func FindMaxBlockHeight(nodes []*Node) uint64 {
	max := uint64(0)

	for _, v := range nodes {
		if !v.rpcHealth.HasError() && max < v.height {
			max = v.height
		}
	}

	return max
}

func (node *Node) CheckHeight(config *HeightReportConfig, target uint64) {
	// ignore on rpc error
	if !node.rpcHealth.HasError() {
		return
	}

	// number of blocks fall behind
	var behind uint64
	if node.height < target {
		behind = target - node.height
	}

	if behind <= config.MaxGap {
		if recovered, elapsed := node.heightHealth.OnSuccess(&config.ErrorTolerantReportConfig); recovered {
			logrus.WithFields(logrus.Fields{
				"node":    node.name,
				"elapsed": elapsed,
				"behind":  behind,
			}).Warn("Node block height is healthy now")
		}
	} else {
		unhealthy, unrecovered, elapsed := node.heightHealth.OnFailure(&config.ErrorTolerantReportConfig)

		if unhealthy {
			logrus.WithFields(logrus.Fields{
				"node":    node.name,
				"elapsed": elapsed,
				"behind":  behind,
			}).Error("Node block height became unhealthy")
		}

		if unrecovered {
			logrus.WithFields(logrus.Fields{
				"node":    node.name,
				"elapsed": elapsed,
				"behind":  behind,
			}).Error("Node block height not recovered yet")
		}
	}
}
