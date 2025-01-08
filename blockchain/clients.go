package blockchain

import (
	"net/http"
	"sync"
	"time"

	"github.com/0glabs/0g-monitor/utils"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

var clients = make(map[string]*resty.Client)
var clientLock sync.RWMutex

func createClient(prefix, url string) (*resty.Client, error) {
	key, err := utils.PeekUrlHost(url)
	if err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"url": url,
		}).Error("failed to parse url")
		return nil, err
	}

	if len(prefix) > 0 {
		key = prefix + "@" + key
	}

	clientLock.RLock()
	c, exists := clients[key]
	clientLock.RUnlock()
	if exists {
		return c, nil
	} else {
		transport := &http.Transport{
			MaxIdleConns:        3,
			MaxIdleConnsPerHost: 3,
			IdleConnTimeout:     30 * time.Second,
		}
		restClient := resty.New().SetTransport(transport).SetHeader("Connection", "keep-alive")
		clientLock.Lock()
		clients[key] = restClient
		clientLock.Unlock()
		return restClient, nil
	}
}
