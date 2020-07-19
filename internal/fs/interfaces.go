package fs

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
