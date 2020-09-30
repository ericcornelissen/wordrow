package main

import (
	"bufio"
	"os"

	"github.com/ericcornelissen/wordrow/internal/cli"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
)

func run(args *cli.Arguments) {
	if hasStdin() {
		logger.SetLogLevel(logger.FATAL)
		errors := runOnStdin(args)
		if errors != nil {
			panic(errors)
		}
	} else {
		setLogLevel(args)
		errors := runOnFiles(args)
		for _, err := range errors {
			logger.Error(err)
		}
	}
}

func runOnFiles(args *cli.Arguments) (errors []error) {
	mapping, errs := getMapping(args)
	if check(&errors, errs) && args.Strict {
		return errs
	}

	filePaths, errs := fs.ResolveGlobs(args.InputFiles...)
	if check(&errors, errs) && args.Strict {
		return errs
	}

	if !args.DryRun {
		errs = processInputFiles(filePaths, mapping)
		check(&errors, errs)
	}

	return errors
}

func runOnStdin(args *cli.Arguments) (errors []error) {
	mapping, errs := getMapping(args)
	if check(&errors, errs) && args.Strict {
		return errors
	}

	readWriter := bufio.NewReadWriter(
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout),
	)

	err := processStdin(readWriter, mapping)
	if err != nil {
		errors = append(errors, err)
	}

	return errors
}

func check(errors *[]error, newErrors []error) bool {
	*errors = append(*errors, newErrors...)
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

	if shouldRun || hasStdin() {
		run(&args)
	}
}
