package fs

import "os"

const (
	// OReadWrite is a flag to open a file in READ and WRITE mode.
	OReadWrite = os.O_RDWR
)

// OpenFile opens a file to
func OpenFile(
	filePath string,
	flag int,
	mode os.FileMode,
) (file *os.File, err error) {
	return os.OpenFile(filePath, flag, mode)
}
