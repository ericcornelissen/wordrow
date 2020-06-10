// Package fs is a simple utilities package that provides functions to interact
// with the file system.
package fs

import (
	"io/ioutil"
	"path/filepath"
)

// ReadFile reads a file given a path and gets it as a File instance. Relative
// file paths are automatically resolved.
//
// The function sets the error if the file couldn't be read.
func ReadFile(filePath string) (File, error) {
	filePath = ResolvePath(filePath)

	binaryFileData, err := ioutil.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return File{}, err
	}

	return File{
		Content: string(binaryFileData),
		Ext:     GetExt(filePath),
		Path:    filePath,
	}, nil
}

// ReadFiles reads files given a list of paths and get them as File instances.
// Relative paths are automatically resolved and globs are automatically
// evaluated. Note that as a result the output list may be larger than the input
// list.
//
// The function sets the error if any of the files couldn't be read.
func ReadFiles(paths []string) ([]File, error) {
	var files []File

	paths, err := ResolveGlobs(paths...)
	if err != nil {
		return files, err
	}

	for _, path := range paths {
		file, err := ReadFile(path)
		if err != nil {
			return files, err
		}

		files = append(files, file)
	}

	return files, nil
}
