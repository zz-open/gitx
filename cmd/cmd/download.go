package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

var (
	remoteUrl string
)

var (
	downloadCmd = &cobra.Command{
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
		},
	}
)

func init() {
	// initCmd.PersistentFlags().StringVarP(&projectName, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
}
