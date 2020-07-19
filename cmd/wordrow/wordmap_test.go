package main

import (
	"fmt"
	"testing"
	"testing/iotest"

	"github.com/ericcornelissen/wordrow/internal/strings"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

func TestParseMapFileArgument(t *testing.T) {
	t.Run("Just a file", func(t *testing.T) {
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
	t.Run("File with explicit format", func(t *testing.T) {
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
}

func TestProcessMapFile(t *testing.T) {
	t.Run("Read something, correct format", func(t *testing.T) {
		var wordmap wordmaps.WordMap

		content := "foo,bar"
		handle := strings.NewReader(content)
		format := "csv"

		err := processMapFile(handle, format, &wordmap)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		wordmapSize := wordmap.Size()
		if wordmapSize != 1 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}

		from := wordmap.GetFrom(0)
		if from != "foo" {
			t.Errorf("Incorrect first from value (got '%s')", from)
		}

		to := wordmap.GetTo(0)
		if to != "bar" {
			t.Errorf("Incorrect first to value (got '%s')", to)
		}
	})
	t.Run("Read something, incorrect format", func(t *testing.T) {
		var wordmap wordmaps.WordMap

		content := "foobar"
		handle := strings.NewReader(content)
		format := "csv"

		err := processMapFile(handle, format, &wordmap)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}

		wordmapSize := wordmap.Size()
		if wordmapSize != 0 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}
	})
	t.Run("Read nothing", func(t *testing.T) {
		var wordmap wordmaps.WordMap

		content := ""
		handle := strings.NewReader(content)
		format := "csv"

		err := processMapFile(handle, format, &wordmap)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		wordmapSize := wordmap.Size()
		if wordmapSize != 0 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}
	})
	t.Run("Reading error", func(t *testing.T) {
		var wordmap wordmaps.WordMap

		content := "Hello world"
		handle := iotest.TimeoutReader(strings.NewReader(content))
		format := "csv"

		err := processMapFile(handle, format, &wordmap)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}

		wordmapSize := wordmap.Size()
		if wordmapSize != 0 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}
	})
}

func TestProcessInlineMapping(t *testing.T) {
	t.Run("Correct format", func(t *testing.T) {
		var wordmap wordmaps.WordMap

		mapping := "hello,hey"
		err := processInlineMapping(mapping, &wordmap)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		wordmapSize := wordmap.Size()
		if wordmapSize != 1 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}

		from := wordmap.GetFrom(0)
		if from != "hello" {
			t.Errorf("Incorrect first from value (got '%s')", from)
		}

		to := wordmap.GetTo(0)
		if to != "hey" {
			t.Errorf("Incorrect first to value (got '%s')", to)
		}
	})
	t.Run("Incorrect format", func(t *testing.T) {
		var wordmap wordmaps.WordMap

		mapping := "foobar"
		err := processInlineMapping(mapping, &wordmap)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}

		wordmapSize := wordmap.Size()
		if wordmapSize != 0 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}
	})
}

func TestProcessInlineMappings(t *testing.T) {
	t.Run("Correct formats", func(t *testing.T) {
		var wordmap wordmaps.WordMap

		mappings := []string{"hello,hey", "world,planet"}
		err := processInlineMappings(mappings, &wordmap)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		wordmapSize := wordmap.Size()
		if wordmapSize != 2 {
			t.Fatalf("Unexpected wordmap size (got %d)", wordmapSize)
		}

		from := wordmap.GetFrom(0)
		if from != "hello" {
			t.Errorf("Incorrect first from value (got '%s')", from)
		}

		to := wordmap.GetTo(0)
		if to != "hey" {
			t.Errorf("Incorrect first to value (got '%s')", to)
		}

		from = wordmap.GetFrom(1)
		if from != "world" {
			t.Errorf("Incorrect second from value (got '%s')", from)
		}

		to = wordmap.GetTo(1)
		if to != "planet" {
			t.Errorf("Incorrect second to value (got '%s')", to)
		}
	})
	t.Run("Correct format, Incorrect format", func(t *testing.T) {
		var wordmap wordmaps.WordMap

		mappings := []string{"hello,hey", "worldplanet"}
		err := processInlineMappings(mappings, &wordmap)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}
	})
	t.Run("Incorrect format, Correct format", func(t *testing.T) {
		var wordmap wordmaps.WordMap

		mappings := []string{"hellohey", "world,planet"}
		err := processInlineMappings(mappings, &wordmap)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}
	})
}
