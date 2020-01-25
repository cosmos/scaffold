package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

// clones an sdk repository and replaces the username and repo name with options from UserRepoArgs
func clone(args UserRepoArgs) error {
	tmp, err := ioutil.TempDir("", "scaffold-clone-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmp) // clean up

	cmd := exec.Command("git", "clone", "git@github.com:"+args.CloneUser+"/"+args.CloneRepo, tmp)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("rm", "-rf", tmp+"/.git")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	err = filepath.Walk(tmp, visit(args.CloneUser+"/"+args.CloneRepo, args.User+"/"+args.Repo))
	if err != nil {
		return (err)
	}
	err = filepath.Walk(tmp, visit("\\/"+args.CloneUser+"\\/"+args.CloneRepo+"/", "\\/"+args.User+"\\/"+args.Repo+"/"))
	if err != nil {
		return (err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return (err)
	}

	newdir := cwd + "/" + args.Repo

	if _, err := os.Stat(newdir); os.IsNotExist(err) {
		err := os.Rename(tmp, newdir)
		if err != nil {
			return (err)
		}
	} else {
		return errors.New("Folder " + newdir + " already exists")
	}

	if err != nil {
		return (err)
	}
	return nil
}
func visit(search, replace string) filepath.WalkFunc {
	return func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !!fi.IsDir() {
			return nil
		}
		matched, err := filepath.Match("*", fi.Name())
		if err != nil {
			panic(err)
			return err
		}
		if matched {
			read, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			newContents := strings.Replace(string(read), search, replace, -1)
			err = ioutil.WriteFile(path, []byte(newContents), 0)
			if err != nil {
				panic(err)
			}
		}
		return nil
	}
}

// Give this the name of the tutorial and any arguments necessary for
// said tutorial
func scaffold(name string, op string, args UserRepoArgs) error {

	// Fetch the scaffold files for the tutorial
	files := files(name)
	// Create the necessary folders for the files
	err := createUniqueFolders(files, op, args)
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
func files(name string) (out []string) {
	for _, n := range AssetNames() {
		if filepath.HasPrefix(n, name) {
			out = append(out, n)
		}
	}
	return
}

// Creates a unique folder for each passed in file
func createUniqueFolders(files []string, op string, args UserRepoArgs) error {
	for _, dir := range dirs(files, args) {
		err := os.MkdirAll(filepath.Join(op, dir), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func replaceDir(oldPath string, args UserRepoArgs) string {
	dir := args.Dir
	nameLowerCase := args.NameLowerCase
	repo := args.Repo
	newPath := strings.Replace(oldPath, dir, nameLowerCase, 1)
	if dir == "module" {
		return newPath
	}
	newPath = strings.Replace(newPath, dir, nameLowerCase, 1)
	return strings.Replace(newPath, nameLowerCase, repo, 1)
}

// Given a list of filepaths returns a list of directories
func dirs(files []string, args UserRepoArgs) (out []string) {
	for _, n := range files {
		replaced := replaceDir(filepath.Dir(n), args)
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

		replaced := replaceDir(f, args)

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
