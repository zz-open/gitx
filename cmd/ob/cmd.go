package ob

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/zz-open/zb/common"
)

var (
	url string
	Cmd *cobra.Command
)

func init() {
	Cmd = &cobra.Command{
		Use:   "ob [url]",
		Short: "打开默认浏览器",
		Long:  `使用默认浏览器打开url`,
		Args: func(_ *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires at least one arg")
			}

			url = args[0]
			return nil
		},
		Run: func(_ *cobra.Command, args []string) {
			common.OpenUrl(url)
		},
	}
}
