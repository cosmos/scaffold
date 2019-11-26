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

// tutCmd represents the tutorial generator, either nameservice or hellochain can be created
var tutCmd = &cobra.Command{
	Use:   "tutorial [tutorial-name] [user] [repo]",
	Short: "Generates one of the tutorial apps, currently either the 'nameservice' or 'hellochain'",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		ns := userRepoArgs{args[1], args[2]}
		if args[0] == "hellochain" || args[0] == "nameservice" {
			err := scaffold(args[0], outputPath, ns)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}
