package ftp

import "os"

// FileInfo contains useful information about a file or directory
type FileInfo interface {
	os.FileInfo

	Owner() string
	Group() string
}
