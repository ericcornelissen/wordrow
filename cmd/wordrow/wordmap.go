package main

import (
	"io/ioutil"

	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/errors"
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
func parseMapFileArgument(argument string) (filePath string, format string) {
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
	wordmap *wordmaps.WordMap,
) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	content := string(data)
	return wordmap.AddFile(&content, format)
}

// Open the map file specified by `fileArgument` and add its mapping to the
// `wordmap`. If the file cannot be opened or processing failed the function
// returns an error.
func openAndProcessMapFile(
	fileArgument string,
	wordmap *wordmaps.WordMap,
) error {
	filePath, format := parseMapFileArgument(fileArgument)

	logger.Debugf("Opening '%s' as a '%s' formatted map file", filePath, format)
	handle, err := fs.OpenFile(filePath, fs.OReadOnly)
	if err != nil {
		return errors.Newf("Could not open '%s' (%s mode)", filePath, fs.OReadOnly)
	}

	defer handle.Close()
	return processMapFile(handle, format, wordmap)
}

// Open the map files specified by `filePaths` and add their mapping to the
// `wordmap`. Any error that occurs is returned after all files have been
// processed.
func openAndProcessMapFiles(
	filePaths []string,
	wordmap *wordmaps.WordMap,
) (errs []error) {
	for _, filePath := range filePaths {
		logger.Debugf("Processing '%s' as a map file", filePath)
		err := openAndProcessMapFile(filePath, wordmap)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

// Add a CLI defined mapping to the `wordmap`. If the mapping is invalid this
// function returns an error (and leave `wordmap` unchanged).
func processInlineMapping(mapping string, wordmap *wordmaps.WordMap) error {
	values := stringsx.Split(mapping, ",")
	if len(values) != 2 {
		return errors.Newf("Invalid CLI defined mapping '%s'", mapping)
	}

	from, to := values[0], values[1]
	if stringsx.IsEmpty(from) || stringsx.IsEmpty(to) {
		return errors.Newf("Missing value in CLI defined mapping '%s'", mapping)
	}

	wordmap.AddOne(from, to)
	return nil
}

// Add all CLI defined mappings to the `wordmap`. Any error that occurs is
// returned after all mappings have been processed.
func processInlineMappings(
	mappings []string,
	wordmap *wordmaps.WordMap,
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
) (wordmap wordmaps.WordMap, errs []error) {
	errs = openAndProcessMapFiles(mapFiles, &wordmap)
	errs = append(errs, processInlineMappings(inlineMappings, &wordmap)...)
	return wordmap, errs
}
