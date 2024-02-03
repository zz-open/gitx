package download

import (
	"fmt"
	"path/filepath"

	"github.com/zz-open/gitx/common"
	"github.com/zz-open/gitx/download/internal/github"
)

func DownloadFile(svc *github.ServiceContext) error {
	repository := svc.Repository
	v, err := github.RequestRawGithubUserContent(repository.RawUserContentUrl())
	if err != nil {
		return err
	}

	localPath := filepath.Join(svc.Outpath, repository.Filename)
	err = common.WriteFile(v, localPath)
	if err != nil {
		return err
	}

	fmt.Printf("文件下载成功\n")
	return nil
}

func DownloadFile2(svc *github.ServiceContext) error {
	repository := svc.Repository
	v, err := github.RequestReposContentWithFilePath(repository.ContentApiUrl(), svc.Token)
	if err != nil {
		return err
	}

	fileContent, err := common.DecodeGithubContent(v.Encoding, v.Content)
	if err != nil {
		return err
	}

	err = common.WriteFile(fileContent, filepath.Join(svc.Outpath, repository.Filename))
	if err != nil {
		return err
	}

	fmt.Printf("文件下载成功\n")
	return nil
}
