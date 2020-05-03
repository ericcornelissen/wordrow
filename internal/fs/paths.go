package fs

import "os"
import "path"
import "path/filepath"
import "regexp"

import "github.com/ericcornelissen/wordrow/internal/errors"
import "github.com/ericcornelissen/wordrow/internal/logger"

import "github.com/yargevad/filepathx"

// Regular expression for glob strings.
var globExpr = regexp.MustCompile(`[\*\?\[\]]`)

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

// Get the extension of a file give the file path.
func getExt(path string) string {
	return filepath.Ext(path)
}

// ResolveGlobs resolves any number of globs or file paths into distinct file paths.
//
// The function sets the error if at least one malformed pattern is found. Only
// the last malformed pattern is reported. The list of paths will contain all
// paths for valid not-malformed patterns.
func ResolveGlobs(patterns ...string) (paths []string, rerr error) {
	for _, pattern := range patterns {
		if !globExpr.MatchString(pattern) {
			paths = append(paths, pattern)
			continue
		}

		matches, err := filepathx.Glob(pattern)
		if err != nil {
			rerr = errors.Newf("Malformed pattern (%s)", pattern)
		} else {
			paths = append(paths, matches...)
		}
	}

	return paths, rerr
}

// ResolvePath resolves a single absolute or relative path to an absolute path.
//
// The function panics if the (current) working directory is needed but could
// not be found.
func ResolvePath(inputPath string) string {
	if filepath.IsAbs(inputPath) {
		return inputPath
	}

	outputPath := path.Join(getCwd(), inputPath)
	return outputPath
}

// ResolvePaths resolves any number of absolute or relative paths to absolute
// paths only.
//
// The function panics if the (current) working directory is needed but could
// not be found.
func ResolvePaths(inputPaths ...string) []string {
	var paths []string
	for _, inputPath := range inputPaths {
		outputPath := ResolvePath(inputPath)
		paths = append(paths, outputPath)
	}

	return paths
}
