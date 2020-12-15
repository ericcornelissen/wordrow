package markdown

import (
	"bufio"
	"bytes"
	"regexp"

	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/common"
	"github.com/ericcornelissen/wordrow/internal/mappings/errors"
)

// Regular expression of a MarkDown table row.
var tableDividerExpr = regexp.MustCompile(`^\s*\|(\s*-+\s*\|){2,}\s*$`)

// Check whether or not a line in a MarkDown file is part of a table.
func isTableRow(row string) bool {
	row = stringsx.TrimSpace(row)
	return stringsx.HasPrefix(row, "|") && stringsx.HasSuffix(row, "|")
}

// Parse a row of a MarkDown table into it's column values.
//
// The error will be set if the row has an unexpected format, for example an
// incorrect number of columns.
func parseTableRow(row string) ([]string, error) {
	rowValuesCount := 4

	rowValues := stringsx.Split(row, "|")
	if len(rowValues) < rowValuesCount {
		return nil, errors.NewIncorrectFormat(row)
	}

	rowValues = rowValues[1 : len(rowValues)-1]
	rowValues = stringsx.MapAll(rowValues, stringsx.TrimSpace)
	if stringsx.Any(rowValues, stringsx.IsEmpty) {
		return nil, errors.NewMissingValue(row)
	}

	return rowValues, nil
}

// Parse the header of a MarkDown table.
//
// The error will be set if the table header has an unexpected format, such as
// an incorrect number of columns or a missing divider.
func verifyTableHeader(tableLines []string) (err error) {
	headerLine := tableLines[0]
	dividerLine := tableLines[1]
	firstTableRow := tableLines[2]

	if _, e := parseTableRow(headerLine); e != nil {
		err = errors.Newf("Incorrect table header (in '%s')", headerLine)
	} else if _, e = parseTableRow(dividerLine); e != nil {
		err = errors.Newf("Missing table divider (in '%s')", dividerLine)
	} else if !tableDividerExpr.MatchString(dividerLine) {
		err = errors.Newf("Incorrect table divider (in '%s')", dividerLine)
	} else if _, e = parseTableRow(firstTableRow); e != nil {
		err = errors.Newf("Missing table body (in '%s')", firstTableRow)
	}

	return err
}

// Parse a MarkDown table body and put its values into the `mapping`.
//
// The error will be set if any table row has an incorrect format.
func parseTableBody(rows []string, mapping map[string]string) error {
	for _, row := range rows {
		if !isTableRow(row) {
			break
		}

		rowValues, err := parseTableRow(row)
		if err != nil {
			return err
		}

		last := len(rowValues) - 1
		additionalMappings := common.MapFrom(rowValues[0:last], rowValues[last])
		common.MergeMaps(mapping, additionalMappings)
	}

	return nil
}

// Parse a MarkDown table and put its values into the `mapping`.
//
// The error will be set if the table head or any table row has an incorrect
// format.
func parseTable(lines []string, mapping map[string]string) (int, error) {
	minRowCount := 3
	tableHeadOffset := 2
	sizeBefore := len(mapping)

	if len(lines) < minRowCount {
		return 0, errors.Newf("Incomplete table (starting at '%s')", lines[0])
	}

	if err := verifyTableHeader(lines); err != nil {
		return 0, err
	}

	err := parseTableBody(lines[tableHeadOffset:], mapping)
	if err != nil {
		return 0, err
	}

	return (tableHeadOffset + (len(mapping) - sizeBefore)), nil
}

// Parse a MarkDown (MD) formatted file into a map[string]string.
//
// The error will be set if any error occurred while parsing the MD file.
func Parse(rawFileData *string) (map[string]string, error) {
	mapping := make(map[string]string, 1)

	lines := stringsx.Split(*rawFileData, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if isTableRow(line) {
			tableLength, err := parseTable(lines[i:], mapping)
			if err != nil {
				return mapping, err
			}

			i += tableLength
		}
	}

	return mapping, nil
}

var pipeByte = []byte("|")

// Check whether or not a line in a MarkDown file is part of a table.
func isTableRowBytes(line []byte) bool {
	line = bytes.TrimSpace(line)
	return bytes.HasPrefix(line, pipeByte) && bytes.HasSuffix(line, pipeByte)
}

