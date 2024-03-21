package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/zz-open/zb/cmd/dsn"
	"github.com/zz-open/zb/cmd/ghd"
)

var (
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
	rootCmd = &cobra.Command{
		Use:   "zb",
		Short: "通用命令行工具",
		Long:  `通用命令行工具`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(_UI)
		},
	}

	rootCmd.AddCommand(ghd.Cmd)
	rootCmd.AddCommand(dsn.Cmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true // 禁用自动补全子命令
	rootCmd.Version = fmt.Sprintf("%s %s/%s", "0.0.1", runtime.GOOS, runtime.GOARCH)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
