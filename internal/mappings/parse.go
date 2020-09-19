// Package mappings provides two structures for functionality to parse files
// into a map[string]string. The supported formats are:
// - CSV
// - MarkDown
package mappings

import (
	"regexp"

	"github.com/ericcornelissen/wordrow/internal/errors"
	"github.com/ericcornelissen/wordrow/internal/mappings/csv"
	"github.com/ericcornelissen/wordrow/internal/mappings/markdown"
)

var (
	// Regular expression of names considered as MarkDown format.
	mdPattern = regexp.MustCompile(`(?i)\.?(md(te?xt)?|markdown|mdown|mkdown|mkd|mdwn|mkdn)`)

	// Regular expression of names considered as CSV format.
	csvPattern = regexp.MustCompile(`(?i)\.?csv`)
)

// A parse function is a function that takes the contents of a file as a string
// and outputs a map[string]string. If the file is not formatted correctly the
// function may output an error.
type parseFunction func(fileContent *string) (map[string]string, error)

// Get the parseFunction for a given format.
func getParserForFormat(format string) (parseFunction, error) {
	if mdPattern.MatchString(format) {
		return markdown.Parse, nil
	} else if csvPattern.MatchString(format) {
		return csv.Parse, nil
	}

	return nil, errors.Newf("Unknown format '%s'", format)
}

// ParseFile parses a file formatted in a certain way into a map[string]string.
// The function sets the error if the parsing failed, e.g. when the format is
// unknown or if content is improperly formatted.
func ParseFile(content *string, format string) (map[string]string, error) {
	parseFn, err := getParserForFormat(format)
	if err != nil {
		return nil, err
	}

	mapping, err := parseFn(content)
	if err != nil {
		return nil, err
	}

	return mapping, nil
}
