package blockchain

import (
	"fmt"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/0glabs/0g-monitor/utils"
	"github.com/Conflux-Chain/go-conflux-util/metrics"
	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/sirupsen/logrus"
)

var (
	blockTxCntRecord       map[uint64]int
	blockFailedTxCntRecord map[uint64]int
	blockFailedTxCntLock   sync.RWMutex
)

func MustMonitorFromViper() {
	var config Config
	viper.MustUnmarshalKey("blockchain", &config)
	Monitor(config)
}

func Monitor(config Config) {
	logrus.WithFields(logrus.Fields{
		"nodes":      len(config.Nodes),
		"validators": len(config.Validators),
	}).Info("Start to monitor blockchain")

	createMetricsForChain()

	// Connect to all fullnodes
	var nodes []*Node
	for name, url := range config.Nodes {
		logrus.WithField("name", name).WithField("url", url).Debug("Start to monitor fullnode")
		nodes = append(nodes, MustNewNode(name, url))
	}

	var validators []*Validator
	url, _ := url.Parse(config.CosmosRest)
	for name, address := range config.Validators {
		logrus.WithField("name", name).WithField("address", address).Debug("Start to monitor validator")
		validators = append(validators, MustNewValidator(url, name, address))
	}

	consensus := MustNewConsensus(config.CometbftRPC)

	blockTxCntRecord = make(map[uint64]int, config.BlockTxCntLimit)
	blockFailedTxCntRecord = make(map[uint64]int, config.BlockTxCntLimit)

	// Monitor once immediately
	monitorAllOnce(&config, nodes, validators, consensus)

	// Monitor node status periodically
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	monitorValidatorCnt := 0
	monitorNodeCnt := 0
	monitorMempoolCnt := 0
	for range ticker.C {
		monitorValidatorCnt++
		monitorNodeCnt++
		monitorMempoolCnt++

		if monitorNodeCnt%config.NodeInterval == 0 {
			monitorNodeOnce(&config, nodes, consensus)
			monitorNodeCnt = 0
		}

		if monitorValidatorCnt%config.ValidatorInterval == 0 {
			monitorValidatorOnce(&config, validators)
			monitorValidatorCnt = 0
		}

		if monitorMempoolCnt%config.MempoolInterval == 0 {
			monitorMempoolOnce(&config, consensus)
			monitorMempoolCnt = 0
		}
	}
}

func createMetricsForChain() {
	metrics.GetOrRegisterGauge(validatorActiveCountPattern).Update(0)
	metrics.GetOrRegisterGauge(validatorActiveCountUnhealthPattern).Update(0)

	metrics.GetOrRegisterGauge(failedTxCountUnhealthPattern).Update(0)
	metrics.GetOrRegisterGauge(failedTxCountPattern).Update(0)

	metrics.GetOrRegisterGauge(blockTxCountPattern).Update(0)
}

func monitorAllOnce(config *Config, nodes []*Node, validators []*Validator, consensus *Consensus) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		if logrus.IsLevelEnabled(logrus.DebugLevel) {
			logrus.WithFields(logrus.Fields{
				"costed": utils.PrettyElapsed(elapsed),
			}).Debug("executed monitorAllOnce:")
		}
	}()

	var allTasks sync.WaitGroup

	allTasks.Add(1)
	go utils.SafeStartGoroutine(func() {
		defer allTasks.Done()
		monitorNodeOnce(config, nodes, consensus)
	})

	allTasks.Add(1)
	go utils.SafeStartGoroutine(func() {
		defer allTasks.Done()
		monitorValidatorOnce(config, validators)
	})

	allTasks.Add(1)
	go utils.SafeStartGoroutine(func() {
		defer allTasks.Done()
		monitorMempoolOnce(config, consensus)
	})

	allTasks.Wait()
}

