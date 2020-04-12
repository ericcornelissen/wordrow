package fs

import "io/ioutil"

import "github.com/ericcornelissen/wordrow/internal/logger"

// Read a file and get the contents as a string.
//
// The function sets the error if the file couldn't be read. If it is set, the
// error is already logged.
func ReadFile(filePath string) (string, error) {
  binaryFileData, err := ioutil.ReadFile(filePath)
  if err != nil {
    logger.Errorf("File not found ('%s')", filePath)
    return "", err
  } else {
    return string(binaryFileData), nil
  }
}
