package main

import (
	"os"
	"testing"
	"testing/iotest"

	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

func TestDoReplace(t *testing.T) {
	var wordmap wordmaps.WordMap
	wordmap.AddOne("foo", "bar")

	file := fs.File{
		Content: "Foo",
		Ext:     ".txt",
		Path:    "foo.txt",
	}

	fixed := doReplace(file, &wordmap)
	if fixed == file.Content {
		t.Errorf("File 0 should have been fixed")
	}

	file = fs.File{
		Content: "Bar",
		Ext:     ".txt",
		Path:    "bar.txt",
	}

	fixed = doReplace(file, &wordmap)
	if fixed != file.Content {
		t.Errorf("File 0 should have been fixed")
	}
}

func TestDoWriteBack(t *testing.T) {
	w := iotest.NewWriteLogger("", os.Stdout)
	doWriteBack(w, "Hello world!")
}
