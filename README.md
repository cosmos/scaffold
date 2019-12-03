# Cosmos Scaffold

Requirements:
go: 1.13+

This is a scaffolding tool for CosmosSDK based applications. To build the binary, have [golang installed](https://golang.org/doc/install) and then run:

```bash
make
```

The binary should be installed in your `$GOBIN` so you can then run:

```bash
scaffold --help
```
and you will see:

```bash
This CLI helps in scaffolding out CosmosSDK based applications

Usage:
  scaffold [command]

Available Commands:
  app         Generates an an empty application boilerplate 
  help        Help about any command
  module      Generate an empty module for use in the Cosmos-SDK
  tutorial    Generates one of the tutorial apps, currently either the 'nameservice' or 'hellochain'

Flags:
  -c, --config string        config file (default is $HOME/.scaffold.yaml)
  -h, --help                 help for scaffold
  -o, --output-path string   Path to output
  -t, --toggle               Help message for toggle

Use "scaffold [command] --help" for more information about a command.

```

## Tutorials
This command will scaffold out a copy of one of the tutorials from [cosmos-sdk/tutorials](https://github.com/cosmos/tutorials) into a new directory.

For example, to scaffold out the nameservice tutorial use the following format:

```bash
scaffold tutorial nameservice <user/Github org> <repoName>
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

## General-Purpose App

To scaffold out a ready to go general-purpose app into a new directory run the following:

```bash
scaffold app [lvl-1] <user/Github org> <repoName>
```

There are different levels of apps, this mainly separates how many core Cosmos-SDK modules the app is initialized with.

### Levels

- `lvl-1`: Auth, Bank, Distribution, Genutil, Genaccounts, Params, Slashing, Staking, Supply

After you have chosen your level and created your app you will need to `cd` into the directory and run
`go get ./...`. Then you will be able to run `make install`

## Module

To scaffold out an empty module:

1. `cd` into your modules directory (typically named `/x`).
2. Run the command listed below:

```bash
scaffold module <user/Github org> <repoName> <moduleName>
```

This will get you started with writing a module.
The layout of the files follow the module spec, located [here](https://github.com/cosmos/cosmos-sdk/blob/0992c2994ca15131712ab19369f558190434f231/docs/building-modules/structure.md).

## Versioning

This repo will have its own versioning and contain a compatibility table:

| Cosmos-SDK Version | Scaffold Version |
| ------------------ | :--------------: |
| 0.37.\*            |      0.1.0       |
