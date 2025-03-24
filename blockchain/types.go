package blockchain

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrorNotSuccess = errors.New("Error: Not success")
)

var (
	EthRpcPort      = "8545"
	CosmosRpcPort   = "26657"
	CosmosRestPort  = "1317"
	CometbftRpcPort = "26657"
)

type BlockInfo struct {
	Height    uint64
	Timestamp uint64
	Hash      string
	TxHashes  []string
}

type BlockTxInfo struct {
	Height   uint64
	TxHashes []string
}

func ComposeUrl(ip, port, path string) string {
	if len(path) > 0 {
		return fmt.Sprintf("http://%s:%s/%s", ip, port, path)
	}
	return fmt.Sprintf("http://%s:%s", ip, port)
}
