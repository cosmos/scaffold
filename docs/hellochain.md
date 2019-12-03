# hellochain

To scaffold out the hellochain tutorial use the following format:

```bash
scaffold tutorial hellochain <github user||org> <repoName>
```

To change the name of the project to "My Cool Blockchain" (which will create `MyCoolBlockchainApp`) use the following format:

```bash
scaffold tutorial hellochain <user/Github org> <repoName> "My Cool Blockchain"
```

Then just `cd` into the repo folder and:

```bash
go mod tidy
make install
hccli --help
hcd --help
```

Once you have done that run the following commands to start a local development chain:

```bash
hcd init <moniker>
hccli keys add validator
hcd add-genesis-account validator 100000000stake
hcd start
```