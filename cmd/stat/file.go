package stat

import (
	"context"
	"fmt"

	"github.com/0glabs/0g-storage-client/common/shard"
	"github.com/0glabs/0g-storage-client/node"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var (
	txSeq uint64
	root  string

	fileCmd = &cobra.Command{
		Use:   "file",
		Short: "Statistic file distribution of storage node network",
		Run:   statFileDistribution,
	}
)

func init() {
	fileCmd.Flags().Uint64Var(&txSeq, "seq", 0, "Tx seq to stat")
	fileCmd.Flags().StringVar(&root, "root", "", "File root to stat")
	fileCmd.MarkFlagsOneRequired("seq", "root")
	fileCmd.MarkFlagsMutuallyExclusive("seq", "root")

	Cmd.AddCommand(fileCmd)
}

func statFileDistribution(*cobra.Command, []string) {
	type ShardedFileInfo struct {
		File  *node.FileInfo
		Shard shard.ShardConfig
	}

	files := mustStatRpc(func(client *node.ZgsClient, ctx context.Context) (*ShardedFileInfo, error) {
		var info ShardedFileInfo
		var err error

		if info.Shard, err = client.GetShardConfig(ctx); err != nil {
			return nil, err
		}

		if len(root) == 0 {
			info.File, err = client.GetFileInfoByTxSeq(ctx, txSeq)
		} else {
			info.File, err = client.GetFileInfo(ctx, common.HexToHash(root))
		}

		return &info, err
	})

	var rpcFailures int
	fileDistribution := make(map[string]int)
	uploadedDistribution := make(map[uint64]map[uint64]int)

	for _, v := range files {
		if v.Err != nil {
			rpcFailures++
		} else if v.Data.File == nil {
			fileDistribution["Unsynced"]++
		} else if v.Data.File.Finalized {
			fileDistribution["Uploaded"]++

			if m, ok := uploadedDistribution[v.Data.Shard.NumShard]; ok {
				m[v.Data.Shard.ShardId]++
			} else {
				uploadedDistribution[v.Data.Shard.NumShard] = map[uint64]int{
					v.Data.Shard.ShardId: 1,
				}
			}
		} else {
			fileDistribution["Synced"]++
		}
	}

	fmt.Println("\nRPC Failures:", rpcFailures)

	fmt.Println("\nStatus distribution:")
	for status, count := range fileDistribution {
		fmt.Printf("\t%v: %v\n", status, count)
	}

	fmt.Println("\nUploaded distribution:")
	for numShard, id2Counts := range uploadedDistribution {
		fmt.Println("\tNum shard:", numShard)
		for shardId, count := range id2Counts {
			fmt.Printf("\t\tShard %v: %v\n", shardId, count)
		}
	}
}
