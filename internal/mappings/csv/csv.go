package csv

import (
	"bufio"
	"bytes"

	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/errors"
	"github.com/ericcornelissen/wordrow/internal/mappings/common"
)

// Parse a single row of a CSV file and add it to the `mapping`.
//
// The error will be set if the row has an unexpected format, for example an
// incorrect number of columns.
func parseRow(row string, mapping map[string]string) error {
	rowValuesCount := 2

	rowValues := stringsx.Split(row, ",")
	if len(rowValues) < rowValuesCount {
		return errors.Newf(common.IncorrectFormat, row)
	}

	rowValues = stringsx.MapAll(rowValues, stringsx.TrimSpace)
	if stringsx.Any(rowValues, stringsx.IsEmpty) {
		return errors.Newf(common.MissingValue, row)
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
func _parseRow(row []byte, wm *ByteMap) error {
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

// Parse a Comma Separated Values (CSV) file into a WordMap.
//
// The error will be set if any error occurred while parsing the CSV file.
func _parseCsvFile(rawFileData []byte) (ByteMap, error) {
	var wm ByteMap

	lines := bytes.Split(rawFileData, []byte("\n"))
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}

		err := _parseRow(line, &wm)
		if err != nil {
			return wm, err
		}
	}

	return wm, nil
}

func __parseCsvFile(r *bufio.Reader) (ByteMap, error) {
	var wm ByteMap

	var line []byte
	var err error
	for ; err == nil; line, _, err = r.ReadLine() {
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}

		err := _parseRow(line, &wm)
		if err != nil {
			return wm, err
		}
	}

	return wm, nil
}
