package blockchain

import (
	"github.com/Conflux-Chain/go-conflux-util/health"
)

type MempoolMonitorConfig struct {
	highloadCounter health.TimedCounter

	AlarmThreshold uint64 `default:"90"`
	PoolSize       uint64 `default:"5000"`
}

type ValidatorMonitorConfig struct {
	onlinePercentageCounter health.TimedCounter

	AlarmThreshold uint64 `default:"75"`
	MaxSize        uint64 `default:"200"`
}

type Config struct {
	Nodes                     map[string]string
	MempoolInterval           int    `default:"1"`
	NodeInterval              int    `default:"3"`
	ValidatorInterval         int    `default:"15"`
	BlockBehindThreshold      uint64 `default:"10"`
	BlockGapThreshold         uint64 `default:"30"`
	Validators                map[string]string
	MempoolCfg                MempoolMonitorConfig
	ValidatorCfg              ValidatorMonitorConfig
	CosmosRPC                 string `default:"https://127.0.0.1:26657"`
	CosmosRest                string `default:"http://127.0.0.1:1317"`
	CometbftRPC               string `default:"http://127.0.0.1:26657"`
	BlockTxCntLimit           int    `default:"100"`
	FailedTxCntAlarmThreshold int    `default:"2"`
	Mode                      string `default:"localtest"`
	CommonEvtReportCfg        health.TimedCounterConfig
	CriticalEvtReportCfg      health.TimedCounterConfig
}
