package github

type Repository struct {
	Username string `json:"username"`
	Repo     string `json:"repo"`
	Branch   string `json:"branch"`
	Type     string `json:"type"`
	Path     string `json:"path"`
	Filename string `json:"filename"`
	Dirname  string `json:"dirname"`
}

func (repo *Repository) IsFile() bool {
	return repo.Type == "blob"
}

func (repo *Repository) IsDir() bool {
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
