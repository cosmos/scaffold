# Cosmos Scaffold

This is a scaffolding tool for CosmosSDK based applications. To build the binary, have v1.13 [golang installed](https://golang.org/doc/install) and then run:

```bash
make
```

Then the binary should be installed in your `$GOBIN` so you can then run:

```bash
scaffold --help
```

## Tutorial

To scaffold out the `nameservice` tutorial example just run the following:

```bash
cd $GOPATH/src/github.com/{{ .Username }}
scaffold tutorial nameservice myghusername myrepo
```

Then just `cd` into the `myrepo` folder and:

```bash
go mod tidy
make install
nscli --help
nsd --help
```

# Versioning

This repo will have its own versioning and contain a compatibility table:

| Cosmos-SDK Version | Scaffold Version |
| ------------------ | :--------------: |
| 0.37.\*            |      0.1.0       |
