package github

type ReposContentResponse struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	Url         string `json:"url"`
	DownloadUrl string `json:"download_url"`
	GitUrl      string `json:"git_url"`
	Type        string `json:"type"`
	Encoding    string `json:"encoding"`
	Content     string `json:"content"`
}

type ReposGitTreeResponse struct {
}
