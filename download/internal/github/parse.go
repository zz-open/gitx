package github

import (
	"errors"
	"path/filepath"
	"regexp"

	"github.com/zz-open/gitx/common"
)

func UrlParseToRepository(url string) (*Repository, error) {
	if url == "" {
		return nil, errors.New("url 必填")
	}

	matches, err := match(url)
	if err != nil {
		return nil, err
	}

	path := ""
	if matches[7] != "" {
		path = common.FilterTailSlash(matches[7])
	}

	repo := &Repository{
		Username: matches[1],
		Repo:     matches[2],
		Branch:   matches[5],
		Type:     matches[4],
		Path:     path,
	}

	if repo.IsDir() {
		repo.Dirname = filepath.Base(path)
	} else if repo.IsFile() {
		repo.Filename = filepath.Base(path)
	}

	return repo, nil
}

func match(url string) ([]string, error) {
	re := regexp.MustCompile(ReoisitoryRegexp())
	if !re.MatchString(url) {
		return nil, errors.New("请输入正确的git仓库地址")
	}

	matches := re.FindStringSubmatch(url)
	if len(matches) <= 0 {
		return nil, errors.New("无法获取匹配项")
	}

	return matches, nil
}
