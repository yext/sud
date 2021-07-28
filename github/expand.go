package github

import (
	"context"
	"fmt"
	"github.com/gobwas/glob"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"path"
	"strings"
)

// Expands a given repo url pattern with wildcard to the matching repository urls.
func expand(repoUrlPattern string) ([]string, error) {
	prefixUrl := repoUrlPattern[:strings.LastIndex(repoUrlPattern, "/")]
	org := path.Base(prefixUrl)
	pattern := path.Base(repoUrlPattern)
	matchedNames, err := match(org, pattern)
	if err != nil {
		return nil, err
	}
	var expandedRepos []string
	for _, matchedName := range matchedNames {
		expandedRepos = append(expandedRepos, prefixUrl+"/"+matchedName)
	}
	return expandedRepos, nil
}

// Matches the pattern to repository names in the given org.
func match(org string, pattern string) ([]string, error) {
	g := glob.MustCompile(pattern)
	var matchedRepos []string
	repos, err := list(org)
	if err != nil {
		return nil, err
	}
	for _, r := range repos {
		if g.Match(r) {
			matchedRepos = append(matchedRepos, r)
		}
	}
	return matchedRepos, nil
}

// Lists all the repository names in the given org.
func list(org string) ([]string, error) {
	var rs []string
	ctx := context.Background()
	const GithubAccessToken = "GITHUB_ACCESS_TOKEN"
	accessToken, found := os.LookupEnv(GithubAccessToken)
	if !found {
		return nil, fmt.Errorf("please set the env var " + GithubAccessToken +
			" to the value of access token generated from https://github.com/settings/tokens")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.ListByOrg(ctx, org, &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 200,
		},
	})
	if err != nil {
		return nil, err
	}
	for _, r := range repos {
		rs = append(rs, *r.Name)
	}
	return rs, nil
}
