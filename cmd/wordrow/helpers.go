package main

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
