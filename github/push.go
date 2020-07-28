package github

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

var (
	ee *exec.ExitError
	pe *os.PathError
)

// Pushes the given directories using git push.
func Push(dirs []string) error {
	usr, _ := user.Current()
	home := usr.HomeDir
	for _, dir := range dirs {
		cmd := exec.Command("sh", "-c", "git diff --quiet")
		cmd.Dir = convertTilde(dir, home)
		err := cmd.Run()
		if errors.As(err, &ee) {
			fmt.Println("There are changes in " + dir + ", pushing to github.") // ran, but non-zero exit code
		} else if errors.As(err, &pe) {
			log.Printf("os.PathError: %v", pe) // "no such file ...", "permission denied" etc.
			continue
		} else if err != nil {
			log.Printf("general error: %v", err) // something really bad happened!
			continue
		} else {
			fmt.Println("No changes found. Nothing to push.") // ran without error (exit code zero)
			continue
		}

		cmd = exec.Command("sh", "-c", "git commit -am'sud: updating the repo'")
		cmd.Dir = convertTilde(dir, home)
		err = cmd.Run()
		if err != nil {
			continue
		}

		cmd = exec.Command("sh", "-c", "git push")
		cmd.Dir = convertTilde(dir, home)
		err = cmd.Run()
		if err != nil {
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
