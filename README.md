# X1 Network

X1 is a simple, fast, and secure EVM-compatible network for the next generation of decentralized applications powered by Polygon Edge.

## Explore the Network

- [X1 Explorer](https://explorer.x1-devnet.xen.network/)

## RPC Endpoints

- https://x1-devnet.xen.network/
- wss://x1-devnet.xen.network/ws

## Run Full Node

```shell
# Download and extract the latest release from https://github.com/FairCrypto/polygon-edge/releases
wget https://github.com/FairCrypto/polygon-edge/releases/download/v0.6.3-gas2/polygon-edge_0.6.3-gas2_linux_amd64.tar.gz
tar xvf polygon-edge_0.6.3-gas2_linux_amd64.tar.gz

# Initialize data folders for IBFT and generate validator keys
polygon-edge secrets init --data-dir data

# Download genesis file
wget https://x1-devnet.s3.us-west-2.amazonaws.com/genesis.jso

# Start the node
./polygon-edge server --data-dir ./data --chain genesis.json --jsonrpc 0.0.0.0:8545
```
