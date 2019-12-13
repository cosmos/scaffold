package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClone(t *testing.T) {
	tmp, err := ioutil.TempDir("", "scaffold-clone-test-")
	if err != nil {
		log.Fatal(err)
	}
	// defer os.RemoveAll(tmp) // clean up
	os.Chdir(tmp)
	err = clone(UserRepoArgs{
		User:      "username",
		Repo:      "repository",
		CloneUser: "cosmos",
		CloneRepo: "gaia",
	})
	require.NoError(t, err)
}
