# X1 Network

X1 is a simple, fast, and secure EVM-compatible network for the next generation of decentralized applications powered by Polygon Edge.

## Explore the Network

- [X1 Explorer](https://explorer.x1-devnet.xen.network/)

## RPC Endpoints

- https://x1-devnet.xen.network/
- wss://x1-devnet.xen.network/ws

## Run Full Node

First, open port 1478 to allow for P2P traffic. Optionally, open port 8545 to allow for JSON-RPC traffic.

```shell
# Download the latest release from https://github.com/FairCrypto/x1/releases
wget https://github.com/FairCrypto/x1/releases/download/v0.6.3-gas2/polygon-edge_0.6.3-gas2_linux_amd64.tar.gz
tar xvf polygon-edge_0.6.3-gas2_linux_amd64.tar.gz

# Initialize the data folders for IBFT and generate validator keys
./polygon-edge secrets init --data-dir data

# Download the genesis file
wget https://x1-devnet.s3.us-west-2.amazonaws.com/genesis.json

# Start the node
./polygon-edge server \
    --data-dir ./data \
    --chain ./genesis.json \
    --jsonrpc 0.0.0.0:8545 \
    --libp2p 0.0.0.0:1478 \
    --nat "$(curl ifconfig.me)" \
    --price-limit 100000000000 \
    --block-time 1 \
    --log-level INFO
```

### Run a validator node

*Note: The devnet is currently a POA network and requires each validator to be voted into the network*

2/3 of the validators must vote to add your validator to the network.
Please contact Faircrypto to get your validator added to the network.

> Find your validator key
```shell
polygon-edge ibft status
```

> Vote to add your validator to the network with the following command.
```shell
polygon-edge ibft propose --addr ${VALIDATOR_KEY} --vote auth
```

