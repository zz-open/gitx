package downloader

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"
	"sync"

	"github.com/panjf2000/ants/v2"
	"github.com/zz-open/zbin/common"
	"github.com/zz-open/zbin/ghd/sc"
)

type ResourceChunk struct {
	GitUrl  string `json:"git_url"`
	Path    string `json:"path"`
	OutPath string `json:"outpath"`
}

func NewTreeDownloader(svc *sc.ServiceContext) *TreeDownloader {
	td := &TreeDownloader{
		ServiceContext: svc,
		FileChunks:     make([]*ResourceChunk, 0),
		GitTreeChunks:  make([]*ResourceChunk, 0),
	}

	return td
}

type TreeDownloader struct {
	ServiceContext *sc.ServiceContext
	FileChunks     []*ResourceChunk // 最终要下载的所有文件存到此处
	GitTreeChunks  []*ResourceChunk // 需要递归的顶级目录存到此处
}

func (td *TreeDownloader) Download() error {
	err := td.fetchContents()
	if err != nil {
		return err
	}

	err = td.fetchGitTrees()
	if err != nil {
		return err
	}

	return td.asyncFetchBlob()
}

func (td *TreeDownloader) fetchContents() error {
	svc := td.ServiceContext
	repository := svc.Repository
	httpClient := svc.HttpClient

	resp, err := httpClient.SendDirReposContentRequest(repository.ContentApiUrl())
	if err != nil {
		return err
	}

	if resp == nil || len(resp) == 0 {
		return errors.New("contents api 返回结果为空")
	}

	lastLevelDirname := repository.LastLevelDirname()
	for _, v := range resp {
		if v.IsFile() {
			td.FileChunks = append(td.FileChunks, &ResourceChunk{GitUrl: v.GitUrl, Path: v.Path, OutPath: path.Join(lastLevelDirname, v.Name)})
		} else if v.IsDir() {
			td.GitTreeChunks = append(td.GitTreeChunks, &ResourceChunk{GitUrl: v.GitUrl, Path: v.Path, OutPath: path.Join(lastLevelDirname, v.Name)})
		}
	}

	return nil
}

func (td *TreeDownloader) fetchGitTrees() error {
	if td.GitTreeChunks == nil || len(td.GitTreeChunks) == 0 {
		return nil
	}

	svc := td.ServiceContext
	httpClient := svc.HttpClient

	for _, v := range td.GitTreeChunks {
		node, err := httpClient.SendReposGitTreesRequest(fmt.Sprintf("%s?recursive=1", v.GitUrl))
		if err != nil {
			return err
		}

		if node.Tree == nil || len(node.Tree) == 0 {
			continue
		}

		for _, v1 := range node.Tree {
			if v1.IsBlob() {
				td.FileChunks = append(td.FileChunks, &ResourceChunk{
					GitUrl:  v1.Url,
					Path:    path.Join(v.Path, v1.Path),
					OutPath: path.Join(v.OutPath, v1.Path),
				})
			}
		}
	}

	return nil
}

func (td *TreeDownloader) createFile(chunk *ResourceChunk) error {
	if chunk == nil {
		return nil
	}

	svc := td.ServiceContext
	httpClient := svc.HttpClient

	blob, err := httpClient.SendReposGitBlobsRequest(chunk.GitUrl)
	if err != nil {
		return err
	}

	fileContent, err := common.DecodeGithubContent(blob.Encoding, blob.Content)
	if err != nil {
		return err
	}

	err = common.WriteFile(fileContent, filepath.Join(svc.Outpath, chunk.OutPath))
	if err != nil {
		return err
	}

	return nil
}

func (td *TreeDownloader) asyncFetchBlob() error {
	if len(td.FileChunks) == 0 {
		return errors.New("目录中不包含任何文件")
	}

	var wg sync.WaitGroup

	p, err := ants.NewPoolWithFunc(20, func(v any) {
		c, ok := v.(*ResourceChunk)
		if !ok {
			return
		}

		taskErr := td.createFile(c)
		if taskErr != nil {
			fmt.Printf("err= %s\n", taskErr)
		}

		wg.Done()
	})
	if err != nil {
		return err
	}

	defer p.Release()

	for _, v := range td.FileChunks {
		wg.Add(1)
		_ = p.Invoke(v)
	}

	wg.Wait()
	return nil
}
