package main

import (
	"io"

	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/replacer"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

func doReplace(file fs.File, wordmap *wordmaps.WordMap) string {
	content := file.Content
	return replacer.ReplaceAll(content, *wordmap)
}

func doWriteBack(handle io.Writer, fixed string) error {
	b := []byte(fixed)
	_, err := handle.Write(b)
	return err
}
