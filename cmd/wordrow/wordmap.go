package main

import (
	"io/ioutil"

	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
	"github.com/ericcornelissen/wordrow/internal/mappings"
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

// Add the mapping of the `reader` to the `target`. The `format` argument
// determines how the contents of the file are parsed. This function returns an
// error if either the reading or parsing fails.
func processMapFile(
	reader fs.Reader,
	format string,
	target map[string]string,
) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	content := string(data)
	mapping, err := mappings.ParseFile(&content, format)
	if err != nil {
		return err
	}

	merge(target, mapping)
	return nil
}

// Opens the file provided by the handler and add its mapping to the `mapping`.
// If the file cannot be opened or processing failed the handler returns an
// error.
func openAndProcessMapFileWith(mapping map[string]string) handler {
	return func(fileArgument string) error {
		filePath, format := parseMapFileArgument(fileArgument)

		logger.Debugf("Opening '%s' as a '%s' formatted map file", filePath, format)
		handle, err := fs.OpenFile(filePath, fs.OReadOnly)
		if err != nil {
			return err
		}

		defer handle.Close()

		logger.Debugf("Processing '%s' as a map file", filePath)
		return processMapFile(handle, format, mapping)
	}
}

// Processes the value provided by the handler and add its mapping to the
// `target`. Of the value cannot be parsed as a CSV mapping the handler returns
// an error.
func processInlineMapping(value string, target map[string]string) error {
	mapping, err := mappings.ParseFile(&value, "csv")
	if err != nil {
		return err
	}

	merge(target, mapping)
	return nil
}

// Processes the value provided by the handler and add its mapping to the
// `mapping`. Of the value cannot be parsed as a CSV mapping the handler returns
// an error.
func processInlineMappingWith(mapping map[string]string) handler {
	return func(value string) error {
		logger.Debugf("Processing '%s' as a CLI specified mapping", value)
		return processInlineMapping(value, mapping)
	}
}

// Get a mapping for the specified `mapFiles` and `inlineMappings`. Any error
// that occurs is returned after both have been processed. In case of any error
// the mapping that is returned represents only the arguments that could be
// successfully processed.
func getMapping(
	mapFiles []string,
	inlineMappings []string,
	invertMapping bool,
) (map[string]string, []error) {
	mapping := make(map[string]string)

	errs := forEach(mapFiles, openAndProcessMapFileWith(mapping))
	errs = append(
		errs,
		forEach(inlineMappings, processInlineMappingWith(mapping))...,
	)

	if invertMapping {
		mapping = invert(mapping)
	}

	return mapping, errs
}
