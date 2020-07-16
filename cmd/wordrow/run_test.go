package main

import (
	"testing"

	"github.com/ericcornelissen/wordrow/internal/fs"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

func TestDoReplace(t *testing.T) {
	var wordmap wordmaps.WordMap
	var files []fs.File

	wordmap.AddOne("foo", "bar")

	files = make([]fs.File, 2)
	files[0] = fs.File{
		Content: "Foo",
		Ext:     ".txt",
		Path:    "foo.txt",
	}
	files[1] = fs.File{
		Content: "bar",
		Ext:     ".txt",
		Path:    "bar.txt",
	}

	fixed := doReplace(files, &wordmap)

	if fixed[0] == files[0].Content {
		t.Errorf("File 0 should have been fixed")
	}

	if fixed[1] != files[1].Content {
		t.Errorf("File 1 should not have been fixed")
	}
}
