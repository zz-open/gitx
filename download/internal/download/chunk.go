package download

type FileChunk struct {
	GitUrl string `json:"git_url"`
	Path   string `json:"path"`
}

func NewFileChunk(gitUrl string, path string) *FileChunk {
	chunk := &FileChunk{
		GitUrl: gitUrl,
		Path:   path,
	}

	return chunk
}
