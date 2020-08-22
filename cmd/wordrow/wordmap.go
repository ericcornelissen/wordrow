package main

import (
	"io/ioutil"

	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

// Parse a --map-file argument into its component parts.
//
// A --map-file argument can either be just a file path, or a file path with an
// explicit ":format" appended to it. For example:
//
//   /path/to/file.csv
//   /path/to/file.txt:csv
//
// In the former the file extension is returned as format, in the latter the
// explicitly stated format is returned as format.
//
// If the file does not have an extension it is expected to have an explicit
// ":format" appended to it. For example:
//
//   /path/to/file:csv
//   /path/to/file
//
// In the former the file explicitly stated format, in the latter no format is
// returned.
func parseMapFileArgument(argument string) (filePath, format string) {
	fileExtension := fs.GetExt(argument)

	explicitFormatSplit := stringsx.Split(argument, ":")
	if len(explicitFormatSplit) > 1 {
		explicitFormat := explicitFormatSplit[len(explicitFormatSplit)-1]
		filePath := stringsx.TrimSuffix(argument, ":"+explicitFormat)
		return filePath, explicitFormat
	}

	return argument, fileExtension
}

// Add the mapping of `reader` to the `wordmap`. The `format` argument
// determines how the contents of the file are parsed. This function returns an
// error if either the reading or parsing fails.
func processMapFile(
	reader fs.Reader,
	format string,
	wordmap *wordmaps.StringMap,
) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	content := string(data)
	return wordmap.AddFile(&content, format)
}

// Opens the file provided by the handler and add its mapping to the `wordmap`.
// If the file cannot be opened or processing failed the function returns an
// error.
func openAndProcessMapFileWith(wordmap *wordmaps.StringMap) fileHandler {
	return func(fileArgument string) error {
		filePath, format := parseMapFileArgument(fileArgument)

		logger.Debugf("Opening '%s' as a '%s' formatted map file", filePath, format)
		handle, err := fs.OpenFile(filePath, fs.OReadOnly)
		if err != nil {
			return err
		}

		defer handle.Close()

		logger.Debugf("Processing '%s' as a map file", filePath)
		return processMapFile(handle, format, wordmap)
	}
}

// Add a CLI defined mapping to the `wordmap`. If the mapping is invalid this
// function returns an error (and leave `wordmap` unchanged).
func processInlineMapping(mapping string, wordmap *wordmaps.StringMap) error {
	return wordmap.AddFile(&mapping, "csv")
}

// Add all CLI defined mappings to the `wordmap`. Any error that occurs is
// returned after all mappings have been processed.
func processInlineMappings(
	mappings []string,
	wordmap *wordmaps.StringMap,
) (errs []error) {
	for _, mapping := range mappings {
		logger.Debugf("Processing '%s' as a CLI specified mapping", mapping)
		err := processInlineMapping(mapping, wordmap)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

// Get a WordMap for the specified `mapFiles` and `inlineMappings`. Any error
// that occurs is returned after both have been processed. In case of any error
// the `wordmap` that is returned represents only the arguments that could be
// successfully processed.
func getWordMap(
	mapFiles []string,
	inlineMappings []string,
) (wordmap wordmaps.StringMap, errs []error) {
	errs = forEach(mapFiles, openAndProcessMapFileWith(&wordmap))
	errs = append(errs, processInlineMappings(inlineMappings, &wordmap)...)
	return wordmap, errs
}
