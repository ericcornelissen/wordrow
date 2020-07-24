package fs

import (
	"os"
	"path/filepath"
)

// The default mode used by OpenFile.
const mode = 0600

// OpenFile opens a File handle to interact with the file specified by
// `filePath`.
func OpenFile(filePath string, flag Flag) (file *File, err error) {
	cleanFilePath := filepath.Clean(filePath)
	osFile, err := os.OpenFile(cleanFilePath, flagToFlag(flag), mode)
	return &File{
		handle: osFile,
		path:   filePath,
	}, err
}
