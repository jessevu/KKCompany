package app

import (
	"cli-tool/pkg"
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "show current verison application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(pkg.Version)
		},
	}
}
