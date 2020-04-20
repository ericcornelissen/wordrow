package wordmap

import "strings"

import "github.com/ericcornelissen/wordrow/internal/errors"
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


// Parse a list of Files into a single WordMap.
//
// The function sets the error if an error occurs when parsing any of the
// provided file.
func WordMapFrom(files ...fs.File) (WordMap, error) {
  var wordmap WordMap

  for _, file := range files {
    parserFn, err := getParserForFile(file.Path)
    if err != nil {
      logger.Errorf("The file '%s' is of an unknown type\n", file.Path)
      continue
    }

    err = parseFile(&file.Content, parserFn, &wordmap)
    if err != nil {
      return wordmap, errors.Newf("Error when parsing %s: %s", file.Path, err)
    }
  }

  return wordmap, nil
}
