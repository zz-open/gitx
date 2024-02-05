package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"runtime"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/zz-open/zb/cmd/dsn"
	"github.com/zz-open/zb/cmd/ghd"
	"github.com/zz-open/zb/common"
)

var (
	//go:embed usage.tpl
	usageTpl string

	rootCmd *cobra.Command
)

const _UI = `
|\_____  \|\   __  \    
 \|___/  /\ \  \|\ /_   
     /  / /\ \   __  \  
    /  /_/__\ \  \|\  \ 
   |\________\ \_______\
    \|_______|\|_______|
`

func init() {
	cobra.OnInitialize(initialize)
	rootCmd = &cobra.Command{
		Use:   "zb",
		Short: "命令行工具",
		Long: `命令行工具，包含以下功能:
1. 下载github repository资源
2. 输出dsn示例`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(common.Cyan(_UI))
		},
	}

	rootCmd.AddCommand(ghd.Cmd)
	rootCmd.AddCommand(dsn.Cmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true // 禁用自动补全子命令
	rootCmd.Version = fmt.Sprintf("%s %s/%s", "0.0.1", runtime.GOOS, runtime.GOARCH)
	rootCmd.SetUsageTemplate(usageTpl)
}

func initialize() {
	cobra.AddTemplateFuncs(template.FuncMap{
		"blue":    common.Blue,
		"green":   common.Green,
		"rpadx":   common.Rpadx,
		"rainbow": common.Rainbow,
	})
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
