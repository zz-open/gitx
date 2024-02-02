package github

import (
	"fmt"
)

type Repository struct {
	Protocol    string `json:"protocol"`
	Host        string `json:"host"`
	Username    string `json:"username"`
	ProjectName string `json:"project_name"`
	Branch      string `json:"branch"`
	Type        string `json:"type"`
	Path        string `json:"path"`
	FileName    string `json:"file_name"`
	Token       string `json:"token"`
}

func (repo *Repository) IsFile() bool {
	return repo.Type == "blob"
}

func (repo *Repository) IsDir() bool {
	return repo.Type == "tree"
}

func (repo *Repository) RootUrl() string {
	return fmt.Sprintf("%s://%s/%s/%s", repo.Protocol, repo.Host, repo.Username, repo.ProjectName)
}

func (repo *Repository) BranchUrl() string {
	return fmt.Sprintf("%s://%s/%s/%s/%s", repo.Protocol, repo.Host, repo.Username, repo.ProjectName, repo.Branch)
}

func (repo *Repository) GitTreesApiUrl() string {
	return fmt.Sprintf("%s://%s/repos/%s/%s/git/trees/%s?recursive=1", repo.Protocol, GITHUB_API_URL, repo.Username, repo.ProjectName, repo.Branch)
}

func (repo *Repository) RawUserContentUrl() string {
	return fmt.Sprintf("%s://%s/%s/%s/%s/%s", repo.Protocol, GITHUB_RAW_USER_CONTENT_URL, repo.Username, repo.ProjectName, repo.Branch, repo.Path)
}

func (repo *Repository) ContentApiUrl() string {
	return fmt.Sprintf("%s://%s/repos/%s/%s/contents/%s?ref=%s", repo.Protocol, GITHUB_API_URL, repo.Username, repo.ProjectName, repo.Path, repo.Branch)
}
