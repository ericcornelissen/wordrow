/*
Package cli provides a single function that can be used to parse the command
line argument for the wordrow program. This will provide a custom struct
`Arguments` which specifies the configuration for the program run.

	import "os"

	func main() {
		shouldRun, args := ParseArgs(os.Args)
		...
	}
*/
package cli

import (
	"fmt"

	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/errors"
	"github.com/ericcornelissen/wordrow/internal/logger"
)

// Check if any arguments were provided to the program.
func noArgumentsProvided(args []string) bool {
	return len(args) == 1
}

// Parse an option argument as a (set of) flag(s).
func parseArgumentAsAlias(
	alias string,
	arguments *Arguments,
) (newContext argContext, err error) {
	for _, char := range alias[1:] {
		option := fmt.Sprintf("-%c", char)
		newContext, err = parseArgumentAsOption(option, arguments)
	}

	return newContext, err
}

// Parse an option argument and get the new argument context.
func parseArgumentAsOption(
	option string,
	arguments *Arguments,
) (argContext, error) {
	newContext := contextDefault
	switch option {
	case helpFlag.name:
		arguments.help = true
	case versionFlag.name:
		arguments.Version = true

	// Flags
	case dryRunFlag.name:
		arguments.DryRun = true
	case invertFlag.name, invertFlag.alias:
		arguments.Invert = true
	case silentFlag.name, silentFlag.alias:
		arguments.Silent = true
	case verboseFlag.name, verboseFlag.alias:
		arguments.Verbose = true
	case strictFlag.name, strictFlag.alias:
		arguments.Strict = true

	// Options
	case mapfileOption.name, mapfileOption.alias:
		newContext = contextMapFile
	case mappingOption.name, mappingOption.alias:
		newContext = contextMapping
	default:
		return newContext, errors.Newf("Unknown option '%s'. Use %s for help", option, helpFlag)
	}

	return newContext, nil
}

// Parse a single argument as an option or flag.
//
// The function sets the error if the argument could not be parsed (in the
// provided context).
func doParseOneOption(
	arg string,
	arguments *Arguments,
) (newContext argContext, err error) {
	if equalsSplit := stringsx.Split(arg, "="); len(equalsSplit) > 1 {
		option, value := equalsSplit[0], stringsx.Join(equalsSplit[1:], "=")
		err = doParseProgramArguments([]string{option, value}, arguments)
		return contextDefault, err
	}

	if stringsx.HasPrefix(arg, "--") {
		newContext, err = parseArgumentAsOption(arg, arguments)
	} else {
		newContext, err = parseArgumentAsAlias(arg, arguments)
	}

	return newContext, err
}

// Parse a single argument as a value or option/flag.
//
// The function sets the error if the argument could not be parsed (in the
// provided context).
func doParseOneArgument(
	arg string,
	context argContext,
	arguments *Arguments,
) (newContext argContext, err error) {
	if context == contextDefault && stringsx.HasPrefix(arg, "-") {
		newContext, err = doParseOneOption(arg, arguments)
	} else {
		context.parseValue(arg, arguments)
		newContext = contextDefault
	}

	return newContext, err
}

// Parse a slice of arguments that contains at least one program argument.
//
// The function sets the error if there is any issue with the provided
// arguments.
func doParseProgramArguments(args []string, arguments *Arguments) error {
	context := contextDefault
	for _, arg := range args {
		newContext, err := doParseOneArgument(arg, context, arguments)
		if err != nil {
			return err
		}

		context = newContext
	}

	if context != contextDefault {
		return errors.New("More arguments expected")
	}

	return nil
}

// ParseArgs parses a list of arguments (e.g. `os.Args`) into an Arguments
// instance.
func ParseArgs(args []string) (run bool, arguments Arguments) {
	err := doParseProgramArguments(args[1:], &arguments)
	if err != nil {
		logger.Fatalf("An error occurred while parsing arguments: %s", err)
		return false, arguments
	}

	if noArgumentsProvided(args) || arguments.help {
		printUsage()
		return false, arguments
	}

	return len(arguments.InputFiles) > 0, arguments
}
