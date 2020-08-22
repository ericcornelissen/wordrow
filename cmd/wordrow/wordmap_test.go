package main

import (
	"fmt"
	"testing"
	"testing/iotest"

	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
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
		wordmap := make(wordmaps.StringMap, 1)

		format := "csv"
		expectedFrom, expectedTo := "foo", "bar"
		content := fmt.Sprintf("%s,%s", expectedFrom, expectedTo)
		handle := stringsx.NewReader(content)

		err := processMapFile(handle, format, &wordmap)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		wordmapSize := len(wordmap)
		if wordmapSize != 1 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}

		actualTo, ok := wordmap[expectedFrom]
		if !ok {
			t.Error("From value missing from wordmap")
		}

		if actualTo != expectedTo {
			t.Errorf("Incorrect first to value (got '%s')", actualTo)
		}
	})
	t.Run("Read something, incorrect format", func(t *testing.T) {
		wordmap := make(wordmaps.StringMap, 1)

		format := "csv"
		content := "foobar"
		handle := stringsx.NewReader(content)

		err := processMapFile(handle, format, &wordmap)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}

		wordmapSize := len(wordmap)
		if wordmapSize != 0 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}
	})
	t.Run("Read nothing", func(t *testing.T) {
		wordmap := make(wordmaps.StringMap, 1)

		format := "csv"
		content := ""
		handle := stringsx.NewReader(content)

		err := processMapFile(handle, format, &wordmap)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		wordmapSize := len(wordmap)
		if wordmapSize != 0 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}
	})
	t.Run("Reading error", func(t *testing.T) {
		wordmap := make(wordmaps.StringMap, 1)

		format := "csv"
		content := "foo,bar"
		handle := iotest.TimeoutReader(stringsx.NewReader(content))

		err := processMapFile(handle, format, &wordmap)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}

		wordmapSize := len(wordmap)
		if wordmapSize != 0 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}
	})
}

func TestProcessInlineMapping(t *testing.T) {
	t.Run("Correct format", func(t *testing.T) {
		wordmap := make(wordmaps.StringMap, 1)

		expectedFrom, expectedTo := "hello", "hey"
		mapping := fmt.Sprintf("%s,%s", expectedFrom, expectedTo)

		err := processInlineMapping(mapping, &wordmap)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		wordmapSize := len(wordmap)
		if wordmapSize != 1 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}

		actualTo, ok := wordmap[expectedFrom]
		if !ok {
			t.Error("From value missing from wordmap")
		}

		if actualTo != expectedTo {
			t.Errorf("Incorrect first to value (got '%s')", actualTo)
		}
	})
	t.Run("Incorrect format", func(t *testing.T) {
		wordmap := make(wordmaps.StringMap, 1)
		mapping := "foobar"

		err := processInlineMapping(mapping, &wordmap)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}

		wordmapSize := len(wordmap)
		if wordmapSize != 0 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}
	})
	t.Run("Empty string", func(t *testing.T) {
		wordmap := make(wordmaps.StringMap, 1)

		if err := processInlineMapping("foo,", &wordmap); err == nil {
			t.Errorf("Expected no error but got one (%s)", err)
		}

		if err := processInlineMapping(",bar", &wordmap); err == nil {
			t.Errorf("Expected no error but got one (%s)", err)
		}

		wordmapSize := len(wordmap)
		if wordmapSize != 0 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}
	})
}

func TestProcessInlineMappings(t *testing.T) {
	t.Run("Correct formats", func(t *testing.T) {
		wordmap := make(wordmaps.StringMap, 1)

		expectedFrom0, expectedTo0 := "hello", "hey"
		expectedFrom1, expectedTo1 := "world", "planet"
		mappings := []string{
			fmt.Sprintf("%s,%s", expectedFrom0, expectedTo0),
			fmt.Sprintf("%s,%s", expectedFrom1, expectedTo1),
		}

		err := processInlineMappings(mappings, &wordmap)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		wordmapSize := len(wordmap)
		if wordmapSize != 2 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}

		actualTo0, ok := wordmap[expectedFrom0]
		if !ok {
			t.Error("From first value missing from wordmap")
		}

		if actualTo0 != expectedTo0 {
			t.Errorf("Incorrect first to value (got '%s')", actualTo0)
		}

		actualTo1, ok := wordmap[expectedFrom1]
		if !ok {
			t.Error("From second value missing from wordmap")
		}

		if actualTo1 != expectedTo1 {
			t.Errorf("Incorrect second to value (got '%s')", actualTo1)
		}
	})
	t.Run("Correct format, Incorrect format", func(t *testing.T) {
		wordmap := make(wordmaps.StringMap, 1)
		mappings := []string{"hello,hey", "worldplanet"}

		err := processInlineMappings(mappings, &wordmap)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}
	})
	t.Run("Incorrect format, Correct format", func(t *testing.T) {
		wordmap := make(wordmaps.StringMap, 1)
		mappings := []string{"hellohey", "world,planet"}

		err := processInlineMappings(mappings, &wordmap)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}
	})
}
