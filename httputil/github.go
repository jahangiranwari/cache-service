package httputil

import (
	"os"
	"strings"
)

var allowedRepos []string

func init() {
	allowedRepos = strings.Split(os.Getenv("GITHUB_REPOS"), ",")
}

func GetGitHubRepo(repo string) string {
	if allowRepo(repo) {
		url := os.Getenv("GITHUB_REPOS_URL") + repo
		json, _ := GetJSONWithAuth(url, os.Getenv("GITHUB_TOKEN"))
		return json
	}
	return `{"data": "Forbidden"}`
}

func allowRepo(repo string) bool {
	for _, v := range allowedRepos {
		if v == repo {
			return true
		}
	}
	return false
}
