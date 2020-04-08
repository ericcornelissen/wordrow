package cli

import "github.com/ericcornelissen/wordrow/internal/errors"
import "github.com/ericcornelissen/wordrow/internal/logger"


// A custom integer type for Enum to keep track of the arguments context.
type argumentContext int

// The Enum used to keep track of the argument context.
const (
  // The context where arguments are interpreted as input files.
  contextInputFile argumentContext = iota

  // The context where arguments are interpreted as a configuration file.
  contextConfigFile

  // The context where arguments are interpreted as mapping files.
  contextMapFile

  // The context when parsing finished early.
  contextDone
)


// The Arguments type represents the configuration of the program from the
// Command-Line Interface (CLI).
type Arguments struct {
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
}


// Check if any arguments were provided to the program.
func noArgumentsProvided(args []string) bool {
  return len(args) == 1
}

// Check if a certain argument is an option.
func argumentIsOption(arg string) bool {
  return "-" == arg[:1]
}


// Parse an argument that is not in option within a certain argument context.
func parseArgument(arg string, context argumentContext, arguments *Arguments) {
  switch context {
    case contextInputFile:
      arguments.InputFiles = append(arguments.InputFiles, arg)
    case contextConfigFile:
      arguments.ConfigFile = arg
    case contextMapFile:
      arguments.MapFiles = append(arguments.MapFiles, arg)
  }
}

// Parse an option argument and get the new argument context.
func parseOption(option string, arguments *Arguments) (argumentContext, error) {
  newContext := contextInputFile

  switch option {
    case helpOption:
      printUsage()
      newContext = contextDone
    case versionOption:
      printVersion()
      newContext = contextDone

    // Flags
    case dryRunOption:
      arguments.DryRun = true
    case invertOption, invertOptionAlias:
      arguments.Invert = true
    case silentOption, silentOptionAlias:
      arguments.Silent = true
      logger.Warningf("The %s argument is not yet supported", option)
    case verboseOption, verboseOptionAlias:
      arguments.Verbose = true
      logger.Warningf("The %s argument is not yet supported", option)

    // Options
    case configOption, configOptionAlias:
      newContext = contextConfigFile
      logger.Warningf("The %s argument is not yet supported", option)
    case mapfileOption, mapfileOptionAlias:
      newContext = contextMapFile
    default:
      return newContext, errors.Newf("Unknown option '%s'", option)
  }

  return newContext, nil
}


// Parse a slice of arguments (e.g. `os.Args`) into an Arguments instance.
func ParseArgs(args []string) (bool, Arguments) {
  var arguments Arguments

  if noArgumentsProvided(args) {
    printUsage()
    return false, arguments
  }

  context := contextInputFile
  for _, arg := range args[1:] {
    if argumentIsOption(arg) {
      if context == contextInputFile {
        newContext, err := parseOption(arg, &arguments)
        if err != nil {
          logger.Errorf("Unknown option '%s'. Use %s for help", arg, helpOption)
          return false, arguments
        } else if newContext == contextDone {
          return false, arguments
        } else {
          context = newContext
        }
      } else {
        // TODO missing value for earlier argument
        logger.Warning("missing value")
        return false, arguments
      }
    } else {
      parseArgument(arg, context, &arguments)
      context = contextInputFile
    }
  }

  if context != contextInputFile {
    // TODO arguments incomplete
    logger.Warning("incomplete")
    return false, arguments
  }

  return true, arguments
}
