package main

import "strings"

import "github.com/ericcornelissen/wordrow/internal/errors"
import "github.com/ericcornelissen/wordrow/internal/fs"
import "github.com/ericcornelissen/wordrow/internal/logger"
import "github.com/ericcornelissen/wordrow/internal/wordmaps"

// Parse a --map-file argument into its component parts.
func parseWordMapArgument(argument string) (path string, format string) {
	values := strings.Split(argument, ":")
	if len(values) > 1 {
		path = strings.Join(values[:len(values)-1], ":")
		format = values[len(values)-1]
	} else {
		path = argument
	}

	return path, format
}

// Add the mappings defined in the specified map files to the WordMap.
func processMapFiles(mapFilesArgs []string, wm *wordmaps.WordMap) error {
	for _, mapFileArg := range mapFilesArgs {
		logger.Debugf("Parsing argument '%s'", mapFileArg)
		file, format := parseWordMapArgument(mapFileArg)

		logger.Debugf("Reading '%s'", file)
		mapFile, err := fs.ReadFile(file)
		if err != nil {
			return err
		}

		if format == "" {
			logger.Debugf("Processing '%s' as mapping file", mapFile.Path)
			err = wm.AddFile(&mapFile.Content, mapFile.Ext)
		} else {
			logger.Debugf("Processing '%s' as %s mapping file", mapFile.Path, format)
			err = wm.AddFile(&mapFile.Content, format)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// Add the mappings defined in the CLI to the WordMap.
func processMappings(mappings []string, wm *wordmaps.WordMap) error {
	for _, mapping := range mappings {
		logger.Debugf("Processing CLI specified mapping: '%s'", mapping)
		values := strings.Split(mapping, ",")
		if len(values) != 2 {
			return errors.Newf("Incorrect mapping from CLI: '%s'", mapping)
		}

		wm.AddOne(values[0], values[1])
	}

	return nil
}

// Get a WordMap based on the specified mappings and map files.
func getWordMap(
	mapFilesArgs []string,
	cliMappings []string,
) (wordmap wordmaps.WordMap, err error) {
	logger.Debug("Processing CLI specified map files...")
	err = processMapFiles(mapFilesArgs, &wordmap)

	logger.Debug("Processing CLI specified mappings...")
	err = processMappings(cliMappings, &wordmap)

	return wordmap, err
}
