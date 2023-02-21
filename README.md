# X1 by Xen Crypto

X1 is a simple, fast, and secure EVM-compatible network for the next generation of decentralized applications powered by Polygon Edge.

## Run Full Node

```shell
# Download and extract the latest release from https://github.com/FairCrypto/polygon-edge/releases
wget https://github.com/FairCrypto/polygon-edge/releases/download/v0.6.3-gas2/polygon-edge_0.6.3-gas2_linux_amd64.tar.gz
tar xvf polygon-edge_0.6.3-gas2_linux_amd64.tar.gz

# Initialize data folders for IBFT and generate validator keys
polygon-edge secrets init --data-dir data

# Download genesis file
wget https://x1-devnet.s3.us-west-2.amazonaws.com/genesis.json

# Start the node
./polygon-edge server --data-dir ./data --chain genesis.json --grpc-address :10000 --libp2p :10001 --jsonrpc :10002
```
