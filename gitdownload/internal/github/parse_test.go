package download

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGithubUrl(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/blob/main/test/main.go"
	githubRepository, err := NewGitRepositoryByUrl(url)
	log.Printf("%+v\n", githubRepository)
	assert.Equal(t, nil, err)
}

func TestMatch(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/blob/main/test/main.go"
	matches, err := match(url)
	log.Printf("%+v\n", matches)
	assert.Equal(t, nil, err)
}
