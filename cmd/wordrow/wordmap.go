package main

import "strings"

import "github.com/ericcornelissen/wordrow/internal/errors"
import "github.com/ericcornelissen/wordrow/internal/fs"
import "github.com/ericcornelissen/wordrow/internal/logger"
import "github.com/ericcornelissen/wordrow/internal/wordmaps"

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

func getWordMap(
	mapFilesPaths []string,
	cliMappings []string,
) (wm wordmaps.WordMap, err error) {
	for _, mapFilePath := range mapFilesPaths {
		file, format := parseWordMapArgument(mapFilePath)

		logger.Debugf("Reading '%s'", file)
		mapFile, err := fs.ReadFile(mapFilePath)
		if err != nil {
			return wm, err
		}

		if format == "" {
			logger.Debugf("Processing '%s' as mapping file", mapFile.Path)
			err = wm.AddFile(mapFile)
		} else {
			logger.Debugf("Processing '%s' as %s mapping file", mapFile.Path, format)
			err = wm.AddFileAs(mapFile, format)
		}

		if err != nil {
			return wm, err
		}
	}

	logger.Debug("Processing CLI specified mappings")
	for _, mapping := range cliMappings {
		logger.Debugf("Processing CLI specified mapping: '%s'", mapping)

		values := strings.Split(mapping, ",")
		if len(values) != 2 {
			return wm, errors.Newf("Incorrect mapping from CLI: '%s'", mapping)
		}

		wm.AddOne(values[0], values[1])
	}

	return wm, err
}
