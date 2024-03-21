package dsn

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zz-open/zb/modules/dsn"
)

var (
	Cmd *cobra.Command
)

func init() {
	Cmd = &cobra.Command{
		Use:   "dsn",
		Short: "输出数据库dsn示例",
		Long:  ``,
		Run: func(_ *cobra.Command, args []string) {
			dsnInstance := dsn.NewDsn(
				dsn.DsnWithHost("127.0.0.1"),
				dsn.DsnWithPort(3306),
				dsn.DsnWithUsername("root"),
				dsn.DsnWithPassword("123456"),
				dsn.DsnWithDatabase("test"),
				dsn.DsnWithCharset("utf8mb4"),
			)

			fmt.Printf("示例DSN: %s\n", dsnInstance.ToString())
		},
	}
}
