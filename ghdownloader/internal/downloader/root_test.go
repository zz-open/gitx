package downloader

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zz-open/zbin/ghdownloader/internal/github"
)

func TestRootDownloadZip(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/tree/main"
	outpath := "../../../_test_"
	token := ""

	svc, err := github.NewServiceContext(
		url,
		github.ServiceContextWithOutpath(outpath),
		github.ServiceContextWithToken(token),
	)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, svc)

	var dl Downloader = NewRootDownloader(svc, "zip")
	err = dl.Download()
	assert.Equal(t, nil, err)
}

func TestRootDownloadTar(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/tree/main"
	outpath := "../../../_test_"
	token := ""

	svc, err := github.NewServiceContext(
		url,
		github.ServiceContextWithOutpath(outpath),
		github.ServiceContextWithToken(token),
	)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, svc)

	var dl Downloader = NewRootDownloader(svc, "tar")
	err = dl.Download()
	assert.Equal(t, nil, err)
}
