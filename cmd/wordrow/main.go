package main

import (
	"bufio"
	"os"

	"github.com/ericcornelissen/wordrow/internal/cli"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
)

func run(args *cli.Arguments) (errors []error) {
	mapping, errs := getMapping(args.MapFiles, args.Mappings, args.Invert)
	if check(&errors, errs) && args.Strict {
		return errors
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

func runOnStdin(args *cli.Arguments) (errors []error) {
	mapping, errs := getMapping(args.MapFiles, args.Mappings, args.Invert)
	if check(&errors, errs) && args.Strict {
		return errors
	}

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	err := processBuffer(scanner, writer, mapping)
	if err != nil {
		errors = append(errors, err)
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

	if !shouldRun {
		return
	}

	if hasStdin() {
		logger.SetLogLevel(logger.FATAL)
		errs := runOnStdin(&args)
		panic(errs)
	} else {
		setLogLevel(&args)
		errs := run(&args)
		for _, err := range errs {
			logger.Error(err)
		}
	}
}
