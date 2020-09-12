// Package wordmaps provides two structures for functionality to parse files
// into a map[string]string. The supported formats are:
// - CSV
// - MarkDown
package wordmaps

func addMany(target map[string]string, froms []string, to string) {
	for _, from := range froms {
		target[from] = to
	}
}

// ParseFile parses a file defining a string-to-string mapping under a given
// format.
func ParseFile(content *string, format string) (map[string]string, error) {
	return parseFile(content, format)
}
