package fs

import "os"

// Flag represents the flags with which a file can be opened.
type Flag int

const (
	// OReadOnly is a flag to open a file in READ mode.
	OReadOnly Flag = iota

	// OReadWrite is a flag to open a file in READ and WRITE mode.
	OReadWrite
)

// String returns the flag as a string.
func (flag Flag) String() string {
	names := []string{
		"ReadOnly",
		"ReadWrite",
	}

	return names[flag]
}

// Convert a fs.Flag instance to an os.File flag integer.
func flagToFlag(flag Flag) (f int) {
	switch flag {
	case OReadOnly:
		f = os.O_RDONLY
	case OReadWrite:
		f = os.O_RDWR
	default:
		panic("unknown flag")
	}

	return f
}
