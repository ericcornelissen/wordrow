package csv

import (
	"bufio"
	"bytes"

	"github.com/ericcornelissen/wordrow/internal/mappings/errors"
)

// Parse a single row of a CSV file and add it to the WordMap.
//
// The error will be set if the row has an unexpected format, for example an
// incorrect number of columns.
func parseRow(row []byte, mapping map[string]string) error {
	rowValuesCount := 2

	rowValues := bytes.Split(row, []byte(","))
	if len(rowValues) < rowValuesCount {
		return errors.NewIncorrectFormat(row)
	}

	for i, v := range rowValues {
		rowValues[i] = bytes.TrimSpace(v)
		if len(rowValues[i]) == 0 {
			return errors.NewMissingValue(row)
		}
	}

	last := len(rowValues) - 1
	to := string(rowValues[last])
	for _, from := range rowValues[0:last] {
		mapping[string(from)] = to
	}
	return nil
}

// Parse parses a Comma Separated Values (CSV) file into a map[string]string.
//
// The error will be set if any error occurred while parsing the CSV file.
func Parse(reader *bufio.Reader) (mapping map[string]string, err error) {
	mapping = make(map[string]string, 1)

	var line []byte
	for ; err == nil; line, _, err = reader.ReadLine() {
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}

		err := parseRow(line, mapping)
		if err != nil {
			return mapping, err
		}
	}

	return mapping, nil
}
