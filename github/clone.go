package github

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// Clones the given repo to the directory.
func Clone(repoUrl string, dir string) error {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	cloneOpts := git.CloneOptions{
		URL: repoUrl,
	}
	accessToken, aerr := getAccessToken()
	if aerr == nil {
		cloneOpts.Auth = &http.BasicAuth{
			Username: "name",
			Password: accessToken,
		}
	}

	// TODO: Pass through repo instead of rereading from folder later.
	_, err = git.PlainClone(dir, false, &cloneOpts)
	if err != nil {
		fmt.Println("Clone failed. You may need to specify a personal access token.")
		fmt.Println(aerr)
		return err
	}
	return nil
}
