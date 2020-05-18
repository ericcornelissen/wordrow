package main

import "os"
import "strings"

import "github.com/ericcornelissen/wordrow/internal/cli"
import "github.com/ericcornelissen/wordrow/internal/errors"
import "github.com/ericcornelissen/wordrow/internal/fs"
import "github.com/ericcornelissen/wordrow/internal/logger"
import "github.com/ericcornelissen/wordrow/internal/replacer"
import "github.com/ericcornelissen/wordrow/internal/wordmaps"

func getWordMap(
	mapFilesPaths []string,
	cliMappings []string,
) (wm wordmaps.WordMap, err error) {
	for _, mapFilePath := range mapFilesPaths {
		format := ""

		values := strings.Split(mapFilePath, ":")
		if len(values) > 1 {
			mapFilePath = strings.Join(values[:len(values)-1], ":")
			format = values[len(values)-1]
		}

		logger.Debugf("Reading '%s'", mapFilePath)
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

func run(args cli.Arguments) error {
	wm, err := getWordMap(args.MapFiles, args.Mappings)
	if err != nil {
		return err
	}

	if args.Invert {
		wm.Invert()
	}

	inputFiles, err := fs.ReadFiles(args.InputFiles)
	if err != nil {
		return err
	}

	for _, file := range inputFiles {
		logger.Debugf("Processing '%s' as input file", file.Path)
		fixedFileData := replacer.ReplaceAll(file.Content, wm)

		if !args.DryRun {
			fs.WriteFile(file.Path, fixedFileData)
		} else {
			logger.Printf("Before:\n-------\n%s\n", file.Content)
			logger.Printf("After:\n------\n%s", fixedFileData)
		}
	}

	return nil
}

func setLogLevel(args cli.Arguments) {
	if args.Silent {
		logger.SetLogLevel(logger.ERROR)
	} else if args.Verbose {
		logger.SetLogLevel(logger.DEBUG)
	}
}

func main() {
	shouldRun, args := cli.ParseArgs(os.Args)
	if shouldRun {
		setLogLevel(args)

		err := run(args)
		if err != nil {
			logger.Error(err)
		}
	}
}
