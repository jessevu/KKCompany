package main

import (
	"cli-tool/cmd/app"

	"github.com/alexflint/go-arg"
)

func main() {
	args := loadArgsValid()
	app.Run(args)
}

func loadArgsValid() app.Args {
	var args app.Args
	arg.MustParse(&args)
	return args
}
