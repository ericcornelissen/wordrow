package main

import (
	"io/ioutil"

	"github.com/ericcornelissen/wordrow/internal/errors"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
	"github.com/ericcornelissen/wordrow/internal/replace"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

// Reads the contents from the `reader` and updates the content based on the
// `wordmap`.
func doReplace(
	reader fs.Reader,
	wordmap *wordmaps.WordMap,
) (updatedContent string, er error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return updatedContent, err
	}

	mapping := wordmap.Iter()
	content := string(data)
	return replace.All(content, mapping), nil
}

// Writes the `updatedContents` to the `writer`.
func doWriteBack(writer fs.Writer, updatedContent string) error {
	data := []byte(updatedContent)
	_, err := writer.Write(data)
	return err
}

// Process `file` by reading its content, changing that based on the `wordmap`,
// and writing the updated content back to `file`.If a reading or writing error
// occurs this function returns an error.
func processFile(file fs.ReadWriter, wordmap *wordmaps.WordMap) error {
	logger.Debugf("Reading '%s' and replacing words", file)
	updatedContent, err := doReplace(file, wordmap)
	if err != nil {
		return errors.Newf("Could not read from file '%s'", file)
	}

	logger.Debugf("Writing updated contents to '%s'", file)
	err = doWriteBack(file, updatedContent)
	if err != nil {
		return errors.Newf("Could not write to file '%s'", file)
	}

	return nil
}

// Open the file specified by `filePath` and process it using the `wordmap`. If
// opening the file fails or a reading or writing error occurs this function
// returns an error.
func openAndProcessFile(filePath string, wordmap *wordmaps.WordMap) error {
	logger.Debugf("Opening '%s'", filePath)
	handle, err := fs.OpenFile(filePath, fs.OReadWrite)
	if err != nil {
		return errors.Newf("Could not open '%s' (%s mode)", filePath, fs.OReadWrite)
	}

	defer handle.Close()
	return processFile(handle, wordmap)
}

// Update the contents of all files specified by `filePaths` based on the
// `wordmap`. Any error that occurs is returned after all files have been
// processed.
func processInputFiles(
	filePaths []string,
	wordmap *wordmaps.WordMap,
) (errs []error) {
	for _, filePath := range filePaths {
		logger.Debugf("Processing '%s'", filePath)
		err := openAndProcessFile(filePath, wordmap)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
