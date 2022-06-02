package github

import (
	"fmt"
	"os"
)

// Access tokens created at https://github.com/settings/tokens
// Click on profile image, go to Settings -> Developer settings -> Personal access tokens
func getAccessToken() (string, error) {
	const GithubAccessToken = "GITHUB_ACCESS_TOKEN"
	accessToken, found := os.LookupEnv(GithubAccessToken)
	if !found {
		return "", fmt.Errorf("please set the env var " + GithubAccessToken +
			" to the value of access token generated from https://github.com/settings/tokens" +
			"/n Click on profile image, go to Settings -> Developer settings -> Personal access tokens.")
	}
	return accessToken, nil
}
