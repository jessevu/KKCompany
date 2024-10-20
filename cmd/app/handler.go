package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{
		Use:   "futil",
		Short: "futil is a cli interface",
		Long:  `File Utility`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Usage()
		},
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(
		NewCheckSumCommand(),
		NewLineCountCommand(),
		NewVersionCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
