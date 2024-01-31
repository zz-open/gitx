package download

import "strings"

func FilterTailSlash(str string) string {
	return strings.TrimSuffix("/", str)
}

func IsGitUrl(url string) bool {
	return true
}
