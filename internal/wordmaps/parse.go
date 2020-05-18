package wordmaps

import "github.com/ericcornelissen/wordrow/internal/errors"

var (
	// List of names considered as MarkDown format.
	md = []string{".md"}

	// List of names considered as CSV format.
	csv = []string{".csv"}
)

// A parse function is a function that takes the contents of a file as a string
// and outputs a WordMap. If the file is not formatted correctly the function
// may output an error.
type parseFunction func(fileContent *string) (WordMap, error)

// Get the parseFunction for a given format.
func getParserForFormat(format string) (parseFunction, error) {
	if contains(md, format) {
		return parseMarkDownFile, nil
	} else if contains(csv, format) {
		return parseCsvFile, nil
	}

	return nil, errors.New("Unknown file type")
}

// Parse a string formatted in a certain way into a WordMap.
//
// The function sets the error if the parsing failed, e.g. when the format is
// unknown or if content is improperly formatted.
func parseFile(content *string, format string, wm *WordMap) error {
	parseFn, err := getParserForFormat(format)
	if err != nil {
		return errors.Newf("Unknown type '%s'", format)
	}

	fileMap, err := parseFn(content)
	if err != nil {
		return err
	}

	wm.AddFrom(fileMap)
	return nil
}
