package github

import (
	"errors"
	"path"
	"path/filepath"
	"regexp"

	"github.com/zz-open/zb/common"
)

type Repository struct {
	Username string `json:"username"`
	Repo     string `json:"repo"`
	Branch   string `json:"branch"`
	Type     string `json:"type"`
	Path     string `json:"path"`
	Filename string `json:"filename"`
	Dirname  string `json:"dirname"`
}

func (repo *Repository) LastLevelDirname() string {
	if repo.IsTree() {
		return path.Base(repo.Path)
	} else if repo.IsBlob() {
		return path.Base(path.Dir(repo.Path))
	}

	return ""
}

func (repo *Repository) IsBlob() bool {
	return repo.Type == "blob"
}

func (repo *Repository) IsTree() bool {
	return repo.Type == "tree"
}

func (repo *Repository) IsRoot() bool {
	return repo.Path == ""
}

func (repo *Repository) RootUrl() string {
	return RepositoryRootUrl(repo.Username, repo.Repo)
}

func (repo *Repository) BranchUrl() string {
	return RepositoryBranchUrl(repo.Username, repo.Repo, repo.Branch)
}

func (repo *Repository) ZipballUrl() string {
	return ZipballUrl(repo.Username, repo.Repo, repo.Branch)
}

func (repo *Repository) TarballUrl() string {
	return TarballUrl(repo.Username, repo.Repo, repo.Branch)
}

func (repo *Repository) ArchiveZipUrl() string {
	return ArchiveZipUrl(repo.Username, repo.Repo, repo.Branch)
}

func (repo *Repository) GitTreesApiUrl(isRecursive bool) string {
	return GitTreesApiUrl(repo.Username, repo.Repo, repo.Branch, isRecursive)
}

func (repo *Repository) RawUserContentUrl() string {
	return RawUserContentUrl(repo.Username, repo.Repo, repo.Branch, repo.Path)
}

func (repo *Repository) ContentApiUrl() string {
	return ContentApiUrl(repo.Username, repo.Repo, repo.Branch, repo.Path)
}

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

	if repo.IsTree() {
		repo.Dirname = filepath.Base(path)
	} else if repo.IsBlob() {
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
