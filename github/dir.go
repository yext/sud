package github

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strings"
)

// Clones the repos that are github.
func CloneToDirs(repos []string) ([]string, error) {
	var dirs []string
	var err error

	usr, _ := user.Current()
	home := usr.HomeDir
	sudDir := filepath.Join(home, ".sud", "repos")

	err = clear(sudDir)
	if err != nil {
		fmt.Println("unable to clear directory " + sudDir)
		return repos, err
	}

	var githubRepos []string
	for _, dir := range repos {
		if !strings.HasPrefix(dir, "http://github.com") &&
			!strings.HasPrefix(dir, "https://github.com") {
			dirs = append(dirs, dir)
			continue
		}
		base := path.Base(dir)
		if !strings.Contains(base, "*") {
			githubRepos = append(githubRepos, dir)
			continue
		}
		matchedRepos, err := expand(dir)
		if err != nil {
			fmt.Printf("unable to expand %s, error: %v\n", dir, err)
			continue
		}
		if len(matchedRepos) > 0 {
			githubRepos = append(githubRepos, matchedRepos...)
		}
	}

	for _, repo := range githubRepos {
		dir, err := dirForRepoUrl(repo, sudDir)
		if err != nil {
			log.Printf("unable to create dir for repo %s, error: %v ", repo, err)
			continue
		}
		dirs = append(dirs, dir)
	}

	fmt.Println("Working on dirs: ")
	for _, dir := range dirs {
		fmt.Println(dir)
	}
	return dirs, nil
}

func clear(dir string) error {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("rm -rf %q", dir))
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func dirForRepoUrl(repoUrl string, sudDir string) (string, error) {
	dir := sudDir
	splits := strings.Split(repoUrl, "/")
	for _, s := range splits {
		if s == "" {
			continue
		}
		dir = filepath.Join(dir, s)
	}
	err := Clone(repoUrl, filepath.Dir(dir))
	if err != nil {
		return "", err
	}
	return dir, nil
}
