package cli

// Option is a type representing a CLI argument option or flag.
type Option struct {
	// The full version of the Option.
	name string

	// The alias for the Option.
	alias string
}

var (
	// The flag to output the usage of the program.
	helpFlag = Option{
		name: "--help",
	}

	// The flag to output the program version.
	versionFlag = Option{
		name: "--version",
	}

	// The flag to enable dry run. If enabled the program won't make any changes
	// to the input files.
	dryRunFlag = Option{
		name: "--dry-run",
	}

	// The flag to invert the mapping. If enabled the mapping will be used right-
	// to-left instead of left-to-right.
	invertFlag = Option{
		name:  "--invert",
		alias: "-i",
	}

	// The flag to make the program silent.
	silentFlag = Option{
		name:  "--silent",
		alias: "-s",
	}

	// The flag to make the program verbose.
	verboseFlag = Option{
		name:  "--verbose",
		alias: "-v",
	}

	// The option to specify a mapping file.
	mapfileOption = Option{
		name:  "--map-file",
		alias: "-M",
	}

	// The option to specify a single mapping from the CLI.
	mappingOption = Option{
		name:  "--map",
		alias: "-m",
	}
)
