package github

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/zz-open/gitx/common"
)

type ApiRequestFail struct {
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}

func request(url string, token string) (*http.Response, error) {
	headers := make(map[string]string)
	headers["Accept"] = "application/vnd.github+json"
	if token != "" {
		headers["Authorization: Bearer"] = token
	}

	headers["X-GitHub-Api-Version"] = "2022-11-28"
	return common.HttpGet(url, headers)
}

func requestAndReadBody(url string, token string) ([]byte, error) {
	response, err := request(url, token)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	res, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		fail := &ApiRequestFail{}
		err = json.Unmarshal(res, fail)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(fail.Message)
	}

	return res, nil
}

func RequestReposContentWithFilePath(url string, token string) (*ReposContents, error) {
	res, err := requestAndReadBody(url, token)
	if err != nil {
		return nil, err
	}

	data := &ReposContents{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func RequestReposContentWithDirPath(url string, token string) ([]*ReposContents, error) {
	res, err := requestAndReadBody(url, token)
	if err != nil {
		return nil, err
	}

	data := make([]*ReposContents, 0)
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func RequestRawGithubUserContent(url string) ([]byte, error) {
	response, err := common.HttpGet(url, nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	res, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func RequestReposGitTrees(url string, token string) (*ReposGitTrees, error) {
	res, err := requestAndReadBody(url, token)
	if err != nil {
		return nil, err
	}

	data := &ReposGitTrees{
		Tree: make([]ReposGitTreesItem, 0),
	}

	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}

	if data.Truncated {
		return nil, errors.New("超出API限制")
	}

	return data, nil
}

func RequestReposGitBlobs(url string, token string) (*ReposGitBlobs, error) {
	res, err := requestAndReadBody(url, token)
	if err != nil {
		return nil, err
	}

	data := &ReposGitBlobs{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func RequestZipball(url string, token string) (string, []byte, error) {
	resp, err := request(url, token)
	if err != nil {
		return "", nil, err
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/zip" {
		return "", nil, errors.New("Content-Type is not application/zip")
	}

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	filename, err := common.ParseAttachmentFilename(resp.Header.Get("content-disposition"))
	if err != nil {
		return "", nil, err
	}

	return filename, b, nil
}

func RequestTarball(url string, token string) (string, []byte, error) {
	resp, err := request(url, token)
	if err != nil {
		return "", nil, err
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/x-gzip" {
		return "", nil, errors.New("Content-Type is not application/x-gzip")
	}

	// fmt.Println(resp.Header.Get("content-disposition")) // attachment; filename=zzopen-mysqldoc-e3cbfdf.tar.gz
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	filename, err := common.ParseAttachmentFilename(resp.Header.Get("content-disposition"))
	if err != nil {
		return "", nil, err
	}

	return filename, b, nil
}
