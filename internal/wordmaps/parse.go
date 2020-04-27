package wordmaps

import "github.com/ericcornelissen/wordrow/internal/fs"
import "github.com/ericcornelissen/wordrow/internal/errors"


// A parse function is a function that takes the contents of a file as a string
// and outputs a WordMap. If the file is not formatted correctly the function
// may output an error.
type parseFunction func(fileContent *string) (WordMap, error)


// Get the parseFunction for a given File.
func getParserForFile(file *fs.File) (parseFunction, error) {
  if file.Ext == ".md" {
    return parseMarkDownFile, nil
  } else if file.Ext == ".csv" {
    return parseCsvFile, nil
  }

  return nil, errors.New("Unknown file type")
}

// Parse a single file into a WordMap.
//
// The function sets the error if the parsing failed, e.g. when the file is
// improperly formatted.
func parseFile(file *fs.File, wm *WordMap) error {
  parseFn, err := getParserForFile(file)
  if err != nil {
    return errors.Newf("Unknown file type of %s", file.Path)
  }

  fileMap, err := parseFn(&file.Content)
  if err != nil {
    return err
  }

  wm.AddFrom(fileMap)
  return nil
}
