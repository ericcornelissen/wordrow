package wordmaps

import "github.com/ericcornelissen/wordrow/internal/strings"

// Parse a single row of a CSV file and add it to the WordMap.
//
// The error will be set if the row has an unexpected format, for example an
// incorrect number of columns.
func parseRow(row string, wm *WordMap) error {
	rowValues := strings.Split(row, ",")
	if len(rowValues) < 2 {
		return &parseError{"Unexpected row format", row}
	}

	strings.TrimSpaceAll(rowValues)
	if strings.Any(rowValues, strings.IsEmpty) {
		return &parseError{"Missing value", row}
	}

	last := len(rowValues) - 1
	wm.AddMany(rowValues[0:last], rowValues[last])
	return nil
}

// Parse a Comma Separated Values, CSV, file into a WordMap.
//
// The error will be set if any error occured while parsing the CSV file.
func parseCsvFile(rawFileData *string) (WordMap, error) {
	var wm WordMap

	lines := strings.Split(*rawFileData, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.TrimSpace(line) == "" {
			continue
		}

		err := parseRow(line, &wm)
		if err != nil {
			return wm, err
		}
	}

	return wm, nil
}
