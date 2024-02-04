package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/zz-open/zbin/cmd/ghd"
)

var rootCmd = &cobra.Command{
	Use:   "zbin",
	Short: "一款综合命令行工具",
	Long: `包含以下功能:
1. github 仓库下载工具`,
}

func init() {
	rootCmd.AddCommand(ghd.Cmd)

	rootCmd.CompletionOptions.DisableDefaultCmd = true // 禁用自动补全子命令
	rootCmd.Version = fmt.Sprintf("%s %s/%s", "0.0.1", runtime.GOOS, runtime.GOARCH)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
