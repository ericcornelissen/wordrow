package fs

import "os"

// Flag represents the flags with which a file can be opened.
type Flag uint16

const (
	// ORead is a flag to open a file in READ mode.
	ORead = iota

	// OReadWrite is a flag to open a file in READ and WRITE mode.
	OReadWrite
)

// FlagToFlag converts fs.Flag to os.File flag integers.
func flagToFlag(flag Flag) (f int) {
	switch flag {
	case ORead:
		f = os.O_RDONLY
	case OReadWrite:
		f = os.O_RDWR
	}

	return f
}

// OpenFile opens a handle to interact with the specified file.
func OpenFile(filePath string, flag Flag) (file *Handle, err error) {
	osFile, err := os.OpenFile(filePath, flagToFlag(flag), 0600)
	return &Handle{handle: osFile}, err
}
