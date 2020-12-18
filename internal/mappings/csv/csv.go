package csv

import (
	"bufio"
	"bytes"

	"github.com/ericcornelissen/wordrow/internal/mappings/common"
	"github.com/ericcornelissen/wordrow/internal/mappings/errors"
)

// Byte-slice representing a comma (',').
var comma = []byte{','}

// Parse a single row of a CSV file and add it to the map[string]string.
//
// The error will be set if the row has an unexpected format, for example an
// incorrect number of columns.
func parseRow(row []byte, mapping map[string]string) error {
	rowValuesCount := 2

	rowValues := bytes.Split(row, comma)
	if len(rowValues) < rowValuesCount {
		return errors.NewIncorrectFormat(row)
	}

	rowValues, err := common.TrimValues(rowValues)
	if err != nil {
		return errors.NewMissingValue(row)
	}

	common.AddValuesToMap(mapping, rowValues)
	return nil
}

// Parse a Comma Separated Values (CSV) file into a map[string]string.
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
