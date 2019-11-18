# Cosmos Scaffold

This is a scaffolding tool for CosmosSDK based applications. To build the binary, have v1.13 [golang installed](https://golang.org/doc/install) and then just run:

```bash
make
```

Then the binary should be installed in your `$GOBIN` so you can then run:

```bash
scaffold --help
```

To scaffold out the `nameservice` tutorial example just run the following:

```bash
cd $GOPATH/src/github.com/{{ .Username }}
scaffold nameservice myghusername myrepo
```

Then just `cd` into the `myrepo` folder and:

```bash
go mod tidy
make install
nscli --help
nsd --help
```
