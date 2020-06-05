package wordmaps

import "strings"

// TrimSpaceAll trims leading and trailing spaces in all strings. Note that this
// function operates in place.
func trimSpaceAll(v []string) {
	for i, s := range v {
		v[i] = strings.TrimSpace(s)
	}
}

// AnyString returns true if at least one item in the list fulfills the
// condition and false otherwise.
func anyString(v []string, condition func(string) bool) bool {
	for _, s := range v {
		if condition(s) {
			return true
		}
	}

	return false
}

// IsEmptyString returns true if the string is empty and false otherwise.
func isEmptyString(s string) bool {
	return s == ""
}

// Parse a single row of a CSV file and add it to the WordMap.
//
// The error will be set if the row has an unexpected format, for example an
// incorrect number of columns.
func parseRow(row string, wm *WordMap) error {
	rowValues := strings.Split(row, ",")
	if len(rowValues) < 2 {
		return &parseError{"Unexpected row format", row}
	}

	trimSpaceAll(rowValues)
	if anyString(rowValues, isEmptyString) {
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
