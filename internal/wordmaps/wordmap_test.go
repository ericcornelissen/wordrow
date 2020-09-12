package wordmaps

import (
	"fmt"
	"testing"
)

func TestParseFileUnknownType(t *testing.T) {
	content := "Hello world"
	format := ".bar"

	_, err := ParseFile(&content, format)
	if err == nil {
		t.Error("Expected error to be set but it was not")
	}
}

func TestParseFileKnownType(t *testing.T) {
	from, to := "foo", "bar"
	content := fmt.Sprintf("%s,%s", from, to)
	format := ".csv"

	mapping, err := ParseFile(&content, format)
	if err != nil {
		t.Fatalf("Error should not be set for this test (got '%s')", err)
	}

	expected := make([][]string, 1)
	expected[0] = []string{from, to}
	checkMapping(t, mapping, expected)
}
