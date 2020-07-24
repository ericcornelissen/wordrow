// Package fs is a simple utilities package that provides functions to interact
// with the file system.
package fs

// Reader is the interface that wraps the basic Read method.
type Reader interface {
	// Read file contents into `data`. The first return value represent the number
	// of bytes read. The second return value can be used to set an error if
	// reading failed.
	Read(data []byte) (n int, err error)
}

// Writer is the interface that wraps the basic Write method.
type Writer interface {
	// Write `data` into the file. The first return value represent the number of
	// bytes written. The second return value can be used to set an error if
	// writing failed.
	Write(data []byte) (n int, err error)
}

// ReadWriter is the interface that groups the basic Read and Write methods.
type ReadWriter interface {
	Reader
	Writer
}
