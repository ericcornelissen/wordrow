package fs

import "os"

// The default mode used by OpenFile.
const mode = 0600

// OpenFile opens a File handle to interact with the file specified by
// `filePath`.
func OpenFile(filePath string, flag Flag) (file *File, err error) {
	osFile, err := os.OpenFile(filePath, flagToFlag(flag), mode)
	return &File{
		handle: osFile,
		path:   filePath,
	}, err
}
