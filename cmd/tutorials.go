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
	"os"
	"path/filepath"
	"strings"
	"text/template"

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
		err := scaffoldTutorial("nameservice", outputPath, ns)
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
		err := scaffoldTutorial("hellochain", outputPath, ns)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

// Give this the name of the tutorial and any arguements necessary for
// said tutorial
func scaffoldTutorial(name string, op string, args interface{}) error {
	// Fetch the scaffold files for the tutorial
	files := tutorialFiles(name)

	// Create the necessary folders for the files
	err := createUniqueFolders(files, op)
	if err != nil {
		return err
	}

	// Create the files and fill them with the applied templates
	err = applyTemplates(files, op, args)
	if err != nil {
		return err
	}

	return nil
}

// Get all the files from the tutorial
func tutorialFiles(name string) (out []string) {
	for _, n := range AssetNames() {
		if filepath.HasPrefix(n, name) {
			out = append(out, n)
		}
	}
	return
}

// Creates a unique folder for each passed in file
func createUniqueFolders(files []string, op string) error {
	for _, dir := range dirs(files) {
		err := os.MkdirAll(filepath.Join(op, dir), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// Given a list of filepaths returns a list of directories
func dirs(files []string) (out []string) {
	for _, n := range files {
		out = append(out, filepath.Dir(n))
	}
	return uniqueStrings(out)
}

// Returns a list of unique strings from a list
func uniqueStrings(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}
	return us
}

// Given a list of templates, create the mirror rendered template files
// in the given output path. Should run createUniqueFolders first
func applyTemplates(files []string, op string, args interface{}) error {
	for _, f := range files {
		// fetch the template
		tb, err := Asset(f)
		if err != nil {
			return err
		}

		// Remove the .tmpl from the suffix
		np := filepath.Join(op, strings.TrimSuffix(f, ".tmpl"))

		// Create the *template.Template with name f
		t, err := template.New(np).Parse(string(tb))
		if err != nil {
			return err
		}

		// Create the file at path f
		f, err := os.Create(np)
		if err != nil {
			return err
		}

		// Render template and write it to file
		err = t.Execute(f, args)
		if err != nil {
			return err
		}

		fmt.Printf("Created %s...\n", np)
	}
	return nil
}

// Register the commands in init
func init() {
	rootCmd.AddCommand(nsCmd)
	rootCmd.AddCommand(hcCmd)
}
