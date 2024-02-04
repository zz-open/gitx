package ghd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zz-open/zbin/ghdownloader"
)

var (
	personalToken string
	url           string
	outpath       string
)

var (
	Cmd = &cobra.Command{
		Use:   "ghd [resourceUrl]",
		Short: "github repository 下载工具",
		Long:  ``,
		Args: func(_ *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires at least one arg")
			}

			url = args[0]
			return nil
		},
		Run: func(_ *cobra.Command, args []string) {
			err := ghdownloader.Download(url, outpath, personalToken)
			if err != nil {
				fmt.Printf("err: %s\n", err)
			}
		},
	}
)

func init() {
	Cmd.PersistentFlags().StringVarP(&personalToken, "token", "", "", "github personal token")
	Cmd.PersistentFlags().StringVarP(&outpath, "outpath", "o", ".", "本地保存路径")
}
