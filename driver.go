package ftp

import "io"

// DriverFactory is in charge of creating a Driver
// For each client that connects to the server, a new FTPDriver is required.
// Create an implementation of this interface and provide it to FTPServer.
type DriverFactory interface {
	NewDriver() (Driver, error)
}

// Driver contains methods for interacting with the filesystem.
// You will create an implementation of this interface that speaks to your
// chosen persistence layer. A new instance of your driver will be created
// for each client that connects and delegate to it as required.
type Driver interface {
	// Init init
	Init(*Conn)

	// params  - a file path
	// returns - FileInfo for the path
	//         - an error if the file doesn't exist or the user lacks
	//           permissions
	Stat(string) (FileInfo, error)

	// params  - path
	// returns - an error if the user is not allowed to change path
	ChangeDir(string) error

	// params  - path, function on file or subdir found
	// returns - error
	ListDir(string, func(FileInfo) error) error

	// params  - path
	// returns - error if the directory couldn't be deleted
	DeleteDir(string) error

	// params  - path
	// returns - error if the file couldn't be deleted
	DeleteFile(string) error

	// params  - from_path, to_path
	// returns - error if the file couldn't be renamed
	Rename(string, string) error

	// params  - path
	// returns - error if the folder couldn't be created
	MakeDir(string) error

	// params  - path, offset
	// returns - bytes to be returned
	//				 - a ReadCloser with the contents
	// 				 - error if any
	GetFile(string, int64) (int64, io.ReadCloser, error)

	// params  - destination path
	//				 - an io.Reader containing the file data
	//				 - whether to append the data or not
	// returns - number of bytes written
	//				 - error if any
	PutFile(string, io.Reader, bool) (int64, error)
}
