package wordmaps

import "reflect"
import "testing"

import "github.com/ericcornelissen/wordrow/internal/fs"

func TestGetParserForUnknownFileType(t *testing.T) {
	file := fs.File{
		Ext:  ".bar",
		Path: "foo.bar",
	}

	_, err := getParserForFile(&file)

	if err == nil {
		t.Error("The error should be set for unknown file types")
	}
}

func TestGetParserForMarkDownFile(t *testing.T) {
	file := fs.File{
		Ext:  ".md",
		Path: "foo.md",
	}

	parserFn, err := getParserForFile(&file)

	if err != nil {
		t.Fatalf("The error should be nil for this test (Error: %s)", err)
	}

	actual, expected := reflect.ValueOf(parserFn), reflect.ValueOf(parseMarkDownFile)
	if actual.Pointer() != expected.Pointer() {
		t.Error("The parser function should be the MarkDown parse function")
	}
}

func TestGetParserForCSVFile(t *testing.T) {
	file := fs.File{
		Ext:  ".csv",
		Path: "foo.csv",
	}

	parserFn, err := getParserForFile(&file)

	if err != nil {
		t.Fatalf("The error should be nil for this test (Error: %s)", err)
	}

	actual, expected := reflect.ValueOf(parserFn), reflect.ValueOf(parseCsvFile)
	if actual.Pointer() != expected.Pointer() {
		t.Error("The parser function should be the CSV parse function")
	}
}

func TestParseFileNoParser(t *testing.T) {
	var wm WordMap
	file := fs.File{
		Content: "",
		Ext:     ".bar",
		Path:    "foo.bar",
	}

	err := parseFile(&file, &wm)

	if err == nil {
		t.Error("The error should set for this test")
	}
}

func TestParseFileUpdatesWordMap(t *testing.T) {
	var wm WordMap
	file := fs.File{
		Content: "this is definitely not a real CSV file",
		Ext:     ".csv",
		Path:    "foo.csv",
	}

	err := parseFile(&file, &wm)

	if err == nil {
		t.Error("The error should set for this test")
	}
}

func TestParseFileParseCSV(t *testing.T) {
	var wm WordMap
	file := fs.File{
		Content: "foo,bar",
		Ext:     ".csv",
		Path:    "foo.csv",
	}

	err := parseFile(&file, &wm)

	if err != nil {
		t.Fatalf("The error should not be set for this test")
	}

	if wm.Size() == 0 {
		t.Error("The size of the wm should be greater than 0")
	}
}

func TestParseFileParseMarkDown(t *testing.T) {
	var wm WordMap
	file := fs.File{
		Content: `
      | From | To  |
      | ---- | --- |
      | foo  | bar |
    `,
		Ext:  ".md",
		Path: "foo.md",
	}

	err := parseFile(&file, &wm)

	if err != nil {
		t.Fatalf("The error should not be set for this test")
	}

	if wm.Size() == 0 {
		t.Error("The size of the wm should be greater than 0")
	}
}
