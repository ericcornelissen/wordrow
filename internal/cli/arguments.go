package cli

// The Arguments type represents the configuration of the program from the
// Command-Line Interface (CLI).
type Arguments struct {
	// Flag indicating if the program usage should be displayed.
	help bool

	// Flag indicating if the program version should be displayed.
	Version bool

	// Flag indicating if this is a dry run.
	DryRun bool

	// Flag indicating if the mapping should be inverted.
	Invert bool

	// Flag indicating if the program should be silent.
	Silent bool

	// Flag indicating if the program should be verbose.
	Verbose bool

	// Flag indicating if the program should halt on non-blocking errors.
	Strict bool

	// List of files to be processed.
	InputFiles []string

	// List of files that specify a mapping.
	MapFiles []string

	// List of mappings defined in the CLI.
	Mappings []string
}
