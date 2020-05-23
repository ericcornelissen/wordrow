package wordmaps

import "reflect"
import "testing"

func TestGetParserForUnknownFileType(t *testing.T) {
	_, err := getParserForFormat(".bar")

	if err == nil {
		t.Error("The error should be set for unknown file types")
	}
}

func TestGetParserForMarkDownFile(t *testing.T) {
	parserFn, err := getParserForFormat(".md")

	if err != nil {
		t.Fatalf("The error should be nil for this test (Error: %s)", err)
	}

	actual, expected := reflect.ValueOf(parserFn), reflect.ValueOf(parseMarkDownFile)
	if actual.Pointer() != expected.Pointer() {
		t.Error("The parser function should be the MarkDown parse function")
	}
}

func TestGetParserForCSVFile(t *testing.T) {
	parserFn, err := getParserForFormat(".csv")

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

	content := ""
	err := parseFile(&content, ".bar", &wm)

	if err == nil {
		t.Error("The error should set for this test")
	}
}

func TestParseFileUpdatesWordMap(t *testing.T) {
	var wm WordMap

	content := "this is definitely not a real CSV file"
	err := parseFile(&content, ".csv", &wm)

	if err == nil {
		t.Error("The error should set for this test")
	}
}

func TestParseFileParseCSV(t *testing.T) {
	var wm WordMap

	content := "foo,bar"
	err := parseFile(&content, ".csv", &wm)

	if err != nil {
		t.Fatalf("The error should not be set for this test")
	}

	if wm.Size() == 0 {
		t.Error("The size of the wm should be greater than 0")
	}
}

func TestParseFileParseMarkDown(t *testing.T) {
	var wm WordMap

	content := `
		| From | To  |
		| ---- | --- |
		| foo  | bar |
	`
	err := parseFile(&content, ".md", &wm)

	if err != nil {
		t.Fatalf("The error should not be set for this test")
	}

	if wm.Size() == 0 {
		t.Error("The size of the wm should be greater than 0")
	}
}
