package github

import (
	"errors"
	"path/filepath"
)

type ServiceContext struct {
	Url        string
	Outpath    string
	Token      string
	Repository *Repository
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

	repository, err := UrlParseToRepository(url, svc.Token)
	if err != nil {
		return nil, err
	}

	svc.Repository = repository
	return svc, nil
}
