package sc

import (
	"errors"
	"path/filepath"

	"github.com/zz-open/zb/ghd/github"
	"github.com/zz-open/zb/ghd/http"
)

type ServiceContext struct {
	Token string

	Url        string
	Outpath    string
	Repository *github.Repository
	HttpClient *http.HttpClient
}

type ServiceContextOption func(*ServiceContext)

func ServiceContextWithOutpath(outpath string) ServiceContextOption {
	return func(svc *ServiceContext) {
		svc.Outpath = outpath
	}
}

func ServiceContextWithToken(token string) ServiceContextOption {
	return func(svc *ServiceContext) {
		svc.Token = token
	}
}

func NewServiceContext(url string, opts ...ServiceContextOption) (*ServiceContext, error) {
	svc := &ServiceContext{
		Url: url,
	}

	for _, option := range opts {
		option(svc)
	}

	if svc.Url == "" {
		return nil, errors.New("url 不合法")
	}

	if svc.Outpath == "" {
		svc.Outpath = "."
	}

	absPath, err := filepath.Abs(svc.Outpath)
	if err != nil {
		return nil, err
	}

	svc.Outpath = absPath

	repository, err := github.UrlParseToRepository(url)
	if err != nil {
		return nil, err
	}

	svc.Repository = repository

	httpClient := http.NewHttpClient(http.HttpClientWithToken(svc.Token))
	svc.HttpClient = httpClient

	return svc, nil
}
