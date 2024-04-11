#!/bin/bash

# clear the terminal
clear

# Stop execution if any command fails
set -e

KEYRING_BACKEND="os"
PASSWORD="password"

export LAYERD_NODE_HOME="$HOME/.layer/alice"

# Remove old test chains (if present)
echo "Removing old test chain data..."
rm -rf ~/.layer

# Build layerd
echo "Building layerd..."
go build ./cmd/layerd

# Initialize the chain
echo "Initializing the chain..."
./layerd init layer --chain-id layer

# Init two different chain nodes with two different folders
echo "Initializing chain nodes..."
echo "alice..."
./layerd init alicemoniker --chain-id layer --home ~/.layer/alice
echo "bill..."
./layerd init billmoniker --chain-id layer --home ~/.layer/bill

# Add a validator account alice
echo "Adding validator accounts..."
echo "alice..."
./layerd keys add alice --keyring-backend $KEYRING_BACKEND --home ~/.layer/alice
echo "bill..."
./layerd keys add bill --keyring-backend $KEYRING_BACKEND --home ~/.layer/bill
echo "charlie..."
yes | ./layerd keys add charlie --keyring-backend os --home ~/.layer/alice > ~/Desktop/charlie_key_info.txt 2>&1

# Extract the mnemonic from the key_info file
echo "Extracting charlie's mnemonic from key_info file..."
grep -A 24 'It is the only way to recover your account if you ever forget your password.' ~/Desktop/charlie_key_info.txt | tail -n 1 > ~/Desktop/charlie_mnemonic.txt


# Update vote_extensions_enable_height in genesis.json
echo "Updating vote_extensions_enable_height in genesis.json..."
echo "main..."
jq '.consensus.params.abci.vote_extensions_enable_height = "1"' ~/.layer/config/genesis.json > temp.json && mv temp.json ~/.layer/config/genesis.json
echo "alice..."
jq '.consensus.params.abci.vote_extensions_enable_height = "1"' ~/.layer/alice/config/genesis.json > temp.json && mv temp.json ~/.layer/alice/config/genesis.json
echo "bill..."
jq '.consensus.params.abci.vote_extensions_enable_height = "1"' ~/.layer/bill/config/genesis.json > temp.json && mv temp.json ~/.layer/bill/config/genesis.json

# Create a tx to give alice loyas to stake
echo "Adding genesis accounts..."
echo "alice..."
./layerd genesis add-genesis-account $(./layerd keys show alice -a --keyring-backend $KEYRING_BACKEND --home ~/.layer/alice)  10000000000000loya --keyring-backend $KEYRING_BACKEND --home ~/.layer/alice
echo "bill..."
./layerd genesis add-genesis-account $(./layerd keys show bill -a --keyring-backend $KEYRING_BACKEND --home ~/.layer/bill) 10000000000000loya --keyring-backend $KEYRING_BACKEND --home ~/.layer/alice
echo "charlie..."
./layerd genesis add-genesis-account $(./layerd keys show charlie -a --keyring-backend $KEYRING_BACKEND --home ~/.layer/alice) 10000000000000loya --keyring-backend $KEYRING_BACKEND --home ~/.layer/alice
# ./layerd genesis add-genesis-account $(./layerd keys show bill -a --keyring-backend os --home ~/.layer/bill) 10000000000000loya --keyring-backend os --home ~/.layer/bill

# Create a tx to stake some loyas for alice
echo "Creating gentx alice..."
./layerd genesis gentx alice 1000000000000loya --chain-id layer --keyring-backend $KEYRING_BACKEND --home ~/.layer/alice --keyring-dir ~/.layer/alice

# Add the transactions to the genesis block:q
echo "Collecting gentxs..."
./layerd genesis collect-gentxs --home ~/.layer/alice

# Export alice key from os backend and import to test backend
echo "Exporting alice key..."
echo $PASSWORD | ./layerd keys export alice --keyring-backend $KEYRING_BACKEND --home ~/.layer/alice > ~/Desktop/alice_keyfile
echo "Importing alice key to test backend..."
echo $PASSWORD | ./layerd keys import alice ~/Desktop/alice_keyfile --keyring-backend test --home ~/.layer/alice

# Export bill key from os backend and import to test backend
echo "Exporting bill key..."
echo $PASSWORD | ./layerd keys export bill --keyring-backend $KEYRING_BACKEND --home ~/.layer/bill > ~/Desktop/bill_keyfile
echo "Importing bill key to test backend..."
echo $PASSWORD | ./layerd keys import bill ~/Desktop/bill_keyfile --keyring-backend test --home ~/.layer/bill

# Export charlie key from os backend and import to test backend
echo "Exporting charlie key..."
echo $PASSWORD | ./layerd keys export charlie --keyring-backend $KEYRING_BACKEND --home ~/.layer/alice > ~/Desktop/charlie_keyfile
echo "Importing charlie key to test backend..."
echo $PASSWORD | ./layerd keys import charlie ~/Desktop/charlie_keyfile --keyring-backend test --home ~/.layer/alice

# Delete the keyfiles
echo "Deleting keyfiles..."
rm ~/Desktop/alice_keyfile
rm ~/Desktop/bill_keyfile
rm ~/Desktop/charlie_keyfile

# Modify timeout_commit in config.toml for alice
echo "Modifying timeout_commit in config.toml for alice..."
sed -i '' 's/timeout_commit = "5s"/timeout_commit = "500ms"/' ~/.layer/alice/config/config.toml

# sleep 30

echo "Start chain..."
./layerd start --home $LAYERD_NODE_HOME --api.enable --api.swagger
