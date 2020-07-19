package main

import (
	"io/ioutil"

	"github.com/ericcornelissen/wordrow/internal/errors"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
	"github.com/ericcornelissen/wordrow/internal/replacer"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

// Do change the contents of `reader` based on the `wordmap`.
func doReplace(
	reader fs.Reader,
	wordmap *wordmaps.WordMap,
) (fixedText string, err error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return fixedText, err
	}

	content := string(data)
	return replacer.ReplaceAll(content, *wordmap), nil
}

// Do write back the `updatedContents` to `writer`.
func doWriteBack(
	writer fs.Writer,
	updatedContent string,
) error {
	data := []byte(updatedContent)
	_, err := writer.Write(data)
	return err
}

// Process `file` by reading its content, changed that based on the `wordmap`,
// and writing the updated content back.
func processFile(
	file fs.ReadWriter,
	wordmap *wordmaps.WordMap,
) error {
	logger.Debugf("Reading '%s' and replacing words", file)
	fixedText, err := doReplace(file, wordmap)
	if err != nil {
		return errors.New("Could not read from file")
	}

	logger.Debugf("Writing updated contents to '%s'", file)
	err = doWriteBack(file, fixedText)
	if err != nil {
		return errors.New("Could not write to file")
	}

	return nil
}

// Open the specified input file and process it.
func openAndProcessFile(
	filePath string,
	wordmap *wordmaps.WordMap,
) error {
	logger.Debugf("Opening '%s'", filePath)
	handle, err := fs.OpenFile(filePath, fs.OReadWrite)
	if err != nil {
		return errors.Newf("Could not open '%s' in %s mode", filePath, fs.OReadWrite)
	}

	defer handle.Close()
	return processFile(handle, wordmap)
}

// Update the contents of all specified files based on the `wordmap`.
func processInputFiles(
	filePaths []string,
	wordmap *wordmaps.WordMap,
) error {
	for _, filePath := range filePaths {
		logger.Debugf("Processing '%s'", filePath)
		err := openAndProcessFile(filePath, wordmap)
		if err != nil {
			return err
		}
	}

	return nil
}
