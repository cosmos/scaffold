# Module

To scaffold out an empty module:

> Note: if you need a basic application see the [`app` section](./app.md) to generate one first

1. `cd` into your modules directory (typically named `/x`).
2. Run the command listed below:

```bash
scaffold module <user/Github org> <repoName> <moduleName>
```

This will get you started with writing a module.
The layout of the files follow the module spec, located [here](https://github.com/cosmos/cosmos-sdk/blob/v0.37.3/docs/modules/SPEC.md).

## Versioning

This repo will have its own versioning and contain a compatibility table:

| Cosmos-SDK Version | Scaffold Version |
| ------------------ | :--------------: |
| 0.37.3             |      0.1.0       |
