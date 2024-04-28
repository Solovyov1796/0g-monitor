# 0G Monitoring Service

This is to monitor all 0G services, including blockchain fullnodes, storage nodes and some other infrastructures.

- Blockchain
    - Fullnode RPC availability via `eth_blockNumber`.
    - Whether block height fall behind, e.g. 30 blocks behind the highest block height.

## Configuration

There are two files to configure monitoring parameters:

- `config.yaml`: change any default configurations and uncomment out.
- `.env`: recommended, overwrite configurations via environment variables prefixed with `ZG_MONITOR`, for example:

```shell
# Enable alert configurations (dingtalk)
export ZG_MONITOR_ALERT_DINGTALK_ENABLED=true
export ZG_MONITOR_ALERT_CUSTOMTAGS="testnet"
export ZG_MONITOR_ALERT_DINGTALK_WEBHOOK="${dingtalk_webhook_url}"
export ZG_MONITOR_ALERT_DINGTALK_SECRET="${dingtalk_secret}"

# Configure blockchain monitoring variables
export ZG_MONITOR_BLOCKCHAIN_NODES_NODE0="http://ip0:8545"
export ZG_MONITOR_BLOCKCHAIN_NODES_NODE1="http://ip1:8545"
export ZG_MONITOR_BLOCKCHAIN_NODES_NODE2="http://ip2:8545"
```

## Run

Before monitoring service launched, do not forget to load the `.env` file if any.

```shell
# Load configurations on prod machine.
source .env

# Start monitoring service.
./0g-monitor
```
