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

## App

To scaffold out a ready to go general-purpose app run the following:

```bash
scaffold app [lvl-1] <user/Github org> <repoName>
```

There are different levels of apps, this mainly separates how many core Cosmos-SDK modules the app is initialized with.

There is no need to make a directory this will happen with the command.

### Levels

- `lvl-1`: Auth, Bank, Distribution, Genutil, Genaccounts, Params, Slashing, Slashing, Supply

- More levels to come...

After you have chosen your level and created your app, you will need to `cd` into the directory then run
`go get ./...` and then you will be able to run `make install`

## Module

To scaffold out a empty module:

- First: `cd` into your modules directory, if you used the scaffold tool to generate your app this will be you `/x` folder
- Second: run the command listed below:

```bash
scaffold module <user/Github org> <repoName> <moduleName>
```

This will get you started with writing a module.
The layout of the files follow the module spec, located [here](https://github.com/cosmos/cosmos-sdk/blob/0992c2994ca15131712ab19369f558190434f231/docs/building-modules/structure.md).

## Tutorials

To scaffold out the `nameservice` tutorial example just run the following:

```bash
scaffold tutorial nameservice <user/Github org> <repoName> "Blockchain Name"
```

Then just `cd` into the `myrepo` folder and:

```bash
go mod tidy
make install
nscli --help
nsd --help
```

## Versioning

This repo will have its own versioning and contain a compatibility table:

| Cosmos-SDK Version | Scaffold Version |
| ------------------ | :--------------: |
| 0.37.\*            |      0.1.0       |
