package pkg

import "errors"

var (
	ErrFileNameEmpty   = errors.New("the --file flag cannot be empty")
	ErrUnsupportedAlgo = errors.New("unsupported algorithm")
)
