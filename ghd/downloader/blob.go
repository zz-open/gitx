package downloader

import (
	"path/filepath"

	"github.com/zz-open/zbin/common"

	"github.com/zz-open/zbin/ghd/sc"
)

func NewBlobDownloader(svc *sc.ServiceContext) *BlobDownloader {
	bd := &BlobDownloader{
		ServiceContext: svc,
	}

	return bd
}

type BlobDownloader struct {
	ServiceContext *sc.ServiceContext
}

func (bd *BlobDownloader) Download() error {
	return bd.download()
}

func (bd *BlobDownloader) download() error {
	svc := bd.ServiceContext
	repository := svc.Repository
	httpClient := svc.HttpClient

	v, err := httpClient.SendRawGithubUserContentRequest(repository.RawUserContentUrl())
	if err != nil {
		return err

	}

	localPath := filepath.Join(svc.Outpath, repository.Filename)
	err = common.WriteFile(v, localPath)
	if err != nil {
		return err
	}

	return nil
}

func (bd *BlobDownloader) download2() error {
	svc := bd.ServiceContext
	repository := svc.Repository
	httpClient := svc.HttpClient

	v, err := httpClient.SendFileReposContentRequest(repository.ContentApiUrl())
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

	return nil
}
