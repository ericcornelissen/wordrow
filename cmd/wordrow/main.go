package main

import (
	"os"

	"github.com/ericcornelissen/wordrow/internal/cli"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
)

func run(args cli.Arguments) error {
	wordmap, err := getWordMap(args.MapFiles, args.Mappings)
	if err != nil {
		return err
	}

	if args.Invert {
		wordmap.Invert()
	}

	filePaths, err := fs.ResolveGlobs(args.InputFiles...)
	if err != nil {
		return err
	}

	if !args.DryRun {
		processFiles(filePaths, &wordmap)
	}

	return err
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

	if args.Version {
		printVersion()
	}

	if shouldRun {
		setLogLevel(args)

		err := run(args)
		if err != nil {
			logger.Error(err)
		}
	}
}
