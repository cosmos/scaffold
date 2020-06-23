# Level Max

To scaffold out the `lvl-max` app use the following format:

```bash
scaffold app lvl-max <github user||org> <repoName>
```

Then just `cd` into the repo folder and:

```bash
go mod tidy
make install
appcli --help
appd --help
```

Once you have done that run the following commands to start a local development chain:

```bash
appd init <moniker>
appcli keys add validator
appd add-genesis-account validator 100000000stake
appd gentx --name validator
appd collect-gentxs
appd start
```

## Contribution

Contributed by [Kimura Yu](https://github.com/KimuraYu45z) , [LCNEM, Inc.](https://github.com/lcnem)
