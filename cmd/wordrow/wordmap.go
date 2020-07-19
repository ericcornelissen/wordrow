package main

import (
	"io/ioutil"

	"github.com/ericcornelissen/wordrow/internal/errors"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
	"github.com/ericcornelissen/wordrow/internal/strings"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

// Parse a --map-file argument into its component parts.
//
// A --map-file argument can either be just a file path, or a file path with
// ":format" appended to it. For example:
//
//   /path/to/file.csv
//   /path/to/file.txt:csv
//
// In the former the file extension is returned as format, in the latter the
// explicitly stated format is returned as format.
func parseMapFileArgument(argument string) (filePath string, format string) {
	fileExtension := fs.GetExt(argument)

	tmp := strings.Split(fileExtension, ":")
	if len(tmp) > 1 {
		explicitFormat := tmp[len(tmp)-1]
		filePath := strings.TrimSuffix(argument, ":"+explicitFormat)
		return filePath, explicitFormat
	}

	return argument, fileExtension
}

// Add the mapping of a map file to the `wordmap`. The `format` argument
// determines how the contents of the file are parsed.
func processMapFile(
	reader fs.Reader,
	format string,
	wordmap *wordmaps.WordMap,
) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	content := string(data)
	return wordmap.AddFile(&content, format)
}

// Open a map file and add its mapping to the `wordmap`.
func openAndProcessMapFile(
	fileArgument string,
	wordmap *wordmaps.WordMap,
) error {
	filePath, format := parseMapFileArgument(fileArgument)

	logger.Debugf("Opening '%s' as a '%s' formatted map file", filePath, format)
	handle, err := fs.OpenFile(filePath, fs.OReadOnly)
	if err != nil {
		return errors.Newf("Could not open '%s' in %s mode", filePath, fs.OReadOnly)
	}

	defer handle.Close()
	return processMapFile(handle, format, wordmap)
}

// Open the specified map files and add their mapping to the `wordmap`.
func openAndProcessMapFiles(
	filePaths []string,
	wordmap *wordmaps.WordMap,
) error {
	for _, filePath := range filePaths {
		err := openAndProcessMapFile(filePath, wordmap)
		if err != nil {
			return err
		}
	}

	return nil
}

// Add the CLI defined mappings to the `wordmap`.
func processInlineMappings(
	mappings []string,
	wm *wordmaps.WordMap,
) error {
	for _, mapping := range mappings {
		logger.Debugf("Processing the CLI defined mapping '%s'", mapping)
		values := strings.Split(mapping, ",")
		if len(values) != 2 {
			return errors.Newf("Invalid CLI defined mapping '%s'", mapping)
		}

		wm.AddOne(values[0], values[1])
	}

	return nil
}

// Get a WordMap for the specified map files and inline mappings.
func getWordMap(
	mapFiles []string,
	inlineMappings []string,
) (wordmap wordmaps.WordMap, err error) {
	err = openAndProcessMapFiles(mapFiles, &wordmap)
	if err != nil {
		return wordmap, err
	}

	err = processInlineMappings(inlineMappings, &wordmap)
	return wordmap, err
}
