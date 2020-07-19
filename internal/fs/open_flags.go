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

func (flag Flag) String() string {
	names := []string{
		"ReadOnly",
		"ReadWrite",
	}

	return names[flag]
}

// FlagToFlag converts fs.Flag to os.File flag integers.
func flagToFlag(flag Flag) (f int) {
	switch flag {
	case OReadOnly:
		f = os.O_RDONLY
	case OReadWrite:
		f = os.O_RDWR
	default:
		panic(1)
	}

	return f
}
