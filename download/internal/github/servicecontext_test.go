package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServiceContext(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	outpath := "../../../_test_"
	token := ""

	svc, err := NewServiceContext(
		url,
		ServiceContextWithOutpath(outpath),
		ServiceContextWithToken(token),
	)
	assert.Equal(t, nil, err)

	assert.NotEqual(t, nil, svc)
	assert.NotEqual(t, nil, svc.Repository)
	assert.NotEqual(t, nil, svc.HttpClient)
}
