package cmd

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// moduleCmd
var moduleCmd = &cobra.Command{
	Use:   "module [user] [repo] [moduleName]",
	Short: "Generate an empty module for use in the Cosmos-SDK",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		nameRaw := args[2]

		// nameCapitalCamelCase := strcase.ToCamel(nameRaw)
		// nameLowerCamelCase := strcase.ToLowerCamel(nameRaw)
		nameLowerCase := strings.ToLower(nameRaw)

		mdl := UserRepoArgs{
			User:          args[0],
			Repo:          args[1],
			Dir:           "module",
			NameRaw:       nameRaw,
			NameLowerCase: nameLowerCase,
		}
		err := scaffold("module", outputPath, mdl)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

// appCmd enables scaffolding of different levels of apps
var appCmd = &cobra.Command{
	Use:   "app [lvl] [user] [repo]",
	Short: "Generates an empty application boilerplate",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		nameRaw := args[2]

		nameCapitalCamelCase := strcase.ToCamel(nameRaw)
		nameLowerCamelCase := strcase.ToLowerCamel(nameRaw)
		nameLowerCase := strings.ToLower(nameRaw)

		ns := UserRepoArgs{
			Dir:                  args[0],
			User:                 args[1],
			Repo:                 args[2],
			NameRaw:              nameRaw,
			NameCapitalCamelCase: nameCapitalCamelCase,
			NameLowerCamelCase:   nameLowerCamelCase,
			NameLowerCase:        nameLowerCase,
		}
		err := scaffold(args[0], outputPath, ns)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
