package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/ericcornelissen/wordrow/internal/cli"
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
	"github.com/ericcornelissen/wordrow/internal/replacer"
)

// Check if the program received input from STDIN.
//
// based on: https://stackoverflow.com/a/38612652
func hasStdin() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	return (fi.Mode() & os.ModeNamedPipe) != 0
}

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
		err = processInputFiles(filePaths, &wordmap)
	}

	return err
}

func runOnStdin(args cli.Arguments, input string) error {
	wm, err := getWordMap(args.MapFiles, args.Mappings)
	if err != nil {
		return err
	}

	if args.Invert {
		wm.Invert()
	}

	fixedInput := replacer.ReplaceAll(input, wm)
	os.Stdout.WriteString(fixedInput)

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

		err := run(args)
		if err != nil {
			logger.Error(err)
		}
	}
}
