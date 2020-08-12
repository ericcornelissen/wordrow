package wordmaps

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
	rowValues := stringsx.Split(row, "|")
	if len(rowValues) < 4 {
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
	} else if tableDividerExpr.MatchString(dividerLine) == false {
		rerr = errors.Newf("Incorrect table divider (in '%s')", dividerLine)
	} else if _, err = parseTableRow(firstTableRow); err != nil {
		rerr = errors.Newf("Missing table body (in '%s')", firstTableRow)
	}

	return rerr
}

// Parse a MarkDown table and put its values into a WordMap.
//
// The error will be set if the table head or any table row has an incorrect
// format.
func parseTable(tableLines []string, wm *WordMap) (int, error) {
	if len(tableLines) < 3 {
		return 0, errors.Newf("Incomplete table (starting at '%s')", tableLines[0])
	}

	if err := parseTableHeader(tableLines); err != nil {
		return 0, err
	}

	sizeBefore := wm.Size()
	for i := 2; i < len(tableLines); i++ {
		row := tableLines[i]
		if !isTableRow(row) {
			break // Table ended
		}

		rowValues, err := parseTableRow(row)
		if err != nil {
			return 0, err
		}

		last := len(rowValues) - 1
		wm.AddMany(rowValues[0:last], rowValues[last])
	}

	return (2 + (wm.Size() - sizeBefore)), nil
}

// Parse a MarkDown (MD) formatted file into a WordMap.
//
// The error will be set if any error occurred while parsing the MD file.
func parseMarkDownFile(rawFileData *string) (WordMap, error) {
	var wm WordMap

	lines := stringsx.Split(*rawFileData, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if isTableRow(line) {
			tableLength, err := parseTable(lines[i:], &wm)
			if err != nil {
				return wm, err
			}

			i = i + tableLength
		}
	}

	return wm, nil
}
