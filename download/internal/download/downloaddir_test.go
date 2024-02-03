package download

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zz-open/gitx/download/internal/github"
)

func TestDownloadDir(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/tree/main/src/common"
	outpath := "./"
	token := ""
	svc, err := github.NewServiceContext(
		url,
		github.ServiceContextWithOutpath(outpath),
		github.ServiceContextWithToken(token),
	)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, svc)

	err = DownloadDir(svc)
	assert.Equal(t, nil, err)
}
