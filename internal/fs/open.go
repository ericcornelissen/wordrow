package fs

import (
	"os"
	"path/filepath"

	"github.com/ericcornelissen/wordrow/internal/errors"
)

// The default mode used by OpenFile.
const mode = 0600

// OpenFile opens a File handle to interact with the file specified by
// `filePath`.
func OpenFile(filePath string, flag Flag) (file *File, err error) {
	osFile, err := os.OpenFile(filepath.Clean(filePath), flagToFlag(flag), mode)
	if err != nil {
		return nil, errors.Newf("Could not open '%s' (%s mode)", filePath, flag)
	}

	return &File{
		handle: osFile,
		path:   filePath,
	}, err
}
