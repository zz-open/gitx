package download

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zz-open/gitx/download/internal/github"
)

func TestDownloadFile(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	outpath := "./"
	token := ""

	svc, err := github.NewServiceContext(
		url,
		github.ServiceContextWithOutpath(outpath),
		github.ServiceContextWithToken(token),
	)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, svc)

	err = DownloadFile(svc)
	assert.Equal(t, nil, err)
}

func TestDownloadFile2(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	outpath := "./"
	token := ""

	svc, err := github.NewServiceContext(
		url,
		github.ServiceContextWithOutpath(outpath),
		github.ServiceContextWithToken(token),
	)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, svc)

	err = DownloadFile2(svc)
	assert.Equal(t, nil, err)
}
