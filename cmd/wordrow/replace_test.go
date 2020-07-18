package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"
	"testing/iotest"

	"github.com/ericcornelissen/wordrow/internal/strings"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

func TestDoReplace(t *testing.T) {
	var wordmap wordmaps.WordMap
	wordmap.AddOne("foo", "bar")

	t.Run("Replace something", func(t *testing.T) {
		content := "Foo Bar"
		handle := strings.NewReader(content)

		fixed, err := doReplace(handle, &wordmap)
		if err != nil {
			t.Fatalf("Unexpected error for reader (%s)", err)
		}

		if fixed == content {
			t.Error("Content should have been changed but wasn't")
		}
	})
	t.Run("Replace nothing", func(t *testing.T) {
		content := "Bar"
		handle := strings.NewReader(content)

		fixed, err := doReplace(handle, &wordmap)
		if err != nil {
			t.Fatalf("Unexpected error for reader (%s)", err)
		}

		if fixed != content {
			t.Errorf("Content should not have been changed but was (got '%s')", fixed)
		}
	})
	t.Run("Reading error", func(t *testing.T) {
		content := "Hello world"
		handle := iotest.TimeoutReader(strings.NewReader(content))

		_, err := doReplace(handle, &wordmap)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}
	})
}

func TestDoWriteBack(t *testing.T) {
	content := "Hello world!"

	t.Run("Write something", func(t *testing.T) {
		handle := new(bytes.Buffer)

		err := doWriteBack(handle, content)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		written := handle.Bytes()
		if string(written) != content {
			t.Errorf("Unexpected value written (got '%s')", written)
		}
	})
	t.Run("Writing error", func(t *testing.T) {
		handle := iotest.TruncateWriter(os.Stdin, 1)

		err := doWriteBack(handle, content)
		if err == nil {
			t.Fatal("Expected an error but got none")
		}
	})
}

func TestProcessFile2(t *testing.T) {
	var wordmap wordmaps.WordMap
	wordmap.AddOne("hello", "hey")
	wordmap.AddOne("world", "planet")

	t.Run("Replace something", func(t *testing.T) {
		content := "Hello world!"

		reader := strings.NewReader(content)
		writer := new(bytes.Buffer)
		bufferedReader := bufio.NewReader(reader)
		bufferedWriter := bufio.NewWriter(writer)
		handle := bufio.NewReadWriter(bufferedReader, bufferedWriter)

		err := processFile2("file.txt", handle, &wordmap)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		bufferedWriter.Flush()
		written := writer.Bytes()
		if string(written) != "Hey planet!" {
			t.Errorf("Unexpected value written (got '%s')", written)
		}
	})
	t.Run("Replace nothing", func(t *testing.T) {
		content := "foobar"

		reader := strings.NewReader(content)
		writer := new(bytes.Buffer)
		bufferedReader := bufio.NewReader(reader)
		bufferedWriter := bufio.NewWriter(writer)
		handle := bufio.NewReadWriter(bufferedReader, bufferedWriter)

		err := processFile2("file.txt", handle, &wordmap)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		bufferedWriter.Flush()
		written := writer.Bytes()
		if string(written) != content {
			t.Errorf("Unexpected value written (got '%s')", written)
		}
	})
	t.Run("Reading error", func(t *testing.T) {
		content := "foobar"

		reader := iotest.TimeoutReader(strings.NewReader(content))
		writer := new(bytes.Buffer)
		bufferedReader := bufio.NewReader(reader)
		bufferedWriter := bufio.NewWriter(writer)
		handle := bufio.NewReadWriter(bufferedReader, bufferedWriter)

		err := processFile2("file.txt", handle, &wordmap)
		if err == nil {
			t.Fatal("Expected an error but got none")
		}

		bufferedWriter.Flush()
		written := writer.Bytes()
		if !strings.IsEmpty(string(written)) {
			t.Errorf("Unexpected value written (got '%s')", written)
		}
	})
	t.Run("Writing error", func(t *testing.T) {
		content := "foobar"
		if len(content) < 2 {
			t.Fatal("Content must be at least 2 bytes to ensure the writer errors")
		}

		reader := strings.NewReader(content)
		writer := iotest.TruncateWriter(os.Stdin, 1)
		bufferedReader := bufio.NewReader(reader)
		bufferedWriter := bufio.NewWriterSize(writer, 1)
		handle := bufio.NewReadWriter(bufferedReader, bufferedWriter)

		err := processFile2("file.txt", handle, &wordmap)
		if err == nil {
			t.Fatal("Expected an error but got none")
		}
	})
}
