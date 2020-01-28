# Cosmos Scaffold

Requirements:
go: 1.13+

This is a basic scaffolding tool for CosmosSDK applications. To build the binary, have [golang installed](https://golang.org/doc/install) and then run:

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

Flags:
  -c, --config string        config file (default is $HOME/.scaffold.yaml)
  -h, --help                 help for scaffold
  -o, --output-path string   Path to output
  -t, --toggle               Help message for toggle

Use "scaffold [command] --help" for more information about a command.
```

- [Usage information about `app`](./docs/app.md)
- [Usage information about `module`](./docs/module.md)