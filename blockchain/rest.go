package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

func IsValidatorJailed(url string) (bool, error) {
	client, err := createClient("rest", url)
	if err != nil {
		return false, err
	}
	var result map[string]interface{}
	resp, err := client.R().SetResult(&result).Get(url)
	if err != nil {
		return false, err
	}

	if resp.StatusCode() != 200 {
		return false, ErrorNotSuccess
	}

	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		jsonStr, _ := json.Marshal(result)
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", string(jsonStr)),
		}).Debug("response of cometbft rpc: validators?height=")
	}

	return result["validator"].(map[string]interface{})["jailed"].(bool), nil
}
