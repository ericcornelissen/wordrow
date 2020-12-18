package main

import (
	"bufio"
	"io/ioutil"

	"github.com/ericcornelissen/wordrow/internal/errors"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
	"github.com/ericcornelissen/wordrow/internal/replace"
)

// Reads the contents from the `reader` and updates the content based on the
// `mapping`.
func doReplace(
	reader fs.Reader,
	mapping map[string]string,
) (updatedContent []byte, er error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return updatedContent, err
	}

	return replace.All(data, mapping), nil
}

// Writes the `updatedContents` to the `writer`.
func doWriteBack(writer fs.Writer, data []byte) error {
	_, err := writer.Write(data)
	return err
}

// Process the `input` provided by the ReadWriter, changing that based on the
// `mapping`, and write the updated content back to the ReadWriter.
func processStdin(rw *bufio.ReadWriter, mapping map[string]string) error {
	input := bufio.NewScanner(rw.Reader)
	output := rw.Writer

	for input.Scan() {
		line := input.Bytes()
		fixedLine := replace.All(line, mapping)
		output.Write(fixedLine)
		output.WriteRune('\n')
	}

	return output.Flush()
}

// Process `file` by reading its content, changing that based on the `mapping`,
// and writing the updated content back to `file`. If a reading or writing error
// occurs this function returns an error.
func processFile(file fs.ReadWriter, mapping map[string]string) error {
	logger.Debugf("Reading '%s' and replacing words", file)
	updatedContent, err := doReplace(file, mapping)
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

// Opens the file provided by the handler and process it using the `mapping`. If
// opening the file fails or a reading or writing error occurs the error is
// outputted to the channel `ch`.
func openAndProcessFileWith(
	ch chan error,
	mapping map[string]string,
) func(value string) {
	return func(filePath string) {
		logger.Debugf("Opening '%s'", filePath)
		handle, err := fs.OpenFile(filePath, fs.OReadWrite)
		if err != nil {
			ch <- err
		}

		defer handle.Close()

		logger.Debugf("Processing '%s'", filePath)
		ch <- processFile(handle, mapping)
	}
}

// Update the contents of all files specified by `filePaths` based on the
// `mapping`. Any error that occurs is returned after all files have been
// processed.
func processInputFiles(
	filePaths []string,
	mapping map[string]string,
) (errs []error) {
	ch := make(chan error, len(filePaths))
	defer close(ch)

	openAndProcessFile := openAndProcessFileWith(ch, mapping)
	for _, filePath := range filePaths {
		go openAndProcessFile(filePath)
	}

	return drain(ch, len(filePaths))
}
