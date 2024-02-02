package directory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/tree/main/cli"
	outpath := "./"
	err := Download(url, outpath)
	assert.Equal(t, nil, err)
}
