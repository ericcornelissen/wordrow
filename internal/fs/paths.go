package fs

import "os"
import "path"
import "path/filepath"

import "github.com/ericcornelissen/wordrow/internal/logger"


// Get the (current) working directory.
//
// The function panics if the (current) working directory could not be found.
func getCwd() string {
  cwd, err := os.Getwd()
  if err != nil {
    logger.Fatal("Current working directory could not be obtained")
    panic(1)
  }

  return cwd
}

// Resolve any number of absolute or relative paths to absolute paths only.
//
// The function panics if the (current) working directory is needed but could
// not be found.
func ResolvePaths(files ...string) []string {
  var paths []string
  for i := 0; i < len(files); i++ {
    file := files[i]
    if filepath.IsAbs(file) {
      paths = append(paths, file)
    } else {
      file = path.Join(getCwd(), file)
      paths = append(paths, file)
    }
  }

  return paths
}
