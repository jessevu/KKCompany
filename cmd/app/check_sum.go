package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCheckSumCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "checksum",
		Short: "Print checksum of file",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}
}
