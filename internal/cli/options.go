package cli

// Option is a type representing a CLI argument option or flag.
type option struct {
	// The full version of the Option.
	name string

	// The alias for the Option.
	alias string
}

var (
	// The flag to output the usage of the program.
	helpFlag = option{
		name: "--help",
	}

	// The flag to output the program version.
	versionFlag = option{
		name: "--version",
	}

	// The flag to enable dry run. If enabled the program won't make any changes
	// to the input files.
	dryRunFlag = option{
		name: "--dry-run",
	}

	// The flag to invert the mapping. If enabled the mapping will be used right-
	// to-left instead of left-to-right.
	invertFlag = option{
		name:  "--invert",
		alias: "-i",
	}

	// The flag to make the program silent.
	silentFlag = option{
		name:  "--silent",
		alias: "-s",
	}

	// The flag to make the program verbose.
	verboseFlag = option{
		name:  "--verbose",
		alias: "-v",
	}

	// The option to specify a mapping file.
	mapfileOption = option{
		name:  "--map-file",
		alias: "-M",
	}

	// The option to specify a single mapping from the CLI.
	mappingOption = option{
		name:  "--map",
		alias: "-m",
	}
)
