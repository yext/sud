package github

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

var (
	ee *exec.ExitError
	pe *os.PathError
)

// Pushes the given directories using git push.
func Push(dirs []string) error {
	for _, dir := range dirs {
		repo, err := git.PlainOpen(dir)
		worktree, err := repo.Worktree()
		status, err := worktree.Status()

		if err != nil {
			log.Printf("general error: %v", err) // something really bad happened!
			continue
		} else if !status.IsClean() {
			fmt.Println("There are changes in " + dir + ", pushing to github.") // ran, but non-zero exit code
		} else {
			fmt.Println("No changes found. Nothing to push.") // ran without error (exit code zero)
			continue
		}

		err = worktree.AddGlob(".")
		if err != nil {
			continue
		}

		commitOpt := git.CommitOptions{
			All: true,
		}
		_, err = worktree.Commit("sud: updating the repo", &commitOpt)
		if err != nil {
			fmt.Println("Failed to commit " + dir)
			fmt.Println(err)
			continue
		}

		accessToken, err := getAccessToken()

		if err != nil {
			fmt.Println(err)
			continue
		}

		pushOpt := git.PushOptions{
			Auth: &http.BasicAuth{
				Username: "name",
				Password: accessToken,
			},
		}
		err = repo.Push(&pushOpt)
		if err != nil {
			fmt.Println("Failed to push " + dir)
			fmt.Println(err)
			continue
		}

		fmt.Println("Pushed to github.")
	}
	return nil
}

func convertTilde(path string, home string) string {
	if path == "~" {
		// In case of "~", which won't be caught by the "else if"
		path = home
	} else if strings.HasPrefix(path, "~/") {
		// Use strings.HasPrefix so we don't match paths like
		// "/something/~/something/"
		path = filepath.Join(home, path[2:])
	}
	return path
}
