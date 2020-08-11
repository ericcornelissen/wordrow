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
