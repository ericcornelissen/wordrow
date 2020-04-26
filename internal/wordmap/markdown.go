package wordmap

import "regexp"
import "strings"


// Regular expression of a MarkDown table row.
var tableDividerExpr = regexp.MustCompile("^\\s*\\|\\s*-+\\s*\\|\\s*-+\\s*\\|\\s*$")


// Check whether or not a line in a MarkDown file is part of a table.
func isTableRow(row string) bool {
  row = strings.TrimSpace(row)
  return strings.HasPrefix(row, "|") && strings.HasSuffix(row, "|")
}


// Parse a row of a MarkDown table into it's column values.
//
// The error will be set if the row has an unexpected format, for example an
// incorrect number of columns.
func parseTableRow(row string) ([]string, error) {
  rowValues := strings.Split(row, "|")
  if len(rowValues) != 4 {
    return nil, &parseError{"Unexpected table row format", row}
  }

  fromValue := strings.TrimSpace(rowValues[1])
  toValue := strings.TrimSpace(rowValues[2])
  if fromValue == "" || toValue == "" {
    return nil, &parseError{"Missing value", row}
  }

  return []string{fromValue, toValue}, nil
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
    rerr = &parseError{"Incorrect table header", headerLine}
  } else if _, err = parseTableRow(dividerLine); err != nil {
    rerr = &parseError{"Missing table header divider", dividerLine}
  } else if tableDividerExpr.MatchString(dividerLine) == false {
    rerr = &parseError{"Missing table header divider", dividerLine}
  } else if _, err = parseTableRow(firstTableRow); err != nil {
    rerr = &parseError{"Missing table body", firstTableRow}
  }

  return rerr
}

// Parse a MarkDown table and put its values into a WordMap.
//
// The error will be set if the table head or any table row has an incorrect
// format.
func parseTable(tableLines []string, wm *WordMap) (int, error) {
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

    wm.AddOne(rowValues[0], rowValues[1])
  }

  return (2 + (wm.Size() - sizeBefore)), nil
}

// Parse a MarkDown, MD, formatted file into a WordMap.
//
// The error will be set if any error occured while parsing the MD file.
func parseMarkDownFile(rawFileData *string) (WordMap, error) {
  var wm WordMap

  lines := strings.Split(*rawFileData, "\n")
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
