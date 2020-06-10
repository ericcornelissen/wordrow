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
	"strings"

	"github.com/ericcornelissen/wordrow/internal/errors"
	"github.com/ericcornelissen/wordrow/internal/logger"
)

// Check if any arguments were provided to the program.
func noArgumentsProvided(args []string) bool {
	return len(args) == 1
}

// Check if a certain argument is an option.
func argumentIsOption(arg string) bool {
	return strings.HasPrefix(arg, "-")
}

// Parse an option argument and get the new argument context.
func parseArgumentAsOption(
	option string,
	arguments *Arguments,
) (argContext, error) {
	newContext := contextInputFile
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

	// Options
	case configOption.name, configOption.alias:
		newContext = contextConfigFile
		logger.Warningf("The %s argument is not yet supported", option)
	case mapfileOption.name, mapfileOption.alias:
		newContext = contextMapFile
	case mappingOption.name, mappingOption.alias:
		newContext = contextMapping
	default:
		return newContext, errors.Newf("Unknown option '%s'. Use %s for help", option, helpFlag)
	}

	return newContext, nil
}

// Parse an argument that is not in option within a certain argument context.
func parseArgumentAsValue(
	value string,
	context argContext,
	arguments *Arguments,
) {
	switch context {
	case contextInputFile:
		arguments.InputFiles = append(arguments.InputFiles, value)
	case contextConfigFile:
		arguments.ConfigFile = value
	case contextMapFile:
		arguments.MapFiles = append(arguments.MapFiles, value)
	case contextMapping:
		arguments.Mappings = append(arguments.Mappings, value)
	}
}

// Parse a single argument, value or option.
//
// The function sets the error if the argument could not be parsed (in the
// provided context).
func doParseOneArgument(
	arg string,
	context argContext,
	arguments *Arguments,
) (argContext, error) {
	if argumentIsOption(arg) {
		if context != contextInputFile {
			return context, errors.Newf("Missing value for %s option", context)
		}

		newContext, err := parseArgumentAsOption(arg, arguments)
		if err != nil {
			return context, err
		}

		return newContext, nil
	}

	parseArgumentAsValue(arg, context, arguments)
	return contextInputFile, nil
}

// Parse a slice of arguments that contains at least one program argument.
//
// The function sets the error if there is any issue with the provided
// arguments.
func doParseProgramArguments(args []string) (Arguments, error) {
	var arguments Arguments

	context := contextInputFile
	for _, arg := range args {
		newContext, err := doParseOneArgument(arg, context, &arguments)
		if err != nil {
			return arguments, err
		}

		context = newContext
	}

	if context != contextInputFile {
		return arguments, errors.New("More arguments expected")
	}

	return arguments, nil
}

// ParseArgs parses a list of arguments (e.g. `os.Args`) into an Arguments
// instance.
func ParseArgs(args []string) (run bool, arguments Arguments) {
	if noArgumentsProvided(args) {
		printUsage()
		return false, arguments
	}

	arguments, err := doParseProgramArguments(args[1:])
	if err != nil {
		logger.Error(err)
		return false, arguments
	}

	if arguments.help {
		printUsage()
		return false, arguments
	}

	return len(arguments.InputFiles) > 0, arguments
}
