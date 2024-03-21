package ghd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zz-open/zb/modules/ghd"
)

var (
	personalToken string
	url           string
	outpath       string
	Cmd           *cobra.Command
)

func init() {
	Cmd = &cobra.Command{
		Use:   "ghd [resourceUrl]",
		Short: "下载 github 项目指定文件",
		Long:  ``,
		Args: func(_ *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires at least one arg")
			}

			url = args[0]
			return nil
		},
		Run: func(_ *cobra.Command, args []string) {
			err := ghd.Download(url, outpath, personalToken)
			if err != nil {
				fmt.Printf("%s\n", err)
			}
		},
	}
	Cmd.PersistentFlags().StringVarP(&personalToken, "token", "", "", "Personal access tokens (classic)")
	Cmd.PersistentFlags().StringVarP(&outpath, "outpath", "o", ".", "本地保存路径")
}
