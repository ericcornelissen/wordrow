package csv

import (
	"bufio"
	"bytes"

	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/common"
	"github.com/ericcornelissen/wordrow/internal/mappings/errors"
)

// Parse a single row of a CSV file and add it to the `mapping`.
//
// The error will be set if the row has an unexpected format, for example an
// incorrect number of columns.
func parseRow(row string, mapping map[string]string) error {
	rowValuesCount := 2

	rowValues := stringsx.Split(row, ",")
	if len(rowValues) < rowValuesCount {
		return errors.NewIncorrectFormat(row)
	}

	rowValues = stringsx.MapAll(rowValues, stringsx.TrimSpace)
	if stringsx.Any(rowValues, stringsx.IsEmpty) {
		return errors.NewMissingValue(row)
	}

	last := len(rowValues) - 1
	additionalMappings := common.MapFrom(rowValues[0:last], rowValues[last])
	common.MergeMaps(mapping, additionalMappings)
	return nil
}

// Parse a Comma Separated Values (CSV) file into a map[string]string.
//
// The error will be set if any error occurred while parsing the CSV file.
func Parse(rawFileData *string) (map[string]string, error) {
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

// Parse a single row of a CSV file and add it to the WordMap.
//
// The error will be set if the row has an unexpected format, for example an
// incorrect number of columns.
func parseRowAsBytes(row []byte, wm *ByteMap) error {
	rowValuesCount := 2

	rowValues := bytes.Split(row, []byte(","))
	if len(rowValues) < rowValuesCount {
		return errors.Newf("Unexpected row format (in '%s')", row)
	}

	for _, v := range rowValues {
		if len(bytes.TrimSpace(v)) == 0 {
			return errors.Newf("Missing value (in '%s')", row)
		}
	}

	last := len(rowValues) - 1
	wm.AddMany(rowValues[0:last], rowValues[last])
	return nil
}

// ParseReader parses a Comma Separated Values (CSV) file from a reader into a
// ByteMap.
//
// The error will be set if any error occurred while parsing the CSV file.
func ParseReader(r *bufio.Reader) (ByteMap, error) {
	var wm ByteMap

	var line []byte
	var err error
	for ; err == nil; line, _, err = r.ReadLine() {
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}

		err := parseRowAsBytes(line, &wm)
		if err != nil {
			return wm, err
		}
	}

	return wm, nil
}
