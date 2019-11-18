/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Arguments for nameservice scaffolding
type userRepoArgs struct {
	User string `json:"user"`
	Repo string `json:"repo"`
}

// nsCmd represents the nameservice tutorial generator
var nsCmd = &cobra.Command{
	Use:   "nameservice [user] [repo]",
	Short: "Generates the nameservice application",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ns := userRepoArgs{args[0], args[1]}
		err := scaffold("nameservice", outputPath, ns)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

// hcCmd represents the hellochain tutorial generator
var hcCmd = &cobra.Command{
	Use:   "hellochain [user] [repo]",
	Short: "Generates the hellochain application",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ns := userRepoArgs{args[0], args[1]}
		err := scaffold("hellochain", outputPath, ns)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

// Register the commands in init
func init() {
	rootCmd.AddCommand(nsCmd)
	rootCmd.AddCommand(hcCmd)
}
