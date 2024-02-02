package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootUrl(t *testing.T) {
	repository := &Repository{
		Protocol:    "https",
		Host:        "github.com",
		Username:    "zzopen",
		ProjectName: "mysqldoc",
		Branch:      "main",
		Type:        "blob",
		Path:        "src/internal/cmd/root.go",
	}

	assert.Equal(t, "https://github.com/zzopen/mysqldoc", repository.RootUrl())
}

func TestBranchUrl(t *testing.T) {
	repository := &Repository{
		Protocol:    "https",
		Host:        "github.com",
		Username:    "zzopen",
		ProjectName: "mysqldoc",
		Branch:      "main",
		Type:        "blob",
		Path:        "src/internal/cmd/root.go",
	}

	assert.Equal(t, "https://github.com/zzopen/mysqldoc/main", repository.BranchUrl())
}

func TestGitTreesApiUrl(t *testing.T) {
	repository := &Repository{
		Protocol:    "https",
		Host:        "github.com",
		Username:    "zzopen",
		ProjectName: "mysqldoc",
		Branch:      "main",
		Type:        "blob",
		Path:        "src/internal/cmd/root.go",
	}

	assert.Equal(t, "https://api.github.com/repos/zzopen/mysqldoc/git/trees/main?recursive=1", repository.GitTreesApiUrl())
}

func TestRawUserContentUrl(t *testing.T) {
	repository := &Repository{
		Protocol:    "https",
		Host:        "github.com",
		Username:    "zzopen",
		ProjectName: "mysqldoc",
		Branch:      "main",
		Type:        "blob",
		Path:        "src/internal/cmd/root.go",
	}

	assert.Equal(t, "https://raw.githubusercontent.com/zzopen/mysqldoc/main/src/internal/cmd/root.go", repository.RawUserContentUrl())
}

func TestContentApiUrl(t *testing.T) {
	repository := &Repository{
		Protocol:    "https",
		Host:        "github.com",
		Username:    "zzopen",
		ProjectName: "mysqldoc",
		Branch:      "main",
		Type:        "blob",
		Path:        "src/internal/cmd/root.go",
	}

	assert.Equal(t, "https://api.github.com/repos/zzopen/mysqldoc/contents/src/internal/cmd/root.go?ref=main", repository.ContentApiUrl())
}
