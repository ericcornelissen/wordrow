package main

import (
	"os"

	"github.com/ericcornelissen/wordrow/internal/cli"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
)

func run(args *cli.Arguments) (errors []error) {
	mapping, errs := getMapping(args.MapFiles, args.Mappings)
	if check(&errors, errs) && args.Strict {
		return errors
	}

	if args.Invert {
		mapping = invert(mapping)
	}

	filePaths, errs := fs.ResolveGlobs(args.InputFiles...)
	if check(&errors, errs) && args.Strict {
		return errors
	}

	if !args.DryRun {
		errs = processInputFiles(filePaths, mapping)
		check(&errors, errs)
	}

	return errors
}

func check(errors *[]error, errs []error) bool {
	*errors = append(*errors, errs...)
	return len(*errors) > 0
}

func setLogLevel(args *cli.Arguments) {
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
		setLogLevel(&args)

		errs := run(&args)
		for _, err := range errs {
			logger.Error(err)
		}
	}
}
