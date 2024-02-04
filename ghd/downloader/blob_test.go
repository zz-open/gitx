package downloader

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zz-open/zbin/ghd/sc"
)

func TestBlobDownloader(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	outpath := "../../../_test_"
	token := ""

	svc, err := sc.NewServiceContext(
		url,
		sc.ServiceContextWithOutpath(outpath),
		sc.ServiceContextWithToken(token),
	)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, svc)

	dl := NewBlobDownloader(svc)
	err = dl.Download()
	assert.Equal(t, nil, err)
}

func TestBlobDownloader2(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	outpath := "../../../_test_"
	token := ""

	svc, err := sc.NewServiceContext(
		url,
		sc.ServiceContextWithOutpath(outpath),
		sc.ServiceContextWithToken(token),
	)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, svc)

	dl := NewBlobDownloader(svc)
	err = dl.download2()
	assert.Equal(t, nil, err)
}