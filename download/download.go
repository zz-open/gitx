package downlaod

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/zz-open/gitx/common"
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

	// 下载整个项目
	if svc.Repository.IsRoot() {
		return nil
	}

	err = download(svc)
	if err != nil {
		return err
	}

	fmt.Println("download success !")
	return nil
}

func download(svc *github.ServiceContext) error {
	treeItem, err := searchTargetTreeItem(svc)
	if err != nil {
		return err
	}

	chunks := make([]*github.FileChunk, 0)
	if treeItem.IsFile() {
		chunks = append(chunks, github.NewFileChunk(treeItem.Url, treeItem.PathChain, treeItem.RawDownloadUrl))
	}

	return AsyncFetch(svc, chunks)
}

// 寻找路径中最后一个资源对应的treeItem 结构数据
func searchTargetTreeItem(svc *github.ServiceContext) (*github.ReposGitTreesItem, error) {
	repository := svc.Repository
	paths := strings.Split(repository.Path, "/")

	pathsLength := len(paths)
	var treeItem *github.ReposGitTreesItem

	var f func(url string, pathsIndex int, pathChain string) error
	f = func(url string, pathsIndex int, pathChain string) error {
		if pathsIndex >= pathsLength {
			return nil
		}

		targetPath := paths[pathsIndex]
		resp, err := github.RequestReposGitTrees(url, repository.Token)
		if err != nil {
			return err
		}

		if resp.Truncated {
			return errors.New("超出API限制")
		}

		if resp.Tree == nil {
			return errors.New("暂无结果")
		}

		for _, v := range resp.Tree {
			if v.Path == targetPath {
				// 遍历到了最后一个
				if pathsIndex+1 >= pathsLength {
					_v := v
					treeItem = &_v
					// 设置带级别的路径变量
					treeItem.PathChain = pathChain
					// 设置raw.githubusercontent.com 下载路径
					treeItem.RawDownloadUrl = github.RawUserContentUrl(repository.Username, repository.Repo, repository.Branch, pathChain)

					return nil
				} else {
					return f(v.Url, pathsIndex+1, fmt.Sprintf("%s/%s", pathChain, v.Path))
				}
			}
		}

		return nil
	}

	// 不使用递归遍历，避免超出API限制
	err := f(repository.GitTreesApiUrl("", false), 0, paths[0])
	if err != nil {
		return nil, err
	}

	return treeItem, nil
}

func AsyncFetch(svc *github.ServiceContext, chunks []*github.FileChunk) error {
	if chunks == nil {
		return nil
	}

	token := svc.Repository.Token

	var wg sync.WaitGroup

	for _, v := range chunks {
		wg.Add(1)
		go func(chunk *github.FileChunk) {
			defer wg.Done()
			resp, err := github.RequestReposGitBlobs(chunk.GitBlobUrl, token)
			if err != nil {
				fmt.Printf("v: %+v 请求失败\n", chunk)
				return
			}

			var fileContent []byte
			if resp.Encoding == "base64" {
				// base64 直接解码
				fileContent, err = base64.StdEncoding.DecodeString(resp.Content)
				if err != nil {
					fmt.Printf("base64解码错误: %s\n", err)
					return
				}

			} else {
				// 非base64 通过 raw.githubusercontent.com 下载
				fileContent, err = github.RequestRawGithubUserContent(chunk.DownloadUrl)
				if err != nil {
					fmt.Printf("请求: %s 失败, %s\n", chunk.DownloadUrl, err)
					return
				}
			}

			localPath := ""
			err = common.WriteFile(localPath, fileContent)
			if err != nil {
				fmt.Printf("文件写入错误: %s\n", err)
				return
			}
		}(v)
	}

	wg.Wait()
	return nil
}
