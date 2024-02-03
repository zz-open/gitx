package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootUrl(t *testing.T) {
	username := "zzopen"
	repo := "mysqldoc"
	assert.Equal(t, "https://github.com/zzopen/mysqldoc", RepositoryRootUrl(username, repo))
}

func TestBranchUrl(t *testing.T) {
	username := "zzopen"
	repo := "mysqldoc"
	branch := "main"
	assert.Equal(t, "https://github.com/zzopen/mysqldoc/main", RepositoryBranchUrl(username, repo, branch))
}

func TestContentApiUrl(t *testing.T) {
	username := "zzopen"
	repo := "mysqldoc"
	branch := "main"
	path := "src/common/query/query.go"
	assert.Equal(t, "https://api.github.com/repos/zzopen/mysqldoc/contents/src/common/query/query.go?ref=main", ContentApiUrl(username, repo, branch, path))
}

func TestRawUserContentUrl(t *testing.T) {
	username := "zzopen"
	repo := "mysqldoc"
	branch := "main"
	path := "src/common/query/query.go"
	assert.Equal(t, "https://raw.githubusercontent.com/zzopen/mysqldoc/main/src/common/query/query.go", RawUserContentUrl(username, repo, branch, path))
}

func TestGitTreesApiUrl(t *testing.T) {
	username := "zzopen"
	repo := "mysqldoc"
	branch := "main"
	assert.Equal(t, "https://api.github.com/repos/zzopen/mysqldoc/git/trees/main?recursive=1", GitTreesApiUrl(username, repo, branch, true))
	assert.Equal(t, "https://api.github.com/repos/zzopen/mysqldoc/git/trees/main", GitTreesApiUrl(username, repo, branch, false))
}
