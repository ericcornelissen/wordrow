package wordmaps

import (
	"reflect"
	"testing"
)

func TestGetParserForUnknownFileType(t *testing.T) {
	_, err := getParserForFormat(".bar")

	if err == nil {
		t.Fatal("The error should be set for unknown formats")
	}
}

func TestGetParserForMarkDownFile(t *testing.T) {
	check := func(t *testing.T, parseFn parseFunction, err error) {
		t.Helper()

		if err != nil {
			t.Fatalf("The error should be nil for this test (got '%s')", err)
		}

		actual := reflect.ValueOf(parseFn)
		expected := reflect.ValueOf(parseMarkDownFile)
		if actual.Pointer() != expected.Pointer() {
			t.Error("The parser function should be the MarkDown parse function")
		}
	}

	t.Run(".md", func(t *testing.T) {
		parseFn, err := getParserForFormat(".md")
		check(t, parseFn, err)
	})
	t.Run(".MD", func(t *testing.T) {
		parseFn, err := getParserForFormat(".MD")
		check(t, parseFn, err)
	})
	t.Run(".mdown", func(t *testing.T) {
		parseFn, err := getParserForFormat(".mdown")
		check(t, parseFn, err)
	})
	t.Run(".markdown", func(t *testing.T) {
		parseFn, err := getParserForFormat(".markdown")
		check(t, parseFn, err)
	})
	t.Run(".mdwn", func(t *testing.T) {
		parseFn, err := getParserForFormat(".mdwn")
		check(t, parseFn, err)
	})
	t.Run(".mkdn", func(t *testing.T) {
		parseFn, err := getParserForFormat(".mkdn")
		check(t, parseFn, err)
	})
	t.Run(".mkdn", func(t *testing.T) {
		parseFn, err := getParserForFormat(".mdtxt")
		check(t, parseFn, err)
	})
	t.Run(".mdtext", func(t *testing.T) {
		parseFn, err := getParserForFormat(".mdtext")
		check(t, parseFn, err)
	})
	t.Run("md", func(t *testing.T) {
		parseFn, err := getParserForFormat("md")
		check(t, parseFn, err)
	})
}

func TestGetParserForCSVFile(t *testing.T) {
	check := func(t *testing.T, parseFn parseFunction, err error) {
		t.Helper()

		if err != nil {
			t.Fatalf("The error should be nil for this test (got '%s')", err)
		}

		actual := reflect.ValueOf(parseFn)
		expected := reflect.ValueOf(parseCsvFile)
		if actual.Pointer() != expected.Pointer() {
			t.Error("The parser function should be the CSV parse function")
		}
	}

	t.Run(".csv", func(t *testing.T) {
		parseFn, err := getParserForFormat(".csv")
		check(t, parseFn, err)
	})
	t.Run(".CSV", func(t *testing.T) {
		parseFn, err := getParserForFormat(".CSV")
		check(t, parseFn, err)
	})
	t.Run("csv", func(t *testing.T) {
		parseFn, err := getParserForFormat("csv")
		check(t, parseFn, err)
	})
}

func TestParseFileNoParser(t *testing.T) {
	content := "Hello world!"

	_, err := parseFile(&content, ".bar")
	if err == nil {
		t.Fatal("The error should set for this test")
	}
}

func TestParseFileUpdatesWordMap(t *testing.T) {
	content := "this is definitely not a real CSV file"

	_, err := parseFile(&content, ".csv")
	if err == nil {
		t.Fatal("The error should set for this test")
	}
}

func TestParseFileParseCSV(t *testing.T) {
	content := "foo,bar"

	wm, err := parseFile(&content, ".csv")
	if err != nil {
		t.Fatalf("The error should not be set for this test (got '%s')", err)
	}

	if len(wm) == 0 {
		t.Error("The size of the WordMap should be greater than 0")
	}
}

func TestParseFileParseMarkDown(t *testing.T) {
	content := `
		| From | To  |
		| ---- | --- |
		| foo  | bar |
	`

	wm, err := parseFile(&content, ".md")
	if err != nil {
		t.Fatalf("The error should not be set for this test (got '%s')", err)
	}

	if len(wm) == 0 {
		t.Error("The size of the WordMap should be greater than 0")
	}
}
