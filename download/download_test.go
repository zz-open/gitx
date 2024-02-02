package downlaod

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zz-open/gitx/download/internal/github"
)

func TestDownload(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	outpath := "./"
	token := ""
	err := Download(url, outpath, token)
	assert.Equal(t, nil, err)
}

func TestSearchTargetTreeItem(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	outpath := "./"
	token := ""
	svc, err := github.NewServiceContext(
		url,
		github.ServiceContextWithOutpath(outpath),
		github.ServiceContextWithToken(token),
	)

	assert.NotEqual(t, nil, svc)
	assert.Equal(t, nil, err)

	treeItem, err := searchTargetTreeItem(svc)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, treeItem)

	log.Printf("treeItem: %+v\n", treeItem)
}
