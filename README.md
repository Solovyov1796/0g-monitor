# 0G Monitoring Service

This is to monitor all 0G services, including blockchain fullnodes, storage nodes and some other infrastructures.

## Configuration

There are two files to configure monitoring parameters:

- `config.yaml`: change any configurations and uncomment out.
- `.env`: overwrite configurations via environment variables prefixed with `ZG_MONITOR`.

Generally, blockchain fullnodes should be configured, e.g. in `.env` file:

```shell
export ZG_MONITOR_LOG_LEVEL=debug

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
