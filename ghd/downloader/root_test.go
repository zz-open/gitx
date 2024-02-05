package downloader

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zz-open/zb/ghd/sc"
)

func TestRootDownloadZip(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/tree/main"
	outpath := TEST_OUT_PATH
	token := ""

	svc, err := sc.NewServiceContext(
		url,
		sc.ServiceContextWithOutpath(outpath),
		sc.ServiceContextWithToken(token),
	)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, svc)

	var dl Downloader = NewRootDownloader(svc, "zip")
	err = dl.Download()
	assert.Equal(t, nil, err)
}

func TestRootDownloadTar(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/tree/main"
	outpath := TEST_OUT_PATH
	token := ""

	svc, err := sc.NewServiceContext(
		url,
		sc.ServiceContextWithOutpath(outpath),
		sc.ServiceContextWithToken(token),
	)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, svc)

	var dl Downloader = NewRootDownloader(svc, "tar")
	err = dl.Download()
	assert.Equal(t, nil, err)
}
