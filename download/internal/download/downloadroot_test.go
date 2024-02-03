package download

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zz-open/gitx/download/internal/github"
)

func TestDownloadRootZip(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/tree/main"
	outpath := "./"
	token := ""

	svc, err := github.NewServiceContext(
		url,
		github.ServiceContextWithOutpath(outpath),
		github.ServiceContextWithToken(token),
	)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, svc)

	err = DownloadRootZip(svc)
	assert.Equal(t, nil, err)
}

func TestDownloadRootGZip(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/tree/main"
	outpath := "./"
	token := ""

	svc, err := github.NewServiceContext(
		url,
		github.ServiceContextWithOutpath(outpath),
		github.ServiceContextWithToken(token),
	)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, svc)

	err = DownloadRootGZip(svc)
	assert.Equal(t, nil, err)
}
