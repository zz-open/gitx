package github

type FileChunk struct {
	GitBlobUrl  string `json:"git_blob_url"`
	Path        string `json:"path"`
	DownloadUrl string `json:"download_url"`
}

func NewFileChunk(gitBlobUrl string, path string, downloadUrl string) *FileChunk {
	chunk := &FileChunk{
		GitBlobUrl:  gitBlobUrl,
		Path:        path,
		DownloadUrl: downloadUrl,
	}

	return chunk
}
