package dicts

import "errors"
import "fmt"
import "strings"

import "github.com/ericcornelissen/wordrow/internal/fs"
import "github.com/ericcornelissen/wordrow/internal/logger"


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
// The error is set if the parsing failed, e.g. when the file is improperly
// formatted.
func parseFile(fileContent *string, parseFn parseFunction, wordmap *WordMap) error {
  fileMap, err := parseFn(fileContent)
  if err != nil {
    return err
  }

  wordmap.AddFrom(fileMap)
  return nil
}


// Parse a list of files (relative or absolute paths) into a single WordMap.
//
// The error is set if an error occurs when parsing any of the provided file.
func WordMapFrom(files ...string) (WordMap, error) {
  var wordmap WordMap

  paths := fs.ResolvePaths(files...)
  for _, filePath := range paths {
    fileContent, err := fs.ReadFile(filePath)
    if err != nil {
      logger.Errorf("could not find '%s'\n", filePath)
      continue
    }

    parserFn, err := getParserForFile(filePath)
    if err != nil {
      logger.Errorf("The file '%s' is of an unknown type\n", filePath)
      continue
    }

    err = parseFile(&fileContent, parserFn, &wordmap)
    if err != nil {
      return wordmap, fmt.Errorf("Error when parsing %s: %s", filePath, err)
    }
  }

  return wordmap, nil
}
