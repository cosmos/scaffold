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

// Arguments for nameservice scaffolding
type UserRepoArgs struct {
	Tutorial             string `json:"tutorial"`
	User                 string `json:"user"`
	Repo                 string `json:"repo"`
	NameRaw              string `json:"nameRaw"`
	NameLowerCase        string `json:"nameLowerCase"`
	NameCapitalCamelCase string `json:"nameCapitalCamelCase"`
	NameLowerCamelCase   string `json:"nameLowerCamelCase"`
}

// tutCmd represents the tutorial generator, either nameservice or hellochain can be created
var tutCmd = &cobra.Command{
	Use:   "tutorial [tutorial-name] [user] [repo] [name]",
	Short: "Generates one of the tutorial apps, currently either the 'nameservice' or 'hellochain'",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {

		nameRaw := args[3]
		nameCapitalCamelCase := strcase.ToCamel(nameRaw)
		nameLowerCamelCase := strcase.ToLowerCamel(nameRaw)
		nameLowerCase := strings.ToLower(nameLowerCamelCase)

		ns := UserRepoArgs{args[0], args[1], args[2], nameRaw, nameLowerCase, nameCapitalCamelCase, nameLowerCamelCase}
		if args[0] == "hellochain" || args[0] == "nameservice" {
			err := scaffold(args[0], outputPath, ns)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}
