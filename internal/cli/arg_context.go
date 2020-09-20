package cli

import "fmt"

// A custom integer type for an Enum to keep track of the arguments context.
type argContext int

// The Enum used to keep track of the argument context.
const (
	// The context where arguments are interpreted as input files.
	contextDefault argContext = iota

	// The context where arguments are interpreted as mapping files.
	contextMapFile

	// The context where arguments are interpreted as a mapping.
	contextMapping
)

// Parse an argument that is not in option within a certain argument context.
func (context argContext) parseValue(value string, arguments *Arguments) {
	switch context {
	case contextDefault:
		arguments.InputFiles = append(arguments.InputFiles, value)
	case contextMapFile:
		arguments.MapFiles = append(arguments.MapFiles, value)
	case contextMapping:
		arguments.Mappings = append(arguments.Mappings, value)
	}
}

// Get an argContext as a human readable string.
func (context argContext) String() string {
	template := "%s/%s"
	names := []string{
		"Unknown",
		fmt.Sprintf(template, mapfileOption.name, mapfileOption.alias),
		fmt.Sprintf(template, mappingOption.name, mappingOption.alias),
	}

	return names[context]
}
