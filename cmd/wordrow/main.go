package main

import (
	"os"

	"github.com/ericcornelissen/wordrow/internal/cli"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
)

func run(args cli.Arguments) (errs []error) {
	wordmap, err := getWordMap(args.MapFiles, args.Mappings)
	if err != nil {
		return []error{err}
	}

	if args.Invert {
		wordmap.Invert()
	}

	filePaths, err := fs.ResolveGlobs(args.InputFiles...)
	if err != nil {
		return []error{err}
	}

	if !args.DryRun {
		errs = processInputFiles(filePaths, &wordmap)
	}

	return errs
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

		errs := run(args)
		if len(errs) > 0 {
			for _, err := range errs {
				logger.Error(err)
			}
		}
	}
}
