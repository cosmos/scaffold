# Contributing

This repo is built to make adding additional templates as easy as possible. The following are steps for adding additional templates.

1. Make your templates: Take the app you want to template and add `{{ .Properties }}` in all the places you want the user entered variables to be added. See the existing templates for some examples.
2. Once your template is ready to go, you will need to add it to either `tutorial`, `app` or other commands. To see how that works check out `cmd/commands.go`
3. Rebuild the `scaffold` (`make install`) and test your scaffold. 
4. Add documentation on how to use your scaffold in the `docs` folder and associate it with the appropriate subcommand.