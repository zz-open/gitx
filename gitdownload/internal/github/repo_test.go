package download

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootUrl(t *testing.T) {
	githubRepo := &GitHubRepository{
		Protocol:    "https",
		Host:        "github.com",
		Username:    "zzopen",
		ProjectName: "mysqldoc",
		Branch:      "main",
		Type:        "blob",
		Path:        "src/internal/cmd/root.go",
	}

	assert.Equal(t, "https://github.com/zzopen/mysqldoc", githubRepo.RootUrl())
}

func TestBranchUrl(t *testing.T) {
	githubRepo := &GitHubRepository{
		Protocol:    "https",
		Host:        "github.com",
		Username:    "zzopen",
		ProjectName: "mysqldoc",
		Branch:      "main",
		Type:        "blob",
		Path:        "src/internal/cmd/root.go",
	}

	assert.Equal(t, "https://github.com/zzopen/mysqldoc/main", githubRepo.BranchUrl())
}

func TestGitTreesApiUrl(t *testing.T) {
	githubRepo := &GitHubRepository{
		Protocol:    "https",
		Host:        "github.com",
		Username:    "zzopen",
		ProjectName: "mysqldoc",
		Branch:      "main",
		Type:        "blob",
		Path:        "src/internal/cmd/root.go",
	}

	assert.Equal(t, "https://api.github.com/repos/zzopen/mysqldoc/git/trees/main?recursive=1", githubRepo.GitTreesApiUrl())
}

func TestRawUserContentUrl(t *testing.T) {
	githubRepo := &GitHubRepository{
		Protocol:    "https",
		Host:        "github.com",
		Username:    "zzopen",
		ProjectName: "mysqldoc",
		Branch:      "main",
		Type:        "blob",
		Path:        "src/internal/cmd/root.go",
	}

	assert.Equal(t, "https://raw.githubusercontent.com/zzopen/mysqldoc/main/src/internal/cmd/root.go", githubRepo.RawUserContentUrl())
}
