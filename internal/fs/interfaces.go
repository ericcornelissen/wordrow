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

// Handle is the struct that implements ReadWriter.
type Handle struct {
	handle *os.File
}

// Close will close this handle.
func (h Handle) Close() error {
	return h.handle.Close()
}

// Read reads into data.
func (h Handle) Read(data []byte) (n int, err error) {
	return h.handle.Read(data)
}

// Write data into the file.
func (h Handle) Write(data []byte) (n int, err error) {
	return h.handle.Write(data)
}
