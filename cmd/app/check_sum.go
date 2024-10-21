package app

import (
	"cli-tool/pkg"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func NewCheckSumCommand() *cobra.Command {
	var filename, algo string
	cmd := &cobra.Command{
		Use:   pkg.CheckSumShortCommand,
		Short: pkg.CheckSumShortCommand,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if filename == pkg.DefaultValue {
				return pkg.ErrFileNameEmpty
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if hash, err := CheckSum(filename, algo); err != nil {
				fmt.Println("error:", err)
			} else {
				fmt.Println(*hash)
			}
		},
	}

	cmd.Flags().StringVarP(&filename, "file", "f", "", "the input file")
	cmd.Flags().StringVar(&algo, "algo", "md5", "checksum algorithm (md5, sha1, sha256)")
	cmd.MarkFlagRequired("file")

	return cmd
}

func CheckSum(filename, option string) (*string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("no such file '%s'", filename)
	}
	defer file.Close()

	var hash string
	switch option {
	case "md5":
		h := md5.New()
		io.Copy(h, file)
		hash = hex.EncodeToString(h.Sum(nil))
	case "sha1":
		h := sha1.New()
		io.Copy(h, file)
		hash = hex.EncodeToString(h.Sum(nil))
	case "sha256":
		h := sha256.New()
		io.Copy(h, file)
		hash = hex.EncodeToString(h.Sum(nil))
	default:
		return nil, pkg.ErrUnsupportedAlgo
	}
	return &hash, nil
}
