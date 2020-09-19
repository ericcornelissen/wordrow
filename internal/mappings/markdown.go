package mappings

import (
	"regexp"

	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/errors"
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
		return nil, errors.Newf(incorrectFormat, row)
	}

	rowValues = rowValues[1 : len(rowValues)-1]
	rowValues = stringsx.MapAll(rowValues, stringsx.TrimSpace)
	if stringsx.Any(rowValues, stringsx.IsEmpty) {
		return nil, errors.Newf(missingValue, row)
	}

	return rowValues, nil
}

// Parse the header of a MarkDown table.
//
// The error will be set if the table header has an unexpected format, such as
// an incorrect number of columns or a missing divider.
func parseTableHeader(tableLines []string) (rerr error) {
	headerLine := tableLines[0]
	dividerLine := tableLines[1]
	firstTableRow := tableLines[2]

	if _, err := parseTableRow(headerLine); err != nil {
		rerr = errors.Newf("Incorrect table header (in '%s')", headerLine)
	} else if _, err = parseTableRow(dividerLine); err != nil {
		rerr = errors.Newf("Missing table divider (in '%s')", dividerLine)
	} else if !tableDividerExpr.MatchString(dividerLine) {
		rerr = errors.Newf("Incorrect table divider (in '%s')", dividerLine)
	} else if _, err = parseTableRow(firstTableRow); err != nil {
		rerr = errors.Newf("Missing table body (in '%s')", firstTableRow)
	}

	return rerr
}

// Parse a MarkDown table and put its values into the `mapping`.
//
// The error will be set if the table head or any table row has an incorrect
// format.
func parseTable(tableLines []string, mapping map[string]string) (int, error) {
	tableHeadOffset := 2
	minRowCount := 3

	if len(tableLines) < minRowCount {
		return 0, errors.Newf("Incomplete table (starting at '%s')", tableLines[0])
	}

	if err := parseTableHeader(tableLines); err != nil {
		return 0, err
	}

	sizeBefore := len(mapping)
	for i := tableHeadOffset; i < len(tableLines); i++ {
		row := tableLines[i]
		if !isTableRow(row) {
			break // Table ended
		}

		rowValues, err := parseTableRow(row)
		if err != nil {
			return 0, err
		}

		last := len(rowValues) - 1
		addToMapping(mapping, rowValues[0:last], rowValues[last])
	}

	return (tableHeadOffset + (len(mapping) - sizeBefore)), nil
}

// Parse a MarkDown (MD) formatted file into a map[string]string.
//
// The error will be set if any error occurred while parsing the MD file.
func parseMarkDownFile(rawFileData *string) (map[string]string, error) {
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
