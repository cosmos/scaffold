# Cosmos Scaffold

Requirements:
go: 1.13+

This is a scaffolding tool for CosmosSDK based applications. To build the binary, have [golang installed](https://golang.org/doc/install) and then just run:

```bash
make
```

Then the binary should be installed in your `$GOBIN` so you can then run:

```bash
scaffold --help
```

## Module

To scaffold out a empty module:

- First: `cd` into your modules directory
- Second: run the command listed below:

```bash
scaffold module <user/Github org> <repoName> <moduleName>
```

This will get you started with writing a module.
The layout of the files follow the module spec, located [here](https://github.com/cosmos/cosmos-sdk/blob/0992c2994ca15131712ab19369f558190434f231/docs/building-modules/structure.md).

## Tutorials

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
