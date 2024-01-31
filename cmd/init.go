package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
	"github.com/zz-open/tpl/internal/git"
)

var (
	remoteUrl string
)

var (
	initCmd = &cobra.Command{
		Use:     "init [remoteUrl]",
		Aliases: []string{"create"},
		Short:   "初始化项目",
		Long:    ``,
		Args: func(_ *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires at least one arg")
			}

			remoteUrl = args[0]
			return nil
		},
		Run: func(_ *cobra.Command, args []string) {
			log.Println(args)
			git.RequestGithubProvideUrl()
		},
	}
)

func init() {
	// initCmd.PersistentFlags().StringVarP(&projectName, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
}
