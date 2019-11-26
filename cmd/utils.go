package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Give this the name of the tutorial and any arguments necessary for
// said tutorial
func scaffold(name string, op string, args UserRepoArgs) error {
	// Fetch the scaffold files for the tutorial
	files := tutorialFiles(name)
	// Create the necessary folders for the files
	err := createUniqueFolders(files, op, args.Tutorial, args.NameLowerCase)
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
	for p, n := range AssetNames() {
		fmt.Println(p, n)
		if filepath.HasPrefix(n, name) {
			out = append(out, n)
		}
	}
	return
}

// Creates a unique folder for each passed in file
func createUniqueFolders(files []string, op, tutorial, nameLowerCase string) error {
	for _, dir := range dirs(files, tutorial, nameLowerCase) {
		fmt.Println("makedir", filepath.Join(op, dir))
		err := os.MkdirAll(filepath.Join(op, dir), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func replaceTutorial(oldPath, tutorial, nameLowerCase string) string {
	return oldPath[:0] + strings.Replace(oldPath[0:], tutorial, nameLowerCase, 1)
}

// Given a list of filepaths returns a list of directories
func dirs(files []string, tutorial, nameLowerCase string) (out []string) {
	for _, n := range files {
		replaced := replaceTutorial(filepath.Dir(n), tutorial, nameLowerCase)
		out = append(out, replaced)
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
func applyTemplates(files []string, op string, args UserRepoArgs) error {
	for _, f := range files {

		// i := strings.LastIndex(f, args.Tutorial)
		// oldFile := f[:0] + strings.Replace(f[0:], args.NameLowerCase, args.Tutorial, 1)
		// fmt.Println("oldFile", oldFile)
		fmt.Println("newFile", f)
		replaced := replaceTutorial(f, args.Tutorial, args.NameLowerCase)

		// fetch the template
		tb, err := Asset(f)
		if err != nil {
			return err
		}

		// Remove the .tmpl from the suffix
		np := filepath.Join(op, strings.TrimSuffix(replaced, ".tmpl"))

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
