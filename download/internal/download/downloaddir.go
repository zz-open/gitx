package download

import (
	"errors"
	"fmt"
	"math/rand"
	"path/filepath"
	"sync"
	"time"

	"github.com/zz-open/gitx/common"
	"github.com/zz-open/gitx/download/internal/github"
)

type DirDownloader struct {
	ServiceContext *github.ServiceContext
	Chunks         []*FileChunk
	GitTreeUrls    []string
	Err            error
}

func (d *DirDownloader) a() {}

func DownloadDir(svc *github.ServiceContext) error {
	downloader := &DirDownloader{
		ServiceContext: svc,
		Chunks:         make([]*FileChunk, 0),
		GitTreeUrls:    []string{},
		Err:            nil,
	}

	downloader.fetchContents()
	if downloader.Err != nil {
		return downloader.Err
	}

	downloader.fetchGitTrees()
	if downloader.Err != nil {
		return downloader.Err
	}

	downloader.fetchBlobs()
	if downloader.Err != nil {
		return downloader.Err
	}

	fmt.Printf("目录下载成功\n")
	return nil
}

func (d *DirDownloader) fetchContents() {
	repository := d.ServiceContext.Repository
	resp, err := github.RequestReposContentWithDirPath(repository.ContentApiUrl(), d.ServiceContext.Token)
	if err != nil {
		d.Err = err
		return
	}

	if resp == nil || len(resp) == 0 {
		d.Err = errors.New("contents api 返回结果为空")
		return
	}

	for _, v := range resp {
		if v.IsFile() {
			d.Chunks = append(d.Chunks, &FileChunk{GitUrl: v.GitUrl, Path: v.Path})
		} else if v.IsDir() {
			d.GitTreeUrls = append(d.GitTreeUrls, v.GitUrl)
		}
	}
}

func (d *DirDownloader) fetchGitTrees() {
	if len(d.GitTreeUrls) == 0 {
		return
	}

	for _, v := range d.GitTreeUrls {
		node, err := github.RequestReposGitTrees(fmt.Sprintf("%s?recursive=1", v), d.ServiceContext.Token)
		if err != nil {
			d.Err = err
			return
		}

		if node.Tree == nil || len(node.Tree) == 0 {
			continue
		}

		for _, v := range node.Tree {
			if v.IsBlob() {
				d.Chunks = append(d.Chunks, &FileChunk{GitUrl: v.Url, Path: v.Path})
			}
		}
	}
}

func (d *DirDownloader) fetchBlobs() {
	if len(d.Chunks) == 0 {
		return
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	outpath := d.ServiceContext.Outpath
	ch := make(chan *github.ReposGitBlobs, 10)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range d.Chunks {
			time.Sleep(200 * time.Millisecond)

			node, err := github.RequestReposGitBlobs(v.GitUrl, d.ServiceContext.Token)
			if err != nil {
				d.Err = err
				return
			}

			node.CustomPath = v.Path
			ch <- node
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range ch {
			fileContent, err := common.DecodeGithubContent(v.Encoding, v.Content)
			if err != nil {
				fmt.Printf("Encoding错误: %s\n", err)
				continue
			}

			err = common.WriteFile(fileContent, filepath.Join(outpath, v.CustomPath))
			if err != nil {
				fmt.Printf("文件写入错误: %s\n", err)
				continue
			}
		}
	}()

	wg.Wait()
}
