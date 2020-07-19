package main

import (
	"io/ioutil"

	"github.com/ericcornelissen/wordrow/internal/errors"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
	"github.com/ericcornelissen/wordrow/internal/replacer"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

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

func doWriteBack(
	writer fs.Writer,
	fixed string,
) error {
	data := []byte(fixed)
	_, err := writer.Write(data)
	return err
}

func processFile(
	filePath string,
	handle fs.ReadWriter,
	wordmap *wordmaps.WordMap,
) error {
	logger.Debugf("Reading '%s' and replacing words", filePath)
	fixedText, err := doReplace(handle, wordmap)
	if err != nil {
		return errors.Newf("Could not read from '%s'", filePath)
	}

	logger.Debugf("Writing updated contents to '%s'", filePath)
	err = doWriteBack(handle, fixedText)
	if err != nil {
		return errors.Newf("Could not write to '%s'", filePath)
	}

	return nil
}

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

	return processFile(filePath, handle, wordmap)
}

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
