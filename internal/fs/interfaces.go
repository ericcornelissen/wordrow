package fs

// Writer is TODO
type Writer interface {
	Write(b []byte) (n int, err error)
}

// Reader is TODO
type Reader interface {
	Read(p []byte) (n int, err error)
}

// ReadWriter is TODO
type ReadWriter interface {
	Reader
	Writer
}

// Handle is TODO
type Handle struct {
	path string
}
