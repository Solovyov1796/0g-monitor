# 0G Monitoring Service

This is to monitor all 0G services, including blockchain fullnodes, storage nodes and some other infrastructures.

- Blockchain
    - Fullnode RPC availability via `eth_blockNumber`.
    - Whether block height fall behind, e.g. 30 blocks behind the highest block height.
    - Whether blockchain height stops growing.

## Configuration

There are two files to configure monitoring parameters:

- `config.yaml`: change any default configurations and uncomment out.
- `.env`: recommended, overwrite configurations via environment variables prefixed with `ZG_MONITOR`, for example:

```shell
# Log Configurations
export ZG_MONITOR_LOG_ALERTHOOK_CHANNELS="dingtalk,telegram"


# Alert Configurations
export ZG_MONITOR_ALERT_CUSTOMTAGS="testnet"
export ZG_MONITOR_ALERT_CHANNELS_DINGTALK_PLATFORM="dingtalk"
export ZG_MONITOR_ALERT_CHANNELS_DINGTALK_WEBHOOK="${dingtalk_webhook_url}"
export ZG_MONITOR_ALERT_CHANNELS_DINGTALK_SECRET="${dingtalk_secret}"
export ZG_MONITOR_ALERT_CHANNELS_TELEGRAM_PLATFORM="telegram"
export ZG_MONITOR_ALERT_CHANNELS_TELEGRAM_APITOKEN="${telegram_api_token}"
export ZG_MONITOR_ALERT_CHANNELS_TELEGRAM_CHATID="${telegram_chat_id}"

# Blockchain Monitoring Configurations
export ZG_MONITOR_BLOCKCHAIN_NODES_NODE0="http://ip0:8545"
export ZG_MONITOR_BLOCKCHAIN_NODES_NODE1="http://ip1:8545"
export ZG_MONITOR_BLOCKCHAIN_NODES_NODE2="http://ip2:8545"

export ZG_MONITOR_BLOCKCHAIN_VALIDATORS_VAL0="kavavalcons17metfygkc7aezdrx6aucnpsh7ggh4lrhrawj66"
export ZG_MONITOR_BLOCKCHAIN_VALIDATORS_VAL1="kavavalcons17metfygkc7aezdrx6aucnpsh7ggh4lrhrawj67"
export ZG_MONITOR_BLOCKCHAIN_VALIDATORS_VAL2="kavavalcons17metfygkc7aezdrx6aucnpsh7ggh4lrhrawj68"
```

## Run

Before monitoring service launched, do not forget to load the `.env` file if any.

```shell
# Load configurations on prod machine.
source .env

# Start monitoring service.
./0g-monitor
```

or execute the `./scripts/start.sh` file to start monitoring service directly.
