package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tutCmd represents the tutorial generator, either nameservice or hellochain can be created
var appCmd = &cobra.Command{
	Use:   "app [lvl] [user] [repo]",
	Short: "Generates a ",
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
