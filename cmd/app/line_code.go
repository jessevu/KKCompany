package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	fileName string
)

func NewLineCountCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "linecount [-f {FILE_NAME} | --file {FILE_NAME}]",
		Short: "Print line count of file",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if fileName == "" {
				return errors.New("the --file flag cannot be empty")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Printf("Error opening file: %v\n", err)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			lineCount := 0
			for scanner.Scan() {
				lineCount++
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading file: %v\n", err)
			}

			fmt.Println(lineCount)
		},
	}

	cmd.Flags().StringVarP(&fileName, "file", "f", "", "input the file name")

	return cmd
}
