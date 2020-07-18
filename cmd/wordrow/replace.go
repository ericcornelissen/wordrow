package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/ericcornelissen/wordrow/internal/errors"
	"github.com/ericcornelissen/wordrow/internal/logger"
	"github.com/ericcornelissen/wordrow/internal/replacer"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

func doReplace(
	handle io.Reader,
	wordmap *wordmaps.WordMap,
) (fixed string, err error) {
	data, err := ioutil.ReadAll(handle)
	if err != nil {
		return fixed, err
	}

	content := string(data)
	return replacer.ReplaceAll(content, *wordmap), nil
}

func doWriteBack(
	handle io.Writer,
	fixed string,
) error {
	data := []byte(fixed)
	_, err := handle.Write(data)
	return err
}

func processFile(
	filePath string,
	handle io.ReadWriter,
	wordmap *wordmaps.WordMap,
) error {
	logger.Debugf("Reading '%s' and replacing words", filePath)
	fixed, err := doReplace(handle, wordmap)
	if err != nil {
		return errors.Newf("Could not read from '%s'", filePath)
	}

	logger.Debugf("Writing updated contents to '%s'", filePath)
	err = doWriteBack(handle, fixed)
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
	handle, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return errors.Newf("Could not open '%s'", filePath)
	}
	defer handle.Close()

	return processFile(filePath, handle, wordmap)
}

func processFiles(
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
