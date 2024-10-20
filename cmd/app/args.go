package app

type Args struct {
	FileName string `arg:"-f,--file,required" help:"file name to read"`
	Version  bool   `arg:"-s,--version" help:"show version"`
}
