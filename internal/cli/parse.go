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
func parseOption(arg string, arguments *Arguments) (argumentContext, error) {
  newState := contextInputFile

  switch arg {
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
      newState = contextConfigFile
      logger.Warningf("The %s argument is not yet supported", arg)
    case mapfileOption, mapfileOptionAlias:
      newState = contextMapFile
    default:
      return newState, errors.New("Unknown option")
  }

  return newState, nil
}

// Parse a slice of arguments (e.g. `os.Args`) into an Arguments instance.
func ParseArgs(args []string) (bool, Arguments) {
  var arguments Arguments

  if len(args) == 1 {
    printUsage()
    return false, arguments
  }

  if helpOption == args[1] {
    printUsage()
    return false, arguments
  }

  if versionOption == args[1] {
    printVersion()
    return false, arguments
  }

  context := contextInputFile
  for i := 1; i < len(args); i++ {
    arg := args[i]
    if "-" == arg[:1] {
      if context == contextInputFile {
        newContext, err := parseOption(arg, &arguments)
        if err != nil {
          logger.Errorf("Unknown option '%s'. Use %s for help", arg, helpOption)
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
