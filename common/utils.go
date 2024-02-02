package common

import (
	"bufio"
	"os"
	"strings"
)

// FilterTailSlash 过滤字符串末尾的 '/'
func FilterTailSlash(str string) string {
	return strings.TrimSuffix(str, "/")
}

// writeFile 写入content 到 outpath
func WriteFile(outpath string, content []byte) error {
	file, err := os.Create(outpath)
	if err != nil {
		return err
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.Write(content)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
