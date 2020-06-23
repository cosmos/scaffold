# Level 1
The Level 1 app includes the following modules:

* Auth
* Bank
* Distribution
* Genutil
* Genaccounts
* Params
* Slashing
* Staking
* Supply


To scaffold out the `lvl-1` app use the following format:

```bash
scaffold app lvl-1 <github user||org> <repoName>
```

Then just `cd` into the repo folder and:

```bash
go mod tidy
make install
acli --help
aud --help
```

Once you have done that run the following commands to start a local development chain:

```bash
# INIT CREATES ~/.appd 
# moniker IS THE NAME OF YOUR VALIDATOR ON THE NETWORK YOU'LL START, CHOOSE YOUR USUAL HANDLE OR VALIDATOR BRAND
appd init <moniker>
# ADD A KEY NAMED VALIDATOR
appcli keys add validator
# ADD AN ACCOUNT TO GENESIS.JSON THAT COORESPONDS TO YOUR VALIDATOR KEY
appd add-genesis-account validator 100000000stake
# SIGN THE GENESIS TRANSACTION WITH YOUR KEY NAMED VALIDATOR
appd gentx --name validator
# COLLECT GENESIS TRANSACTIONS
# THIS ALLOWS THE CHAIN TO START IN A DECENTRALIZED MANNER, WITH GENESIS.JSON APPROVED BY ALL OF THE VALIDATORS-TO-BE
appd collect-gentxs
# START RUNNING THE CHAIN
appd start
```
