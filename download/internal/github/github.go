package github

import "fmt"

const HTTP_PROTOCOL = "http"
const HTTPS_PROTOCOL = "https"

const GITHUB_DOMAIN = "github.com"
const GITHUB_API_URL = "api.github.com"
const GITHUB_RAW_USER_CONTENT_URL = "raw.githubusercontent.com"

func ReoisitoryRegexp() string {
	return fmt.Sprintf("^%s://%s/([^/]+)/([^/]+)(/(tree|blob)/([^/]+)(/(.*))?)?", HTTPS_PROTOCOL, GITHUB_DOMAIN)
}

func RawUserContentUrl(username string, repo string, branch string, path string) string {
	return fmt.Sprintf("%s://%s/%s/%s/%s/%s", HTTPS_PROTOCOL, GITHUB_RAW_USER_CONTENT_URL, username, repo, branch, path)
}

func ContentApiUrl(username string, repo string, branch string, path string) string {
	return fmt.Sprintf("%s://%s/repos/%s/%s/contents/%s?ref=%s", HTTPS_PROTOCOL, GITHUB_API_URL, username, repo, path, branch)
}

func GitTreesApiUrl(username string, repo string, sha1 string, isRecursive bool) string {
	recursive := ""
	if isRecursive {
		recursive = "?recursive=1"
	}

	return fmt.Sprintf("%s://%s/repos/%s/%s/git/trees/%s%s", HTTPS_PROTOCOL, GITHUB_API_URL, username, repo, sha1, recursive)
}

func RepositoryRootUrl(username string, repo string) string {
	return fmt.Sprintf("%s://%s/%s/%s", HTTPS_PROTOCOL, GITHUB_DOMAIN, username, repo)
}

func RepositoryBranchUrl(username string, repo string, branch string) string {
	return fmt.Sprintf("%s://%s/%s/%s/%s", HTTPS_PROTOCOL, GITHUB_DOMAIN, username, repo, branch)
}
