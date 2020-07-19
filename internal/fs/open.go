package fs

import "os"

// OpenFile opens a handle to interact with the specified file.
func OpenFile(filePath string, flag Flag) (file *File, err error) {
	osFile, err := os.OpenFile(filePath, flagToFlag(flag), 0600)
	return &File{
		handle: osFile,
		path:   filePath,
	}, err
}
