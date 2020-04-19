package cli

import "github.com/ericcornelissen/wordrow/internal/errors"
import "github.com/ericcornelissen/wordrow/internal/logger"


// The Arguments type represents the configuration of the program from the
// Command-Line Interface (CLI).
type Arguments struct {
  // Flag indicating if the program usage should be displayed.
  help bool

  // Flag indicating if the program version should be displayed.
  version bool

  // Flag indicating if this is a dry run.
  DryRun bool

  // Flag indicating if the mapping should be inverted.
  Invert bool

  // Flag indicating if the program should be silent.
  Silent bool

  // Flag indicating if the program should be verbose.
  Verbose bool

  // The config file.
  ConfigFile string

  // List of files to be processed.
  InputFiles []string

  // List of files that specify a mapping.
  MapFiles []string

  // List of mappings defined in the CLI.
  Mappings []string
}


// Check if any arguments were provided to the program.
func noArgumentsProvided(args []string) bool {
  return len(args) == 1
}

// Check if a certain argument is an option.
func argumentIsOption(arg string) bool {
  return "-" == arg[:1]
}


// Parse an option argument and get the new argument context.
func parseArgumentAsOption(
  option string,
  arguments *Arguments,
) (argContext, error) {
  newContext := contextInputFile
  switch option {
    case helpOption:
      arguments.help = true
    case versionOption:
      arguments.version = true

    // Flags
    case dryRunOption:
      arguments.DryRun = true
    case invertOption, invertOptionAlias:
      arguments.Invert = true
    case silentOption, silentOptionAlias:
      arguments.Silent = true
    case verboseOption, verboseOptionAlias:
      arguments.Verbose = true

    // Options
    case configOption, configOptionAlias:
      newContext = contextConfigFile
      logger.Warningf("The %s argument is not yet supported", option)
    case mapfileOption, mapfileOptionAlias:
      newContext = contextMapFile
    case mappingOption, mappingOptionAlias:
      newContext = contextMapping
    default:
      return newContext, errors.Newf("Unknown option '%s'. Use %s for help", option, helpOption)
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
  } else {
    parseArgumentAsValue(arg, context, arguments)
    return contextInputFile, nil
  }
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


// Parse a slice of arguments (e.g. `os.Args`) into an Arguments instance.
func ParseArgs(args []string) (bool, Arguments) {
  var arguments Arguments
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

  if arguments.version {
    printVersion()
  }

  return len(arguments.InputFiles) > 0, arguments
}
