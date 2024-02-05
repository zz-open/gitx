package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/zz-open/zb/common"
)

type RequestFailResponse struct {
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}

type HttpClient struct {
	Token      string
	Accept     string
	ApiVersion string
}

type HttpClientOption func(c *HttpClient)

func HttpClientWithToken(token string) HttpClientOption {
	return func(c *HttpClient) {
		c.Token = token
	}
}

func HttpClientWithAccept(accept string) HttpClientOption {
	return func(c *HttpClient) {
		c.Accept = accept
	}
}

func HttpClientWithApiVersion(apiVersion string) HttpClientOption {
	return func(c *HttpClient) {
		c.ApiVersion = apiVersion
	}
}

func NewHttpClient(opts ...HttpClientOption) *HttpClient {
	c := &HttpClient{
		Accept:     "application/vnd.github+json",
		ApiVersion: "2022-11-28",
	}

	for _, op := range opts {
		op(c)
	}

	return c
}

func (c *HttpClient) GetHeaders() map[string]string {
	headers := make(map[string]string)
	if c.Token != "" {
		headers["Authorization"] = fmt.Sprintf("Bearer %s", c.Token)
	}

	if c.Accept != "" {
		headers["Accept"] = c.Accept
	}

	if c.ApiVersion != "" {
		headers["X-GitHub-Api-Version"] = c.ApiVersion
	}

	return headers
}

func (c *HttpClient) SendGet(url string) (*http.Response, error) {
	return common.HttpGet(url, c.GetHeaders())
}

func (c *HttpClient) SendGetAndRead(url string) ([]byte, error) {
	response, err := c.SendGet(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	res, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		fail := &RequestFailResponse{}
		err = json.Unmarshal(res, fail)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(fail.Message)
	}

	return res, nil
}

func (c *HttpClient) SendFileReposContentRequest(url string) (*ReposContents, error) {
	res, err := c.SendGetAndRead(url)
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

func (c *HttpClient) SendDirReposContentRequest(url string) ([]*ReposContents, error) {
	res, err := c.SendGetAndRead(url)
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

func (c *HttpClient) SendReposGitTreesRequest(url string) (*ReposGitTrees, error) {
	res, err := c.SendGetAndRead(url)
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

func (c *HttpClient) SendReposGitBlobsRequest(url string) (*ReposGitBlobs, error) {
	res, err := c.SendGetAndRead(url)
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

func (c *HttpClient) SendZipballRequest(url string) (string, []byte, error) {
	resp, err := c.SendGet(url)
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

func (c *HttpClient) SendTarballRequest(url string) (string, []byte, error) {
	resp, err := c.SendGet(url)
	if err != nil {
		return "", nil, err
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/x-gzip" {
		return "", nil, errors.New("Content-Type is not application/x-gzip")
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

func (c *HttpClient) SendRawGithubUserContentRequest(url string) ([]byte, error) {
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
