package downlaod

import (
	down "github.com/zz-open/gitx/download/internal/download"
	"github.com/zz-open/gitx/download/internal/github"
)

func Download(url string, outpath string, token string) error {
	svc, err := github.NewServiceContext(
		url,
		github.ServiceContextWithOutpath(outpath),
		github.ServiceContextWithToken(token),
	)

	if err != nil {
		return err
	}

	if svc.Repository.IsRoot() {
		return down.DownloadRootZip(svc)
	}

	if svc.Repository.IsFile() {
		return down.DownloadFile(svc)
	}

	if svc.Repository.IsDir() {
		return down.DownloadDir(svc)
	}

	return nil
}
