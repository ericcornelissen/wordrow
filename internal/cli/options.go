package cli


const (
  // The flag to output the usage of the program.
  helpFlag = "--help"

  // The flag to output the program version.
  versionFlag = "--version"

  // The flag to enable dry run. If enabled the program won't make any changes to
  // the input files.
  dryRunFlag = "--dry-run"

  // The flag to invert the mapping. If enabled the mapping will be used right-
  // to-left instead of left-to-right.
  invertFlag = "--invert"

  // The alias for the --invert flag.
  invertFlagAlias = "-i"

  // The flag to make the program silent.
  silentFlag = "--silent"

  // The alias for the --silent flag.
  silentFlagAlias = "-s"

  // The flag to make the program verbose.
  verboseFlag = "--verbose"

  // The alias for the --verbose flag.
  verboseFlagAlias = "-v"

  // The option to specify the configuration file to use.
  configOption = "--config"

  // The alias for the --config option.
  configOptionAlias = "-c"

  // The option to specify a mapping file.
  mapfileOption = "--map-file"

  // The alias for the --map-file option.
  mapfileOptionAlias = "-M"

  // TODO
  mappingOption = "--tmp"

  // TODO
  mappingOptionAlias = "-t"
)
