package fs

import "os"

// Reader is the interface that wraps the basic Read method.
type Reader interface {
	// Read data into `data`.
	Read(data []byte) (n int, err error)
}

// Writer is the interface that wraps the basic Write method.
type Writer interface {
	// Write `data` into the file.
	Write(data []byte) (n int, err error)
}

// ReadWriter is the interface that groups the basic Read and Write methods.
type ReadWriter interface {
	Reader
	Writer
}

// Handle is a file struct that implements ReadWriter.
type Handle struct {
	handle *os.File
	path   string
}

// Close closes this handle.
func (h Handle) Close() error {
	return h.handle.Close()
}

// Read reads the contents of the file into `data`.
func (h Handle) Read(data []byte) (n int, err error) {
	return h.handle.Read(data)
}

// String returns the path of the file in the handle.
func (h Handle) String() string {
	return h.path
}

// Write empties the file and writes `data` into it.
func (h Handle) Write(data []byte) (n int, err error) {
	if err := h.handle.Truncate(0); err != nil {
		return 0, err
	}

	if _, err := h.handle.Seek(0, 0); err != nil {
		return 0, err
	}

	return h.handle.Write(data)
}