func monitorNodeOnce(config *Config, nodes []*Node, consensus *Consensus) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		if logrus.IsLevelEnabled(logrus.DebugLevel) {
			logrus.WithFields(logrus.Fields{
				"costed": utils.PrettyElapsed(elapsed),
			}).Debug("executed monitorNodeOnce:")
		}
	}()

	blockSwitched := false
	var blockTxInfo *BlockTxInfo

	if len(nodes) > 0 {
		swg := utils.NewSizedWaitGroup(len(nodes))

		for i := range nodes {
			swg.Add()
			go func(v *Node) {
				defer swg.Done()
				now := time.Now()
				v.UpdateHeight(config.AvailabilityReport)
				println("PPP> ", v.name, " UpdateHeight-1 costed: ", time.Since(now).String())
				now = time.Now()
				// generate block tx info for new block
				blockFailedTxCntLock.RLock()
				_, existed := blockTxCntRecord[v.currentBlockInfo.Height]
				blockFailedTxCntLock.RUnlock()
				if !existed {
					blockFailedTxCntLock.Lock()
					blockTxCntRecord[v.currentBlockInfo.Height] = len(v.currentBlockInfo.TxHashes)
					blockFailedTxCntLock.Unlock()

					blockTxInfo = &BlockTxInfo{
						Height:   v.currentBlockInfo.Height,
						TxHashes: v.currentBlockInfo.TxHashes,
					}

					if !blockSwitched {
						blockSwitched = true
					}
				}
				println("PPP> ", v.name, " UpdateHeight-2 costed: ", time.Since(now).String())
			}(nodes[i])
		}
		swg.Wait()
	}
	println("PPP> monitorNodeOnce-1 costed: ", time.Since(start).String())
	now := time.Now()
	max := FindMaxBlockHeight(nodes)
	if max == 0 {
		return
	}
	defaultBlockchainHeightHealth.Update(config.BlockchainHeightReport, max)

	logrus.WithField("height", max).Debug("Fullnode status report")
	println("PPP> monitorNodeOnce-2 costed: ", time.Since(now).String())
	now = time.Now()
	for _, v := range nodes {
		v.CheckHeight(&config.NodeHeightReport, max)
	}
	println("PPP> monitorNodeOnce-3 costed: ", time.Since(now).String())

	// detect tx failures and detect fork
	if blockSwitched {
		now = time.Now()
		monitorTxFailures(config, nodes, blockTxInfo)
		println("PPP> monitorNodeOnce-4 costed: ", time.Since(now).String())
		now = time.Now()
		// detect chain fork
		recordor := make(map[uint64]string, 20)
		for _, v := range nodes {
			v.CheckFork(recordor)
		}
		println("PPP> monitorNodeOnce-5 costed: ", time.Since(now).String())
		now = time.Now()
		monitorBlockValidator(config, consensus, blockTxInfo.Height)
		println("PPP> monitorNodeOnce-6 costed: ", time.Since(now).String())
	}
}

func countFailedTx(statusMap map[string]bool) int {
	failedCnt := 0
	for _, status := range statusMap {
		if !status {
			failedCnt++
		}
	}
	return failedCnt
}

func monitorTxFailures(config *Config, nodes []*Node, txInfo *BlockTxInfo) {
	if txInfo != nil {
		now := time.Now()
		blockTxCnt := len(txInfo.TxHashes)
		metrics.GetOrRegisterGauge(blockTxCountPattern).Update(int64(blockTxCnt))

		logrus.Debug(fmt.Sprintf("Block (%d) tx count: %d", txInfo.Height, blockTxCnt))

		if blockTxCnt > 0 {
			rec := make(map[int]bool, len(nodes))
			index := int(time.Now().UnixMilli() % int64(len(nodes)))
			for i := 0; i < len(nodes); i++ {
				nodeIndex := index % len(nodes)
				if _, exists := rec[index]; !exists {
					if nodes[nodeIndex].currentBlockInfo.Height == txInfo.Height {
						rec[index] = true
						statusMap, err := nodes[nodeIndex].FetchBlockReceiptStatus(config.NodeHeightReport.TimedCounterConfig, txInfo.Height)
						if err != nil {
							logrus.WithError(err).
								WithField("node_height", nodes[nodeIndex].currentBlockInfo.Height).
								WithField("height", txInfo.Height).
								WithField("nodeIndex", nodeIndex).
								Info("Failed to fetch block receipt status")
						} else {
							blockFailedTxCntLock.Lock()
							blockFailedTxCntRecord[txInfo.Height] = countFailedTx(statusMap)
							blockFailedTxCntLock.Unlock()
							break
						}
					} else {
						logrus.WithFields(logrus.Fields{
							"height":    nodes[nodeIndex].currentBlockInfo.Height,
							"target":    txInfo.Height,
							"nodeIndex": nodeIndex,
						}).Info("Skip node because of block height not match")
					}
					i++
				}
				index++
			}
			blockFailedTxCntLock.RLock()
			_, existed := blockFailedTxCntRecord[txInfo.Height]
			blockFailedTxCntLock.RUnlock()
			if !existed {
				logrus.WithField("height", txInfo.Height).Info("Failed to fetch block receipt status for this height, set to 0")
				blockFailedTxCntLock.Lock()
				blockFailedTxCntRecord[txInfo.Height] = 0
				blockFailedTxCntLock.Unlock()
			}
		}
		println("PPP> monitorNodeOnce-4-1 costed: ", time.Since(now).String())
		now = time.Now()

		totalTxCnt, failedTxCnt := 0, 0
		for i := 0; i < config.BlockTxCntLimit; i++ {
			if uint64(i) > txInfo.Height {
				break
			}
			targetHeight := txInfo.Height - uint64(i)
			blockFailedTxCntLock.RLock()
			cnt, existed := blockTxCntRecord[targetHeight]
			failedCnt, ok := blockFailedTxCntRecord[targetHeight]
			blockFailedTxCntLock.RUnlock()
			if existed {
				totalTxCnt += cnt
				if ok {
					failedTxCnt += failedCnt
				}
			} else {
				break
			}
		}

		metrics.GetOrRegisterGauge(failedTxCountPattern).Update(int64(failedTxCnt))
		percentage := float64(failedTxCnt*100) / float64(totalTxCnt)
		if failedTxCnt > 0 && percentage-float64(config.FailedTxCntAlarmThreshold) > 0 {
			metrics.GetOrRegisterGauge(failedTxCountUnhealthPattern).Update(1)
		} else {
			metrics.GetOrRegisterGauge(failedTxCountUnhealthPattern).Update(0)
		}
		println("PPP> monitorNodeOnce-4-2 costed: ", time.Since(now).String())
		now = time.Now()

		blockFailedTxCntLock.RLock()
		recordCnt := len(blockTxCntRecord)
		blockFailedTxCntLock.RUnlock()

		if recordCnt > config.BlockTxCntLimit {
			if uint64(config.BlockTxCntLimit) <= nodes[0].currentBlockInfo.Height {
				startHeight := nodes[0].currentBlockInfo.Height - uint64(config.BlockTxCntLimit)

				blockFailedTxCntLock.Lock()
				for k := range blockTxCntRecord {
					if k < startHeight {
						delete(blockTxCntRecord, k)
						delete(blockFailedTxCntRecord, k)
					}
				}
				blockFailedTxCntLock.Unlock()
			}
		}

		println("PPP> monitorNodeOnce-4-3 costed: ", time.Since(now).String())
	} else {
		metrics.GetOrRegisterGauge(blockTxCountPattern).Update(0)
	}
}

