package downlaod

import (
	"errors"
	"fmt"

	"github.com/zz-open/gitx/download/internal/downloader"
	"github.com/zz-open/gitx/download/internal/github"
)

func GithubDownload(url string, outpath string, token string) error {
	svc, err := github.NewServiceContext(
		url,
		github.ServiceContextWithOutpath(outpath),
		github.ServiceContextWithToken(token),
	)

	if err != nil {
		return err
	}

	successTip := ""

	var dl downloader.Downloader
	if svc.Repository.IsRoot() {
		dl = downloader.NewRootDownloader(svc, "zip")
		successTip = "仓库下载成功"
	}

	if svc.Repository.IsBlob() {
		dl = downloader.NewBlobDownloader(svc)
		successTip = "文件下载成功"
	}

	if svc.Repository.IsTree() {
		dl = downloader.NewTreeDownloader(svc)
		successTip = "目录下载成功"
	}

	if dl == nil {
		return errors.New("类型错误, 请检查url")
	}

	err = dl.Download()
	if err != nil {
		return err
	}

	fmt.Println(successTip)
	return nil
}
