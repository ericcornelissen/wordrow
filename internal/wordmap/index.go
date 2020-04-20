package wordmap

import "strings"

import "github.com/ericcornelissen/wordrow/internal/errors"


// A parse function is a function that takes the contents of a file as a string
// and outputs a WordMap. If the file is not formatted correctly the function
// may output an error.
type parseFunction func(fileContent *string) (WordMap, error)


// Get the parseFunction for a given file(name).
func getParserForFile(filePath string) (parseFunction, error) {
  if strings.HasSuffix(filePath, ".md") {
    return parseMarkDownFile, nil
  } else if strings.HasSuffix(filePath, ".csv") {
    return parseCsvFile, nil
  }

  return nil, errors.New("Unknown file type")
}

// Parse a single file into a WordMap.
//
// The function sets the error if the parsing failed, e.g. when the file is
// improperly formatted.
func parseFile(fileContent *string, parseFn parseFunction, wordmap *WordMap) error {
  fileMap, err := parseFn(fileContent)
  if err != nil {
    return err
  }

  wordmap.AddFrom(fileMap)
  return nil
}
