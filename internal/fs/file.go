package fs

import "os"

// File is a struct, representing a file, that implements ReadWriter.
type File struct {
	// The file handle from the operating system (OS).
	handle *os.File

	// The absolute path of the File.
	path string
}

// Close closes the operating system (OS) handle for this File. After calling
// Close the File cannot be used for reading or writing anymore.
func (f File) Close() error {
	return f.handle.Close()
}

// Read reads the contents of the File into `data`. It returns the amount of
// bytes read in the first return value. It may return an error in the second
// return value if reading failed.
func (f File) Read(data []byte) (n int, err error) {
	return f.handle.Read(data)
}

// String returns the absolute path of the File.
func (f File) String() string {
	return f.path
}

// Write empties the File and writes `data` into it. It returns the amount of
// bytes written in the first return value. It may return an error in the second
// return value if emptying or writing failed.
func (f File) Write(data []byte) (n int, err error) {
	if err := f.handle.Truncate(0); err != nil {
		return 0, err
	}

	if _, err := f.handle.Seek(0, 0); err != nil {
		return 0, err
	}

	return f.handle.Write(data)
}
