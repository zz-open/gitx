package directory

import (
	"errors"
	"log"
	"strings"

	"github.com/zz-open/gitdownload/internal/github"
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

	sha1 := lastLevelDirSHA1(repository.Path)
	log.Println("sha1:", sha1)
	// 通过content api 获取 最后一级目录的SHA1

	// 将路径拆分成一级一级的目录
	// 遍历寻找目录
	// 拿到最后一级目录的SHA1
	// 递归遍历最后一级，循环获取内容

	return nil
}

// checkParameter 检查入参
func checkParameter(url string, path string) error {
	if url == "" {
		return errors.New("url 不合法")
	}

	if path == "" {
		return errors.New("path 不合法")
	}

	if strings.HasSuffix(url, "/") {
		return errors.New("url 不是一个文件路径")
	}

	return nil
}

func lastLevelDirSHA1(path string) string {

	path = strings.Trim()
	return ""
}
