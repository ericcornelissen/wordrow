package cli

const (
  // The flag to output the usage of the program.
  helpOption = "--help"

  // The flag to output the program version.
  versionOption = "--version"

  // The flag to enable dry run. If enabled the program won't make any changes to
  // the input files.
  dryRunOption = "--dry-run"

  // The flag to invert the mapping. If enabled the mapping will be used right-
  // to-left instead of left-to-right.
  invertOption = "--invert"

  // The alias for the --invert option.
  invertOptionAlias = "-i"

  // The flag to make the program silent.
  silentOption = "--silent"

  // The alias for the --silent option.
  silentOptionAlias = "-s"

  // The flag to make the program verbose.
  verboseOption = "--verbose"

  // The alias for the --verbose option.
  verboseOptionAlias = "-v"

  // The option to specify the configuration file to use.
  configOption = "--config"

  // The alias for the --config option.
  configOptionAlias = "-c"

  // The option to specify a mapping file.
  mapfileOption = "--map"

  // The alias for the --map option.
  mapfileOptionAlias = "-m"
)
