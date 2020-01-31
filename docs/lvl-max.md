# Level Max

To scaffold out the `lvl-max` app use the following format:

```bash
scaffold app lvl-max <github user||org> <repoName>
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
aud init <moniker>
acli keys add validator
aud add-genesis-account validator 100000000stake
aud gentx --name validator
aud collect-gentxs
aud start
```

## Contribution

Contributed by [Kimura Yu](https://github.com/KimuraYu45z) , [LCNEM, Inc.](https://github.com/lcnem)
