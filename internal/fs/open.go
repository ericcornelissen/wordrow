package fs

import "os"

const (
	// OReadWrite is a flag to open a file in READ and WRITE mode.
	OReadWrite = os.O_RDWR
)

// OpenFile opens a handle to interact with the specified file.
func OpenFile(
	filePath string,
	flag int,
) (handle *Handle, err error) {
	file, err := os.OpenFile(filePath, flag, 0644)
	return &Handle{handle: file}, err
}
