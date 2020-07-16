package main

import (
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
	"github.com/ericcornelissen/wordrow/internal/replacer"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

func doReplace(
	files []fs.File,
	wordmap *wordmaps.WordMap,
) []string {
	result := make([]string, len(files))
	for i, file := range files {
		logger.Debugf("Processing '%s' as input file", file.Path)
		result[i] = replacer.ReplaceAll(file.Content, *wordmap)
	}

	return result
}
