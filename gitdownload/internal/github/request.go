package github

import (
	"encoding/json"

	"github.com/zz-open/gitx/common"
)

func RequestReposContent(url string, token string) (*ReposContentResponse, error) {
	headers := make(map[string]string)
	headers["Accept"] = "application/vnd.github+json"
	if token != "" {
		headers["Authorization: Bearer"] = token
	}

	headers["X-GitHub-Api-Version"] = "2022-11-28"
	res, err := common.HttpGet(url, headers)
	if err != nil {
		return nil, err
	}

	resp := &ReposContentResponse{}
	err = json.Unmarshal(res, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func RequestRawGithubUserContent(url string) ([]byte, error) {
	return common.HttpGet(url, nil)
}

func RequestReposGitTree(url string, token string) (*ReposGitTreeResponse, error) {
	headers := make(map[string]string)
	headers["Accept"] = "application/vnd.github+json"
	if token != "" {
		headers["Authorization: Bearer"] = token
	}

	headers["X-GitHub-Api-Version"] = "2022-11-28"
	res, err := common.HttpGet(url, headers)
	if err != nil {
		return nil, err
	}

	resp := &ReposGitTreeResponse{}
	err = json.Unmarshal(res, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
