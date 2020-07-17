package main

import (
	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/logger"
	"github.com/ericcornelissen/wordrow/internal/replacer"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

func doReplace(
	file fs.File,
	wordmap *wordmaps.WordMap,
) string {
	logger.Debugf("Replacing words in '%s'", file.Path)
	return replacer.ReplaceAll(file.Content, *wordmap)
}

func doWriteBack(
	file fs.File,
	fixed string,
) {
	logger.Debugf("Writing updated contents to '%s'", file.Path)
	fs.WriteFile(file.Path, fixed)
}

func processFiles(
	files []fs.File,
	wordmap *wordmaps.WordMap,
) {
	for _, file := range files {
		logger.Debugf("Processing '%s'", file.Path)
		fixed := doReplace(file, wordmap)
		doWriteBack(file, fixed)
	}
}
