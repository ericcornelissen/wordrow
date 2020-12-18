package main

import (
	"fmt"
	"testing"

	"github.com/ericcornelissen/stringsx"
)

func TestParseMapFileArgument(t *testing.T) {
	t.Run("File with extension, no explicit format", func(t *testing.T) {
		extension := ".txt"
		input := fmt.Sprintf("/foo/bar/test%s", extension)

		filePath, format := parseMapFileArgument(input)
		if filePath != input {
			t.Errorf("Unexpected filepath (got '%s')", filePath)
		}

		if format != extension {
			t.Errorf("Unexpected format (got '%s')", format)
		}
	})
	t.Run("File with extension, with explicit format", func(t *testing.T) {
		extension := ".txt"
		explicitFormat := "csv"
		inputPath := fmt.Sprintf("/hello/world%s", extension)
		input := fmt.Sprintf("%s:%s", inputPath, explicitFormat)

		filePath, format := parseMapFileArgument(input)
		if filePath != inputPath {
			t.Errorf("Unexpected filepath (got '%s')", filePath)
		}

		if format != explicitFormat {
			t.Errorf("Unexpected format (got '%s')", format)
		}
	})
	t.Run("File without extension, no explicit format", func(t *testing.T) {
		input := "/path/to/file/without/extension"

		filePath, format := parseMapFileArgument(input)
		if filePath != input {
			t.Errorf("Unexpected filepath (got '%s')", filePath)
		}

		if format != "" {
			t.Errorf("Unexpected format (got '%s')", format)
		}
	})
	t.Run("File without extension, with explicit format", func(t *testing.T) {
		explicitFormat := "csv"
		inputPath := "/path/to/file/without/extension"
		input := fmt.Sprintf("%s:%s", inputPath, explicitFormat)

		filePath, format := parseMapFileArgument(input)
		if filePath != inputPath {
			t.Errorf("Unexpected filepath (got '%s')", filePath)
		}

		if format != explicitFormat {
			t.Errorf("Unexpected format (got '%s')", format)
		}
	})
}

func TestProcessMapFile(t *testing.T) {
	t.Run("Read something, correct format", func(t *testing.T) {
		format := "csv"
		expectedFrom, expectedTo := "foo", "bar"
		content := fmt.Sprintf("%s,%s", expectedFrom, expectedTo)
		handle := stringsx.NewReader(content)

		mapping, err := processMapFile(handle, format)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		mappingSize := len(mapping)
		if mappingSize != 1 {
			t.Fatalf("Unexpected mapping size (got %d)", mappingSize)
		}

		actualTo, ok := mapping[expectedFrom]
		if !ok {
			t.Error("From value missing from mapping")
		}

		if actualTo != expectedTo {
			t.Errorf("Incorrect first to value (got '%s')", actualTo)
		}
	})
	t.Run("Read something, incorrect format", func(t *testing.T) {
		format := "csv"
		content := "foobar"
		handle := stringsx.NewReader(content)

		mapping, err := processMapFile(handle, format)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}

		mappingSize := len(mapping)
		if mappingSize != 0 {
			t.Fatalf("Unexpected mapping size (got %d)", mappingSize)
		}
	})
	t.Run("Read nothing", func(t *testing.T) {
		format := "csv"
		content := ""
		handle := stringsx.NewReader(content)

		mapping, err := processMapFile(handle, format)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		mappingSize := len(mapping)
		if mappingSize != 0 {
			t.Fatalf("Unexpected mapping size (got %d)", mappingSize)
		}
	})
}

func TestProcessInlineMapping(t *testing.T) {
	t.Run("Correct format", func(t *testing.T) {
		mapping := make(map[string]string, 1)

		expectedFrom, expectedTo := "hello", "hey"
		value := fmt.Sprintf("%s,%s", expectedFrom, expectedTo)

		err := processInlineMapping(value, mapping)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		mappingSize := len(mapping)
		if mappingSize != 1 {
			t.Fatalf("Unexpected mapping size (got %d)", mappingSize)
		}

		actualTo, ok := mapping[expectedFrom]
		if !ok {
			t.Error("From value missing from mapping")
		}

		if actualTo != expectedTo {
			t.Errorf("Incorrect first to value (got '%s')", actualTo)
		}
	})
	t.Run("Incorrect format", func(t *testing.T) {
		mapping := make(map[string]string, 1)
		value := "foobar"

		err := processInlineMapping(value, mapping)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}

		mappingSize := len(mapping)
		if mappingSize != 0 {
			t.Fatalf("Unexpected mapping size (got %d)", mappingSize)
		}
	})
	t.Run("Empty string", func(t *testing.T) {
		mapping := make(map[string]string, 1)

		if err := processInlineMapping("foo,", mapping); err == nil {
			t.Errorf("Expected no error but got one (%s)", err)
		}

		if err := processInlineMapping(",bar", mapping); err == nil {
			t.Errorf("Expected no error but got one (%s)", err)
		}

		mappingSize := len(mapping)
		if mappingSize != 0 {
			t.Fatalf("Unexpected mapping size (got %d)", mappingSize)
		}
	})
}
