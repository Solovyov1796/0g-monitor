package blockchain

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
)

func rpcGetUncommitTxCnt(url string) (int, error) {
	client, err := createClient("cometbft", url)
	if err != nil {
		return 0, err
	}
	var result map[string]interface{}
	resp, err := client.R().SetResult(&result).Get(url + "/num_unconfirmed_txs")
	if err != nil {
		logrus.WithError(err).WithField("url", url).Info("failed to get uncommitted tx count via cometbft rpc")
		return 0, err
	}
	if resp.StatusCode() != 200 {
		logrus.WithError(err).WithField("url", url).WithField("status_code", resp.StatusCode()).Info("get uncommitted tx count via cometbft rpc failed")
		return 0, fmt.Errorf("get uncommitted tx count via cometbft rpc failed, status code: %d", resp.StatusCode())
	}

	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		jsonStr, _ := json.Marshal(result)
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", string(jsonStr)),
		}).Debug("response of cometbft rpc: num_unconfirmed_txs")
	}

	cntStr := result["result"].(map[string]interface{})["total"].(string)

	cnt, err := strconv.ParseInt(cntStr, 10, 64)
	if err != nil {
		logrus.WithError(err).WithField("total", cntStr).Info("failed to convert total string to int")
		return 0, err
	}

	return int(cnt), nil
}

func rpcGetBlockValidatorCnt(url string, height uint64) (int, error) {
	client, err := createClient("cometbft", url)
	if err != nil {
		return 0, err
	}
	var result map[string]interface{}
	resp, err := client.R().SetResult(&result).Get(fmt.Sprintf("%s/validators?height=%d", url, height))
	if err != nil {
		logrus.WithError(err).WithField("url", url).WithField("height", height).Info("failed to get validator list")
		return 0, err
	}
	if resp.StatusCode() != 200 {
		logrus.WithError(err).WithField("url", url).WithField("height", height).WithField("status_code", resp.StatusCode()).Info("failed to get validator list")
		return 0, fmt.Errorf("failed to get validator list, status code: %d", resp.StatusCode())
	}

	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		jsonStr, _ := json.Marshal(result)
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", string(jsonStr)),
		}).Debug("response of cometbft rpc: validators?height=")
	}

	totalStr := result["result"].(map[string]interface{})["total"].(string)

	total, err := strconv.ParseInt(totalStr, 10, 64)
	if err != nil {
		logrus.WithError(err).WithField("total", totalStr).Info("failed to convert total string to int")
		return 0, err
	}

	return int(total), nil
}
