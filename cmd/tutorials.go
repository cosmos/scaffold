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

		tutorial := args[0]
		nameRaw := args[3]
		if nameRaw == "" {
			nameRaw = tutorial
		}
		nameCapitalCamelCase := strcase.ToCamel(nameRaw)
		nameLowerCamelCase := strcase.ToLowerCamel(nameRaw)
		nameLowerCase := strings.ToLower(nameLowerCamelCase)

		ns := UserRepoArgs{
			Tutorial:             tutorial,
			User:                 args[1],
			Repo:                 args[2],
			NameRaw:              nameRaw,
			NameLowerCase:        nameLowerCase,
			NameCapitalCamelCase: nameCapitalCamelCase,
			NameLowerCamelCase:   nameLowerCamelCase,
		}
		if tutorial == "hellochain" || tutorial == "nameservice" {
			err := scaffold(tutorial, outputPath, ns)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}
