package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/ericcornelissen/wordrow/internal/cli"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
	"github.com/ericcornelissen/wordrow/internal/replace"
)

func run(args cli.Arguments) (errors []error) {
	wordmap, errs := getWordMap(args.MapFiles, args.Mappings)
	if check(&errors, errs) && args.Strict {
		return errs
	}

	if args.Invert {
		wordmap.Invert()
	}

	filePaths, errs := fs.ResolveGlobs(args.InputFiles...)
	if check(&errors, errs) && args.Strict {
		return errs
	}

	if !args.DryRun {
		errs = processInputFiles(filePaths, &wordmap)
		check(&errors, errs)
	}

	return errors
}

func runOnStdin(args cli.Arguments, input string) []error {
	wm, err := getWordMap(args.MapFiles, args.Mappings)
	if err != nil {
		return err
	}

	if args.Invert {
		wm.Invert()
	}

	fixedInput := replace.All(input, wm.Iter())
	os.Stdout.WriteString(fixedInput)

	return nil
}

func check(errors *[]error, errs []error) bool {
	*errors = append(*errors, errs...)
	return len(*errors) > 0
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

	if hasStdin() {
		logger.SetLogLevel(logger.FATAL)

		var sb strings.Builder

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			sb.WriteString(line)
			sb.WriteByte('\n')
		}

		input := sb.String()
		err := runOnStdin(args, input)
		if err != nil {
			panic(err)
		}
	} else if shouldRun {
		setLogLevel(args)

		errs := run(args)
		for _, err := range errs {
			logger.Error(err)
		}
	}
}
