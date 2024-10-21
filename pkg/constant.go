package pkg

var Version = "v0.0.1"

const DefaultValue = ""

const (
	LineCountUseCommand = "linecount -f {FILE_NAME} | --file {FILE_NAME}"
	CheckSumUseCommand  = "checksum -f {FILE_NAME} | --file {FILE_NAME} {ALGO}"
)

const (
	LineCountShortCommand = "Print line count of file"
	CheckSumShortCommand  = "Print checksum of file"
)