// Parse a row of a MarkDown table into it's column values.
//
// The error will be set if the row has an unexpected format, for example an
// incorrect number of columns.
func parseTableRowBytes(row []byte) ([][]byte, error) {
	rowValuesCount := 4

	rowValues := bytes.Split(row, pipeByte)
	if len(rowValues) < rowValuesCount {
		return nil, errors.NewIncorrectFormat(row)
	}

	rowValues = rowValues[1 : len(rowValues)-1]

	for i, rowValue := range rowValues {
		rowValues[i] = bytes.TrimSpace(rowValue)
		if len(rowValues[i]) == 0 {
			return nil, errors.NewMissingValue(row)
		}
	}

	return rowValues, nil
}

// Parse the header of a MarkDown table.
//
// The error will be set if the table header has an unexpected format, such as
// an incorrect number of columns or a missing divider.
func verifyTableHeaderBytes(r *bufio.Reader) (err error) {
	dividerLine, _, _ := r.ReadLine()
	if _, e := parseTableRowBytes(dividerLine); e != nil {
		err = errors.Newf("Missing table divider (in '%s')", dividerLine)
	} else if !tableDividerExpr.Match(dividerLine) {
		err = errors.Newf("Incorrect table divider (in '%s')", dividerLine)
	}

	return err
}

// Parse a MarkDown table body and put its values into the `mapping`.
//
// The error will be set if any table row has an incorrect format.
func parseTableBodyBytes(r *bufio.Reader, mapping *ByteMap) (err error) {
	row, _, err := r.ReadLine()
	if err != nil || !isTableRowBytes(row) {
		return errors.Newf("Missing table body (in '%s')", row)
	}

	for ; err == nil; row, _, err = r.ReadLine() {
		if !isTableRowBytes(row) {
			break
		}

		rowValues, err := parseTableRowBytes(row)
		if err != nil {
			return err
		}

		last := len(rowValues) - 1
		mapping.AddMany(rowValues[0:last], rowValues[last])
	}

	return nil
}

// Parse a MarkDown table and put its values into the `mapping`.
//
// The error will be set if the table head or any table row has an incorrect
// format.
func parseTableBytes(r *bufio.Reader, mapping *ByteMap) error {
	if err := verifyTableHeaderBytes(r); err != nil {
		return err
	}

	err := parseTableBodyBytes(r, mapping)
	if err != nil {
		return err
	}

	return nil
}

// ParseReader parses a MarkDown (MD) file from a reader into a ByteMap.
//
// The error will be set if any error occurred while parsing the MarkDown file.
func ParseReader(r *bufio.Reader) (mapping ByteMap, err error) {
	var line []byte
	for ; err == nil; line, _, err = r.ReadLine() {
		if isTableRowBytes(line) {
			err := parseTableBytes(r, &mapping)
			if err != nil {
				return mapping, err
			}
		}
	}

	return mapping, nil
}

// Parse a MarkDown table body and put its values into the `mapping`.
//
// The error will be set if any table row has an incorrect format.
func parseTableBodyBytesString(r *bufio.Reader, mapping map[string]string) (err error) {
	row, _, err := r.ReadLine()
	if err != nil || !isTableRowBytes(row) {
		return errors.Newf("Missing table body (in '%s')", row)
	}

	for ; err == nil; row, _, err = r.ReadLine() {
		if !isTableRowBytes(row) {
			break
		}

		rowValues, err := parseTableRowBytes(row)
		if err != nil {
			return err
		}

		last := len(rowValues) - 1
		to := string(rowValues[last])
		for _, from := range rowValues[0:last] {
			mapping[string(from)] = to
		}
	}

	return nil
}

// Parse a MarkDown table and put its values into the `mapping`.
//
// The error will be set if the table head or any table row has an incorrect
// format.
func parseTableBytesString(r *bufio.Reader, mapping map[string]string) error {
	if err := verifyTableHeaderBytes(r); err != nil {
		return err
	}

	err := parseTableBodyBytesString(r, mapping)
	if err != nil {
		return err
	}

	return nil
}

// ParseReaderString parses a MarkDown (MD) file from a reader into a
// map[string]string.
//
// The error will be set if any error occurred while parsing the MarkDown file.
func ParseReaderString(r *bufio.Reader) (mapping map[string]string, err error) {
	mapping = make(map[string]string, 1)

	var line []byte
	for ; err == nil; line, _, err = r.ReadLine() {
		if isTableRowBytes(line) {
			err := parseTableBytesString(r, mapping)
			if err != nil {
				return mapping, err
			}
		}
	}

	return mapping, nil
}
