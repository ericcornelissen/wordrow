package mappings

import (
	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/errors"
)

// Parse a single row of a CSV file and add it to the `mapping`.
//
// The error will be set if the row has an unexpected format, for example an
// incorrect number of columns.
func parseRow(row string, mapping map[string]string) error {
	rowValuesCount := 2

	rowValues := stringsx.Split(row, ",")
	if len(rowValues) < rowValuesCount {
		return errors.Newf(incorrectFormat, row)
	}

	rowValues = stringsx.MapAll(rowValues, stringsx.TrimSpace)
	if stringsx.Any(rowValues, stringsx.IsEmpty) {
		return errors.Newf(missingValue, row)
	}

	last := len(rowValues) - 1
	addToMapping(mapping, rowValues[0:last], rowValues[last])
	return nil
}

// Parse a Comma Separated Values (CSV) file into a map[string]string.
//
// The error will be set if any error occurred while parsing the CSV file.
func parseCsvFile(rawFileData *string) (map[string]string, error) {
	mapping := make(map[string]string, 1)

	lines := stringsx.Split(*rawFileData, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if stringsx.TrimSpace(line) == "" {
			continue
		}

		err := parseRow(line, mapping)
		if err != nil {
			return mapping, err
		}
	}

	return mapping, nil
}
