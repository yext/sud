package github

import (
	"fmt"
	"os"
	"os/exec"
)

// Clones the given repo to the directory.
func Clone(repoUrl string, dir string) error {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	cmd := exec.Command("sh", "-c", fmt.Sprintf("git clone %q", repoUrl))
	cmd.Dir = dir
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
