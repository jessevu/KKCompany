package app

import (
	"bufio"
	"cli-tool/pkg"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewLineCountCommand() *cobra.Command {
	var filename string
	cmd := &cobra.Command{
		Use:   pkg.LineCountUseCommand,
		Short: pkg.LineCountShortCommand,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if filename == pkg.DefaultValue {
				return pkg.ErrFileNameEmpty
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if count, err := LineCount(filename); err != nil {
				fmt.Println("error:", err)
			} else {
				fmt.Println(*count)
			}
		},
	}

	cmd.Flags().StringVarP(&filename, "file", "f", "", "input the file name")
	cmd.MarkFlagRequired("file")
	return cmd
}

func LineCount(filePath string) (*int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("no such file '%s'", filePath)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %s", err)
	}

	return &lineCount, nil
}
