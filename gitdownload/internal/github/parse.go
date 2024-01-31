package download

import (
	"errors"
	"regexp"
)

func NewGitRepositoryByUrl(url string) (*GitHubRepository, error) {
	return parseGithubUrl(url)
}

func parseGithubUrl(url string) (*GitHubRepository, error) {
	if url == "" {
		return nil, errors.New("url 必填")
	}

	matches, err := match(url)
	if err != nil {
		return nil, err
	}

	path := ""
	if matches[7] != "" {
		path = FilterTailSlash(matches[7])
	}

	repo := &GitHubRepository{
		Protocol:    HTTPS_PROTOCOL,
		Host:        GITHUB_DOMAIN,
		Username:    matches[1],
		ProjectName: matches[2],
		Branch:      matches[5],
		Type:        matches[4],
		Path:        path,
	}

	return repo, nil
}

func match(url string) ([]string, error) {
	re := regexp.MustCompile(GithubReoisitoryRegexp())
	if !re.MatchString(url) {
		return nil, errors.New("请输入正确的git仓库地址")
	}

	matches := re.FindStringSubmatch(url)
	if len(matches) <= 0 {
		return nil, errors.New("无法获取匹配项")
	}

	return matches, nil
}
