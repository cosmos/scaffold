# Nameservice

To scaffold out the nameservice tutorial use the following format:

```bash
scaffold tutorial nameservice <github user||org> <repoName>
```

To change the name of the project to "My Cool Blockchain" (which will create `MyCoolBlockchainApp`) use the following format:

```bash
scaffold tutorial nameservice <user/Github org> <repoName> "My Cool Blockchain"
```

Then just `cd` into the repo folder and:

```bash
go mod tidy
make install
nscli --help
nsd --help
```

Once you have done that run the following commands to start a local development chain:

```bash
nsd init <moniker>
nscli keys add validator
nsd add-genesis-account validator 100000000stake
nsd gentx --name validator
nsd collect-gentxs
nsd start
```
