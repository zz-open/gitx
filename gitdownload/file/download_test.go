package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/blob/main/cli/common.mk"
	outpath := "./"
	err := Download(url, outpath)
	assert.Equal(t, nil, err)
}

func TestFetchFileContent(t *testing.T) {
}

func TestWriteFile(t *testing.T) {

}

func TestCheckParameter(t *testing.T) {
}
