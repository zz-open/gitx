package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServiceContext(t *testing.T) {
	outpath := ""
	token := ""
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"

	svc, err := NewServiceContext(
		url,
		ServiceContextWithOutpath(outpath),
		ServiceContextWithToken(token),
	)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, svc)
}
