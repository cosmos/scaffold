package cmd

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// tutCmd represents the tutorial generator, either nameservice or hellochain can be created
var tutCmd = &cobra.Command{
	Use:   "tutorial [tutorial-name] [user] [repo] [name]",
	Short: "Generates one of the tutorial apps, currently either the 'nameservice' or 'hellochain'",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		nameRaw := ""
		dir := args[0]

		if len(args) == 4 {
			nameRaw = args[3]
		}

		if nameRaw == "" {
			nameRaw = dir
		}
		nameCapitalCamelCase := strcase.ToCamel(nameRaw)
		nameLowerCamelCase := strcase.ToLowerCamel(nameRaw)
		nameLowerCase := strings.ToLower(nameLowerCamelCase)

		ns := UserRepoArgs{
			Dir:                  dir,
			User:                 args[1],
			Repo:                 args[2],
			NameRaw:              nameRaw,
			NameLowerCase:        nameLowerCase,
			NameCapitalCamelCase: nameCapitalCamelCase,
			NameLowerCamelCase:   nameLowerCamelCase,
		}
		if dir == "hellochain" || dir == "nameservice" {
			err := scaffold(dir, outputPath, ns)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}

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
		ns := UserRepoArgs{
			Dir:  args[0],
			User: args[1],
			Repo: args[2],
		}
		err := scaffold(args[0], outputPath, ns)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
