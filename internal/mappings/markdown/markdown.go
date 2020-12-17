package markdown

import (
	"bufio"
	"bytes"
	"regexp"

	"github.com/ericcornelissen/wordrow/internal/mappings/errors"
)

// Byte-slice representing a pipe/vertical bar ('|')
var pipeByte = []byte{'|'}

// Regular expression of a MarkDown table row.
var tableDividerExpr = regexp.MustCompile(`^\s*\|(\s*-+\s*\|){2,}\s*$`)

// Check whether or not a line in a MarkDown file is part of a table.
func isTableRow(line []byte) bool {
	line = bytes.TrimSpace(line)
	return bytes.HasPrefix(line, pipeByte) && bytes.HasSuffix(line, pipeByte)
}

// Parse a row of a MarkDown table into it's column values.
//
// The error will be set if the row has an unexpected format, for example an
// incorrect number of columns.
func parseTableRow(row []byte) ([][]byte, error) {
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
// The error will be set if the table header has an unexpected format.
func verifyTableHeader(headerLine []byte) (err error) {
	if _, e := parseTableRow(headerLine); e != nil {
		err = errors.Newf("Incorrect table header (in '%s')", headerLine)
	}

	return err
}

// Parse the divider of a MarkDown table.
//
// The error will be set if the table divider has an unexpected format.
func verifyTableDivider(reader *bufio.Reader) (err error) {
	problem := false

	dividerLine, _, err := reader.ReadLine()
	if err != nil {
		problem = true
	} else if _, e := parseTableRow(dividerLine); e != nil {
		problem = true
	} else if !tableDividerExpr.Match(dividerLine) {
		problem = true
	}

	if problem {
		err = errors.Newf("Missing table divider (in '%s')", dividerLine)
	}

	return err
}

// Parse a MarkDown table body and put its values into the `mapping`.
//
// The error will be set if any table row has an incorrect format.
func parseTableBody(
	reader *bufio.Reader,
	mapping map[string]string,
) (err error) {
	row, _, err := reader.ReadLine()
	if err != nil || !isTableRow(row) {
		return errors.Newf("Missing table body (in '%s')", row)
	}

	for ; err == nil; row, _, err = reader.ReadLine() {
		if !isTableRow(row) {
			break
		}

		rowValues, err := parseTableRow(row)
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
func parseTable(reader *bufio.Reader, mapping map[string]string) error {
	if err := verifyTableDivider(reader); err != nil {
		return err
	}

	err := parseTableBody(reader, mapping)
	if err != nil {
		return err
	}

	return nil
}

// Parse parses a MarkDown (MD) file into a map[string]string.
//
// The error will be set if any error occurred while parsing the MarkDown file.
func Parse(reader *bufio.Reader) (mapping map[string]string, err error) {
	mapping = make(map[string]string, 1)

	var line []byte
	for ; err == nil; line, _, err = reader.ReadLine() {
		if !isTableRow(line) {
			continue
		}

		if err := verifyTableHeader(line); err != nil {
			return mapping, err
		}

		err := parseTable(reader, mapping)
		if err != nil {
			return mapping, err
		}
	}

	return mapping, nil
}
