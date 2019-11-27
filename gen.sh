#!/usr/bin/env bash

rm -rf ~/.aud
rm -rf ~/.acli
aud init a --chain-id testchain
echo "12345678" | acli keys add test1
echo "12345678" | acli keys add test2

aud add-genesis-account $(acli keys show test1 -a) 10000000000000000000000000stake,1000000legends
acli config output json
acli config indent true
acli config trust-node true

echo "12345678" | aud gentx --name test1

echo "Collecting genesis txs..."
aud collect-gentxs

echo "Validating genesis file..."
aud validate-genesis

aud start 