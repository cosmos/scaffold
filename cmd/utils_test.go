package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/iancoleman/strcase"
	"github.com/stretchr/testify/require"
)

func TestScaffoldTutorialNameservice(t *testing.T) {
	tmp, err := ioutil.TempDir("", "scaffold-test-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmp) // clean up
	os.Chdir(tmp)

	user := "username"
	repo := "repository"
	dir := "nameservice"
	nameRaw := dir
	nameCapitalCamelCase := strcase.ToCamel(nameRaw)
	nameLowerCamelCase := strcase.ToLowerCamel(nameRaw)
	nameLowerCase := strings.ToLower(nameLowerCamelCase)
	ns := UserRepoArgs{
		Dir:                  dir,
		User:                 user,
		Repo:                 repo,
		NameRaw:              nameRaw,
		NameLowerCase:        nameLowerCase,
		NameCapitalCamelCase: nameCapitalCamelCase,
		NameLowerCamelCase:   nameLowerCamelCase,
	}
	err = scaffold(dir, outputPath, ns)
	require.NoError(t, err)
}
func TestScaffoldTutorialHelloChain(t *testing.T) {
	tmp, err := ioutil.TempDir("", "scaffold-test-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmp) // clean up
	os.Chdir(tmp)

	user := "username"
	repo := "repository"
	dir := "hellochain"
	nameRaw := dir
	nameCapitalCamelCase := strcase.ToCamel(nameRaw)
	nameLowerCamelCase := strcase.ToLowerCamel(nameRaw)
	nameLowerCase := strings.ToLower(nameLowerCamelCase)
	ns := UserRepoArgs{
		Dir:                  dir,
		User:                 user,
		Repo:                 repo,
		NameRaw:              nameRaw,
		NameLowerCase:        nameLowerCase,
		NameCapitalCamelCase: nameCapitalCamelCase,
		NameLowerCamelCase:   nameLowerCamelCase,
	}
	err = scaffold(dir, outputPath, ns)
	require.NoError(t, err)
}

func TestScaffoldApp(t *testing.T) {
	tmp, err := ioutil.TempDir("", "scaffold-test-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmp) // clean up
	os.Chdir(tmp)

	lvl := "lvl-1"
	dir := lvl
	user := "username"
	repo := "repository"

	ns := UserRepoArgs{
		Dir:  dir,
		User: user,
		Repo: repo,
	}
	err = scaffold(dir, outputPath, ns)
	require.NoError(t, err)
}

func TestScaffoldModule(t *testing.T) {
	tmp, err := ioutil.TempDir("", "scaffold-test-")
	if err != nil {
		log.Fatal(err)
	}
	// defer os.RemoveAll(tmp) // clean up
	os.Chdir(tmp)

	dir := "module"
	user := "username"
	repo := "repository"
	nameRaw := "myModule"
	nameLowerCase := strings.ToLower(nameRaw)
	mdl := UserRepoArgs{
		User:          user,
		Repo:          repo,
		Dir:           dir,
		NameRaw:       nameRaw,
		NameLowerCase: nameLowerCase,
	}
	err = scaffold(dir, outputPath, mdl)
	require.NoError(t, err)
}
