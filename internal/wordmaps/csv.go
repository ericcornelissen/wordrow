package wordmaps

import (
	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/errors"
)

// Parse a single row of a CSV file and add it to the WordMap.
//
// The error will be set if the row has an unexpected format, for example an
// incorrect number of columns.
func parseRow(row string, wm *WordMap) error {
	rowValues := stringsx.Split(row, ",")
	if len(rowValues) < 2 {
		return errors.Newf(incorrectFormat, row)
	}

	rowValues = stringsx.MapAll(rowValues, stringsx.TrimSpace)
	if stringsx.Any(rowValues, stringsx.IsEmpty) {
		return errors.Newf(missingValue, row)
	}

	last := len(rowValues) - 1
	wm.AddMany(rowValues[0:last], rowValues[last])
	return nil
}

// Parse a Comma Separated Values (CSV) file into a WordMap.
//
// The error will be set if any error occurred while parsing the CSV file.
func parseCsvFile(rawFileData *string) (WordMap, error) {
	var wm WordMap

	lines := stringsx.Split(*rawFileData, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if stringsx.TrimSpace(line) == "" {
			continue
		}

		err := parseRow(line, &wm)
		if err != nil {
			return wm, err
		}
	}

	return wm, nil
}
