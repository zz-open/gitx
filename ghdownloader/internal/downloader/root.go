package downloader

import (
	"errors"
	"path/filepath"

	"github.com/zz-open/zbin/common"
	"github.com/zz-open/zbin/ghdownloader/internal/github"
)

func NewRootDownloader(svc *github.ServiceContext, archiveType string) *RootDownloader {
	rd := &RootDownloader{
		ServiceContext: svc,
		ArchiveType:    archiveType,
	}

	return rd
}

type RootDownloader struct {
	ServiceContext *github.ServiceContext
	ArchiveType    string
}

func (rd *RootDownloader) IsZip() bool {
	return rd.ArchiveType == "zip"
}

func (rd *RootDownloader) IsTar() bool {
	return rd.ArchiveType == "tar"
}

func (rd *RootDownloader) Download() error {
	var filename string
	var b []byte
	var err error

	httpClient := rd.ServiceContext.HttpClient

	if rd.IsZip() {
		filename, b, err = httpClient.SendZipballRequest(rd.ServiceContext.Repository.ZipballUrl())
	} else if rd.IsTar() {
		filename, b, err = httpClient.SendTarballRequest(rd.ServiceContext.Repository.TarballUrl())
	}

	if err != nil {
		return err
	}

	if b == nil || len(b) == 0 {
		return errors.New("归档文件内容为空")
	}

	err = common.WriteFile(b, filepath.Join(rd.ServiceContext.Outpath, filename))
	if err != nil {
		return err
	}

	return nil
}
