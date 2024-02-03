package download

import (
	"fmt"
	"path/filepath"

	"github.com/zz-open/gitx/common"
	"github.com/zz-open/gitx/download/internal/github"
)

func DownloadRootZip(svc *github.ServiceContext) error {
	filename, v, err := github.RequestZipball(svc.Repository.ZipballUrl(), svc.Token)
	if err != nil {
		return err
	}

	err = common.WriteFile(v, filepath.Join(svc.Outpath, filename))
	if err != nil {
		return err
	}

	fmt.Printf("zip文件下载成功\n")
	return nil
}

func DownloadRootGZip(svc *github.ServiceContext) error {
	filename, v, err := github.RequestTarball(svc.Repository.TarballUrl(), svc.Token)
	if err != nil {
		return err
	}

	err = common.WriteFile(v, filepath.Join(svc.Outpath, filename))
	if err != nil {
		return err
	}

	fmt.Printf("gzip文件下载成功\n")
	return nil
}
