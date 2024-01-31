package download

import "fmt"

const HTTP_PROTOCOL = "http"
const HTTPS_PROTOCOL = "https"

const GITHUB_DOMAIN = "github.com"
const GITHUB_API_URL = "api.github.com"
const GITHUB_RAW_USER_CONTENT_URL = "raw.githubusercontent.com"

func GithubReoisitoryRegexp() string {
	return fmt.Sprintf("^%s://%s/([^/]+)/([^/]+)(/(tree|blob)/([^/]+)(/(.*))?)?", HTTPS_PROTOCOL, GITHUB_DOMAIN)
}
