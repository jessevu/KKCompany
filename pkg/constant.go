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

const (
	MD5    = "md5"
	SHA1   = "sha1"
	SHA256 = "sha256"
)
