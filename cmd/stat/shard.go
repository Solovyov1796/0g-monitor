package stat

import (
	"context"
	"fmt"

	"github.com/0glabs/0g-storage-client/common/shard"
	"github.com/0glabs/0g-storage-client/node"
	"github.com/spf13/cobra"
)

var shardCmd = &cobra.Command{
	Use:   "shard",
	Short: "Statistic shard config of storage node network",
	Run:   statShard,
}

func init() {
	Cmd.AddCommand(shardCmd)
}

func statShard(*cobra.Command, []string) {
	shards := mustStatRpc(func(client *node.ZgsClient, ctx context.Context) (shard.ShardConfig, error) {
		return client.GetShardConfig(ctx)
	})

	// stat shard configs
	var rpcFailures int
	invalidShardNodes := make(map[string]shard.ShardConfig)
	shardCounter := newShardCounter()

	for node, rpcResult := range shards {
		if rpcResult.Err != nil {
			rpcFailures++
		} else if rpcResult.Data.NumShard > 1024 {
			invalidShardNodes[node] = rpcResult.Data
		} else {
			shardCounter.insert(rpcResult.Data)
		}
	}

	fmt.Println("\nRPC Failures:", rpcFailures)

	if len(invalidShardNodes) > 0 {
		fmt.Println("\nInvalid shard config nodes:")
		for node, config := range invalidShardNodes {
			fmt.Printf("\t%v: %v / %v\n", node, config.ShardId, config.NumShard)
		}
	}
	shardCounter.print("\nShard distribution")
}

type shardCounter struct {
	shard2Id2Count map[uint64]map[uint64]int
}

func newShardCounter() shardCounter {
	return shardCounter{
		shard2Id2Count: make(map[uint64]map[uint64]int),
	}
}

func (counter *shardCounter) insert(config shard.ShardConfig) {
	if id2Count, ok := counter.shard2Id2Count[config.NumShard]; ok {
		id2Count[config.ShardId]++
	} else {
		counter.shard2Id2Count[config.NumShard] = map[uint64]int{
			config.ShardId: 1,
		}
	}
}

func (counter *shardCounter) print(label string) {
	fmt.Printf("%v, replica = %v:\n", label, counter.replica())

	for numShard, id2Count := range counter.shard2Id2Count {
		fmt.Println("\tNum shard:", numShard)

		for id, count := range id2Count {
			fmt.Printf("\t\t%v: %v\n", id, count)
		}
	}
}

func (counter *shardCounter) replica() int {
	var result int

	for numShard, id2Count := range counter.shard2Id2Count {
		// any shard id missded
		if uint64(len(id2Count)) < numShard {
			continue
		}

		min := id2Count[0]

		for _, count := range id2Count {
			if min > count {
				min = count
			}
		}

		result += min
	}

	return result
}
