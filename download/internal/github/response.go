package github

// https://api.github.com/repos/(username)/(repository)/git/trees/(SHA1) =============

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

// https://api.github.com/repos/(username)/(repository)/git/trees/(SHA1) =============

type ReposGitTrees struct {
	Sha       string              `json:"sha"`
	Url       string              `json:"url"`
	Truncated bool                `json:"truncated"`
	Tree      []ReposGitTreesItem `json:"tree"`
}

type ReposGitTreesItem struct {
	Path           string `json:"path"`
	Sha            string `json:"sha"`
	Type           string `json:"type"`
	Url            string `json:"url"`
	PathChain      string `json:"-"`
	RawDownloadUrl string `json:"-"`
}

func (item *ReposGitTreesItem) IsFile() bool {
	return item.Type == "blob"
}

func (item *ReposGitTreesItem) IsDir() bool {
	return item.Type == "tree"
}

// https://api.github.com/repos/(username)/(repository)/git/blobs/(SHA1) =============

type ReposGitBlobs struct {
	Sha      string `json:"sha"`
	Size     int    `json:"size"`
	NodeId   string `json:"node_id"`
	Url      string `json:"url"`
	Encoding string `json:"encoding"`
	Content  string `json:"content"`
}
