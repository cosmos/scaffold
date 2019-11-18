package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type moduleArgs struct {
	User       string `json:"user"`
	Repo       string `json:"repo"`
	ModuleName string `json:"module_name"`
}

var moduleCMD = &cobra.Command{
	Use:   "module [user] [repo] [moduleName]",
	Short: "Generate a empty module for use in the Cosmos-SDK",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		mdl := moduleArgs{
			User:       args[0],
			Repo:       args[1],
			ModuleName: args[2],
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
