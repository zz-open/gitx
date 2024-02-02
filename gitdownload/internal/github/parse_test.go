package github

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGithubUrl(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/blob/main/test/main.go"
	repository, err := NewRepositoryByUrl(url)
	log.Printf("%+v\n", repository)
	assert.Equal(t, nil, err)
}

func TestMatch(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/blob/main/test/main.go"
	matches, err := match(url)
	log.Printf("%+v\n", matches)
	assert.Equal(t, nil, err)
}
