package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"
	"testing/iotest"

	"github.com/ericcornelissen/stringsx"
)

func TestDoReplace(t *testing.T) {
	mapping := make(map[string]string, 1)
	mapping["foo"] = "bar"

	t.Run("Replace something", func(t *testing.T) {
		content := "Foo Bar"
		handle := stringsx.NewReader(content)

		fixed, err := doReplace(handle, mapping)
		if err != nil {
			t.Fatalf("Unexpected error for reader (%s)", err)
		}

		if fixed == content {
			t.Error("Content should have been changed but wasn't")
		}
	})
	t.Run("Replace nothing", func(t *testing.T) {
		content := "Bar"
		handle := stringsx.NewReader(content)

		fixed, err := doReplace(handle, mapping)
		if err != nil {
			t.Fatalf("Unexpected error for reader (%s)", err)
		}

		if fixed != content {
			t.Errorf("Content should not have been changed but was (got '%s')", fixed)
		}
	})
	t.Run("Reading error", func(t *testing.T) {
		content := "Hello world"
		handle := iotest.TimeoutReader(stringsx.NewReader(content))

		_, err := doReplace(handle, mapping)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}
	})
	t.Run("Empty reader", func(t *testing.T) {
		handle := stringsx.NewReader("")

		fixed, err := doReplace(handle, mapping)
		if err != nil {
			t.Fatalf("Unexpected error for reader (%s)", err)
		}

		if fixed != "" {
			t.Errorf("Updated content should have been empty (got '%s')", fixed)
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
		if len(content) < 2 {
			t.Fatal("Content must be at least 2 bytes to ensure the writer errors")
		}

		handle := iotest.TruncateWriter(os.Stdin, 1)

		err := doWriteBack(handle, content)
		if err == nil {
			t.Fatal("Expected an error but got none")
		}
	})
}

func TestProcessBuffer(t *testing.T) {
	from0, to0 := "hello", "hey"
	from1, to1 := "world", "planet"

	mapping := make(map[string]string, 2)
	mapping[from0] = to0
	mapping[from1] = to1

	t.Run("Replace something", func(t *testing.T) {
		content := fmt.Sprintf("%s %s", from0, from1)
		expectedWritten := fmt.Sprintf("%s %s\n", to0, to1)

		reader := stringsx.NewReader(content)
		writer := new(bytes.Buffer)
		scanner := bufio.NewScanner(reader)
		bufferedWriter := bufio.NewWriter(writer)

		err := processBuffer(scanner, bufferedWriter, mapping)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		written := writer.Bytes()
		if string(written) != expectedWritten {
			t.Errorf("Unexpected value written (got '%s')", written)
		}
	})
	t.Run("Replace nothing", func(t *testing.T) {
		content := "foobar\n"
		if stringsx.Contains(content, from0) || stringsx.Contains(content, from1) {
			t.Fatal("Content cannot contain a string that may be replaced")
		}

		reader := stringsx.NewReader(content)
		writer := new(bytes.Buffer)
		scanner := bufio.NewScanner(reader)
		bufferedWriter := bufio.NewWriter(writer)

		err := processBuffer(scanner, bufferedWriter, mapping)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		bufferedWriter.Flush()
		written := writer.Bytes()
		if string(written) != content {
			t.Errorf("Unexpected value written (got '%s')", written)
		}
	})
	t.Run("Writing error", func(t *testing.T) {
		content := "foobar"
		if len(content) < 2 {
			t.Fatal("Content must be at least 2 bytes to ensure the writer errors")
		}

		reader := stringsx.NewReader(content)
		writer := iotest.TruncateWriter(os.Stdin, 1)
		scanner := bufio.NewScanner(reader)
		bufferedWriter := bufio.NewWriterSize(writer, 109)

		err := processBuffer(scanner, bufferedWriter, mapping)
		if err == nil {
			t.Fatal("Expected an error but got none")
		}
	})
}

func TestProcessFile(t *testing.T) {
	from0, to0 := "hello", "hey"
	from1, to1 := "world", "planet"

	mapping := make(map[string]string, 2)
	mapping[from0] = to0
	mapping[from1] = to1

	t.Run("Replace something", func(t *testing.T) {
		content := fmt.Sprintf("%s %s", from0, from1)
		expectedWritten := fmt.Sprintf("%s %s", to0, to1)

		reader := stringsx.NewReader(content)
		writer := new(bytes.Buffer)
		bufferedReader := bufio.NewReader(reader)
		bufferedWriter := bufio.NewWriter(writer)
		handle := bufio.NewReadWriter(bufferedReader, bufferedWriter)

		err := processFile(handle, mapping)
		if err != nil {
			t.Fatalf("Unexpected error (%s)", err)
		}

		bufferedWriter.Flush()
		written := writer.Bytes()
		if string(written) != expectedWritten {
			t.Errorf("Unexpected value written (got '%s')", written)
		}
	})
	t.Run("Replace nothing", func(t *testing.T) {
		content := "foobar"
		if stringsx.Contains(content, from0) || stringsx.Contains(content, from1) {
			t.Fatal("Content cannot contain a string that may be replaced")
		}

		reader := stringsx.NewReader(content)
		writer := new(bytes.Buffer)
		bufferedReader := bufio.NewReader(reader)
		bufferedWriter := bufio.NewWriter(writer)
		handle := bufio.NewReadWriter(bufferedReader, bufferedWriter)

		err := processFile(handle, mapping)
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

		reader := iotest.TimeoutReader(stringsx.NewReader(content))
		writer := new(bytes.Buffer)
		bufferedReader := bufio.NewReader(reader)
		bufferedWriter := bufio.NewWriter(writer)
		handle := bufio.NewReadWriter(bufferedReader, bufferedWriter)

		err := processFile(handle, mapping)
		if err == nil {
			t.Fatal("Expected an error but got none")
		}

		bufferedWriter.Flush()
		written := writer.Bytes()
		if string(written) != "" {
			t.Errorf("Unexpected value written (got '%s')", written)
		}
	})
	t.Run("Writing error", func(t *testing.T) {
		content := "foobar"
		if len(content) < 2 {
			t.Fatal("Content must be at least 2 bytes to ensure the writer errors")
		}

		reader := stringsx.NewReader(content)
		writer := iotest.TruncateWriter(os.Stdin, 1)
		bufferedReader := bufio.NewReader(reader)
		bufferedWriter := bufio.NewWriterSize(writer, 1)
		handle := bufio.NewReadWriter(bufferedReader, bufferedWriter)

		err := processFile(handle, mapping)
		if err == nil {
			t.Fatal("Expected an error but got none")
		}
	})
}
