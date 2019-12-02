package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var moduleCMD = &cobra.Command{
	Use:   "module [user] [repo] [moduleName]",
	Short: "Generate a empty module for use in the Cosmos-SDK",
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

// Register the commands in init
func init() {
	rootCmd.AddCommand(moduleCMD)
}
