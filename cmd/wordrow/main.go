package main

import (
	"bufio"
	"os"

	"github.com/ericcornelissen/wordrow/internal/cli"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
)

func run(args *cli.Arguments) (errors, warnings []error) {
	if hasStdin() {
		logger.SetLogLevel(logger.FATAL)
		errors, warnings = runOnStdin(args)
	} else {
		setLogLevel(args)
		errors, warnings = runOnFiles(args)
	}

	return errors, warnings
}

func runOnFiles(args *cli.Arguments) (errors, warnings []error) {
	mapping, errs := getMapping(args)
	if check(&warnings, errs) && args.Strict {
		return nil, warnings
	}

	filePaths, errs := fs.ResolveGlobs(args.InputFiles...)
	if check(&warnings, errs) && args.Strict {
		return nil, errs
	}

	if !args.DryRun {
		errs = processInputFiles(filePaths, mapping)
		check(&errors, errs)
	}

	return errors, warnings
}

func runOnStdin(args *cli.Arguments) (errors, warnings []error) {
	mapping, errs := getMapping(args)
	if check(&warnings, errs) && args.Strict {
		return nil, warnings
	}

	readWriter := bufio.NewReadWriter(
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout),
	)

	err := processStdin(readWriter, mapping)
	if err != nil {
		errors = append(errors, err)
	}

	return errors, warnings
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

func exit(errors, warnings []error, strict bool) {
	if len(errors) > 0 {
		os.Exit(runtimeErrorExitCode)
	}

	if len(warnings) > 0 && strict {
		os.Exit(runtimeWarningExitCode)
	}

	os.Exit(0)
}

func main() {
	shouldRun, args := cli.ParseArgs(os.Args)
	if args.Version {
		printVersion()
	}

	if !shouldRun && !hasStdin() {
		os.Exit(missingArgumentExitCode)
	}

	errors, warnings := run(&args)
	exit(errors, warnings, args.Strict)
}
