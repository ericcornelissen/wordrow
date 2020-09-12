package wordmaps

import (
	"regexp"

	"github.com/ericcornelissen/wordrow/internal/errors"
)

var (
	// Regular expression of names considered as MarkDown format.
	md = regexp.MustCompile(`(?i)\.?(md(te?xt)?|markdown|mdown|mkdown|mkd|mdwn|mkdn)`)

	// Regular expression of names considered as CSV format.
	csv = regexp.MustCompile(`(?i)\.?csv`)
)

// A parse function is a function that takes the contents of a file as a string
// and outputs a map[string]string. If the file is not formatted correctly the
// function may output an error.
type parseFunction func(fileContent *string) (map[string]string, error)

// Get the parseFunction for a given format.
func getParserForFormat(format string) (parseFunction, error) {
	if md.MatchString(format) {
		return parseMarkDownFile, nil
	} else if csv.MatchString(format) {
		return parseCsvFile, nil
	}

	return nil, errors.Newf(unknownFormat, format)
}

// Parse a string formatted in a certain way into a map[string]string.
//
// The function sets the error if the parsing failed, e.g. when the format is
// unknown or if content is improperly formatted.
func parseFile(content *string, format string) (map[string]string, error) {
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