func monitorValidatorOnce(config *Config, validators []*Validator) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		if logrus.IsLevelEnabled(logrus.DebugLevel) {
			logrus.WithFields(logrus.Fields{
				"costed": utils.PrettyElapsed(elapsed),
			}).Debug("executed monitorValidatorOnce:")
		}
	}()

	jailedCnt := int32(0)

	if len(validators) > 0 {
		swg := utils.NewSizedWaitGroup(len(validators))

		for i := range validators {
			swg.Add()
			go func(v *Validator) {
				defer swg.Done()
				v.Update(config.ValidatorReport)
				if v.jailed {
					atomic.AddInt32(&jailedCnt, 1)
				}
			}(validators[i])
		}
		swg.Wait()
	}

	println("PPP> monitorValidatorOnce-1 costed: ", time.Since(start).String())

	activeValidatorCount := len(validators) - int(jailedCnt)
	metrics.GetOrRegisterGauge(validatorActiveCountPattern).Update(int64(activeValidatorCount))
	percentage := 100 * float64(activeValidatorCount) / float64(len(validators))
	if percentage-float64(67) >= 0 {
		metrics.GetOrRegisterGauge(validatorActiveCountUnhealthPattern).Update(0)
	} else {
		metrics.GetOrRegisterGauge(validatorActiveCountUnhealthPattern).Update(1)
	}

	logrus.WithField("active", activeValidatorCount).WithField("jailed", jailedCnt).Debug("Validators status report")
}

func monitorMempoolOnce(config *Config, consensus *Consensus) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		if logrus.IsLevelEnabled(logrus.DebugLevel) {
			logrus.WithFields(logrus.Fields{
				"costed": utils.PrettyElapsed(elapsed),
			}).Debug("executed monitorMempoolOnce:")
		}
	}()

	unconfirmedTxCnt := consensus.UpdateUncommitTxCnt(config.MempoolReport.TimedCounterConfig)

	metrics.GetOrRegisterGauge(mempoolUncommitTxCntPattern).Update(int64(unconfirmedTxCnt))
	percentage := float64(unconfirmedTxCnt*100) / float64(config.MempoolReport.PoolSize)
	metrics.GetOrRegisterGauge(mempoolLoadPattern).Update(int64(percentage))
	logrus.Debug("Mempool status report: unconfirmed tx count = ", unconfirmedTxCnt, ", percentage = ", percentage)
	if percentage-float64(config.MempoolReport.AlarmThreshold) > 0 {
		metrics.GetOrRegisterGauge(mempoolHighLoadPattern).Update(1)
	} else {
		metrics.GetOrRegisterGauge(mempoolHighLoadPattern).Update(0)
	}
}

func monitorBlockValidator(config *Config, consensus *Consensus, blockHeight uint64) {
	blkValidatorCnt := consensus.GetBlockValidatorCnt(config.MempoolReport.TimedCounterConfig, blockHeight)
	logrus.Debug(fmt.Sprintf("count of validator who signed block %d = %d", blockHeight, blkValidatorCnt))
	metrics.GetOrRegisterGauge(blockValidatorCountPattern).Update(int64(blkValidatorCnt))
}

func FindMaxBlockHeight(nodes []*Node) uint64 {
	max := uint64(0)

	for _, v := range nodes {
		if v.rpcHealth.IsSuccess() && max < v.currentBlockInfo.Height {
			max = v.currentBlockInfo.Height
		}
	}

	return max
}
