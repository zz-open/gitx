package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tpl",
	Short: "基于模板创建项目",
	Long: `基于模板创建项目，支持以下模板：
	1. 内置了作者自己封装的模板
	2. 第三方模板，如vite等`,
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true // 禁用自动补全子命令
	rootCmd.Version = fmt.Sprintf("%s %s/%s", "0.0.1", runtime.GOOS, runtime.GOARCH)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
