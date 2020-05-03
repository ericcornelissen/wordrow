package fs

import "io/ioutil"

import "github.com/ericcornelissen/wordrow/internal/logger"

// The standard write mode to write files.
const mode = 0644

// WriteFile writes a string to a file.
//
// If the file could not be written to, an error is logged.
func WriteFile(filePath string, content string) {
	contentAsBytes := []byte(content)
	err := ioutil.WriteFile(filePath, contentAsBytes, mode)
	if err != nil {
		logger.Errorf("Could not write to file '%s' (%d)", filePath, mode)
	}
}
