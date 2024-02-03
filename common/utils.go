package common

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// ParseAttachmentFilename 提取content-disposition中的 filename
func ParseAttachmentFilename(contentDisposition string) (string, error) {
	_, params, err := mime.ParseMediaType(contentDisposition)
	return params["filename"], err
}

// FilterTailSlash 过滤字符串末尾的 '/'
func FilterTailSlash(str string) string {
	return strings.TrimSuffix(str, "/")
}

// writeFile 文件内容写入outpath，且路径自动创建
func WriteFile(content []byte, outpath string) error {
	err := EnsureDir(outpath)
	if err != nil {
		return err
	}

	return os.WriteFile(outpath, content, 0666)
}

// EnsureDir 确保路径中的目录已创建
func EnsureDir(outpath string) error {
	var err error
	if outpath == "" {
		return errors.New("path 不能为空")
	}

	if !path.IsAbs(outpath) {
		outpath, err = filepath.Abs(outpath)
		if err != nil {
			return err
		}
	}

	err = os.MkdirAll(path.Dir(outpath), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// DecodeGithubContent 解析github content字段
func DecodeGithubContent(encoding string, content string) ([]byte, error) {
	if encoding == "base64" {
		return base64.StdEncoding.DecodeString(content)
	}

	return nil, errors.New("不存在的编码")
}

func RandomDelay(maxdelay int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	delay := rand.Float64() * float64(maxdelay)
	fmt.Printf("Delay time: %.2fs\n", delay)
	time.Sleep(time.Duration(delay * float64(time.Second)))
}
