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

// Get an argContext as a human readable string.
func (context argContext) String() string {
	names := []string{
		"Unknown",
		fmt.Sprintf("%s/%s", mapfileOption.name, mapfileOption.alias),
		fmt.Sprintf("%s/%s", mappingOption.name, mappingOption.alias),
	}

	return names[context]
}
