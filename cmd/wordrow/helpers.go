package main

import "os"

// FileHandler represents a function to handle a file given its file path.
type fileHandler func(filePath string) error

// ForEach executes a fileHandler for each provided filePath. Any error that
// occurs is accumulated and only returned once all files are handled.
func forEach(filePaths []string, handler fileHandler) (errs []error) {
	for _, filePath := range filePaths {
		err := handler(filePath)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

// Check if the program received input from STDIN.
//
// based on: https://stackoverflow.com/a/38612652
func hasStdin() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	return (fi.Mode() & os.ModeNamedPipe) != 0
}
