package download

import "fmt"

type GitHubRepository struct {
	Protocol    string `json:"protocol"`
	Host        string `json:"host"`
	Username    string `json:"username"`
	ProjectName string `json:"project_name"`
	Branch      string `json:"branch"`
	Type        string `json:"type"`
	Path        string `json:"path"`
}

func (gr *GitHubRepository) RootUrl() string {
	return fmt.Sprintf("%s://%s/%s/%s", gr.Protocol, gr.Host, gr.Username, gr.ProjectName)
}

func (gr *GitHubRepository) BranchUrl() string {
	return fmt.Sprintf("%s://%s/%s/%s/%s", gr.Protocol, gr.Host, gr.Username, gr.ProjectName, gr.Branch)
}

func (gr *GitHubRepository) GitTreesApiUrl() string {
	return fmt.Sprintf("%s://%s/repos/%s/%s/git/trees/%s?recursive=1", gr.Protocol, GITHUB_API_URL, gr.Username, gr.ProjectName, gr.Branch)
}

func (gr *GitHubRepository) RawUserContentUrl() string {
	return fmt.Sprintf("%s://%s/%s/%s/%s/%s", gr.Protocol, GITHUB_RAW_USER_CONTENT_URL, gr.Username, gr.ProjectName, gr.Branch, gr.Path)
}
