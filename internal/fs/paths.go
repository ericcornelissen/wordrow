package fs

import (
	"path/filepath"
	"regexp"

	"github.com/ericcornelissen/wordrow/internal/errors"
	"github.com/yargevad/filepathx"
)

// Regular expression for glob strings.
var globExpr = regexp.MustCompile(`[\*\?\[\]]`)

// GetExt returns the extension of a given file path.
func GetExt(path string) string {
	return filepath.Ext(path)
}

// ResolveGlobs resolves any number of globs or file paths into distinct file
// paths. The function returns an error for every invalid pattern.
func ResolveGlobs(patterns ...string) (paths []string, errs []error) {
	for _, pattern := range patterns {
		if !globExpr.MatchString(pattern) {
			paths = append(paths, pattern)
			continue
		}

		matches, err := filepathx.Glob(pattern)
		if err != nil {
			errs = append(errs, errors.Newf("Malformed pattern (%s)", pattern))
		} else {
			paths = append(paths, matches...)
		}
	}

	return paths, errs
}
