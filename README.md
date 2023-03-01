# X1 Network

X1 is a simple, fast, and secure EVM-compatible network for the next generation of decentralized applications powered by Polygon Edge.

## Explore the Network

- [X1 Explorer](https://explorer.x1-devnet.xen.network/)

## RPC Endpoints

- https://x1-devnet.xen.network/
- wss://x1-devnet.xen.network/ws

## Run Full Node

### Firewall

- Allow traffic on port 1478 for the P2P traffic
- Optionally, allow traffic on port 8545 for the JSON-RPC traffic

### Disk Space

100GB of disk space is required for the chain data, but the blockchain data is growing at a rate of .5GB per day.
We recommend starting with 300GB of disk space to be safe.

### Install

Quick start a full node and run in the foreground

```shell
# Install dependencies (ex: ubuntu)
sudo apt-get update
sudo apt-get install -y wget tar jq curl

# Download the latest release from https://github.com/FairCrypto/x1/releases
wget https://github.com/FairCrypto/x1/releases/download/v0.6.3-gas2/polygon-edge_0.6.3-gas2_linux_amd64.tar.gz

# Extract and move the binary to your path
tar xvf polygon-edge_0.6.3-gas2_linux_amd64.tar.gz
sudo mv polygon-edge /usr/local/bin

# Initialize the data folders for IBFT and generate validator keys
polygon-edge secrets init --data-dir data

# Download the genesis file
wget https://x1-devnet.s3.us-west-2.amazonaws.com/genesis.json

# Start the node
polygon-edge server \
    --data-dir ./data \
    --chain ./genesis.json \
    --jsonrpc 0.0.0.0:8545 \
    --libp2p 0.0.0.0:1478 \
    --nat "$(curl ifconfig.me)" \ # external IP address
    --price-limit 100000000000 \
    --block-time 1 \
    --log-level INFO
```

### Production

There's a few more steps to prepare your node for a production environment.

```shell
# Create a user to run your node
sudo adduser --system --group x1

# Move your data directory to the new user's home directory
sudo mv data genesis.json /home/x1
sudo chown -R x1:x1 /home/x1

# Create a systemd service to run your node as a daemon
sudo sh -c 'echo "[Unit]
Description=Polygon Edge Server
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=10
User=x1
LimitNOFILE=infinity
SyslogIdentifier=polygon-edge
WorkingDirectory=/home/x1
ExecStart=/usr/local/bin/polygon-edge server \
    --data-dir ./data \
    --chain ./genesis.json \
    --jsonrpc 0.0.0.0:8545 \
    --libp2p 0.0.0.0:1478 \
    --nat "$(curl ifconfig.me)" \
    --price-limit 100000000000 \
    --block-time 1 \
    --log-level INFO

[Install]
WantedBy=multi-user.target
" > /etc/systemd/system/polygon-edge.service'

# Enable and start the service. The service will now start on boot.
sudo systemctl enable --now polygon-edge

# Check the status of the service
sudo systemctl status polygon-edge

# Check the logs
journalctl -u polygon-edge -f
```


### Run a validator node

*Note: The devnet is currently a POA network and requires each validator to be voted into the network*

2/3 of the validators must vote to add your validator to the network.
Please contact Faircrypto to get your validator added to the network.

> Find your validator key

```shell
polygon-edge ibft status
```

> Vote to add your validator to the network
```shell
VALIDATOR_KEY=$(polygon-edge ibft status --json | jq -r '.validator_key')
polygon-edge ibft propose --addr ${VALIDATOR_KEY} --vote auth
```

