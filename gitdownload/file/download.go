package file

import (
	"encoding/base64"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/zz-open/gitx/common"
	"github.com/zz-open/gitx/gitdownload/internal/github"
)

func Download(url string, outpath string) error {
	var err error
	err = checkParameter(url, outpath)
	if err != nil {
		return err
	}

	repository, err := github.NewRepositoryByUrl(url)
	if err != nil {
		return err
	}

	content, err := fetchFileContent(repository)
	if err != nil {
		return err
	}

	outAbsPath, err := filepath.Abs(outpath)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", outAbsPath, repository.FileName)
	err = common.WriteFile(filePath, content)
	if err != nil {
		return err
	}

	return nil
}

// fetchFileContent 请求github reset api
func fetchFileContent(repository *github.Repository) ([]byte, error) {
	resp, err := github.RequestReposContent(repository.ContentApiUrl(), repository.Token)
	if err != nil {
		return nil, err
	}

	if resp.Type != "file" {
		return nil, errors.New("api response 'type' is not a file")
	}

	var fileContent []byte
	if resp.Encoding == "base64" {
		// base64 直接解码
		fileContent, err = base64.StdEncoding.DecodeString(resp.Content)
	} else {
		// 非base64 通过 raw.githubusercontent.com 下载
		fileContent, err = github.RequestRawGithubUserContent(resp.DownloadUrl)
	}

	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

// checkParameter 检查入参
func checkParameter(url string, path string) error {
	if url == "" {
		return errors.New("url 不合法")
	}

	if path == "" {
		return errors.New("path 不合法")
	}

	return nil
}
