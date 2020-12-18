package testing

import "testing"

func TestCheckMapping(t *testing.T) {
	t.Run("empty mapping and no expectations", func(t *testing.T) {
		mock := newMockT()

		mapping := make(map[string]string)
		expected := make([][]string, 0)
		CheckMapping(&mock, mapping, expected)

		if mock.helperCallCount != 1 {
			t.Error("t.Helper was not called")
		}
	})
	t.Run("succeeds if mapping matches expected", func(t *testing.T) {
		mock := newMockT()

		mapping := make(map[string]string)
		mapping["foo"] = "bar"

		expected := make([][]string, 1)
		expected[0] = []string{"foo", "bar"}

		CheckMapping(&mock, mapping, expected)

		if mock.helperCallCount != 1 {
			t.Error("t.Helper was not called")
		}

		if mock.errorCallCount != 0 {
			t.Errorf("t.Error should not be called (got %d calls)", mock.errorCallCount)
		}
	})
	t.Run("errors if values don't match", func(t *testing.T) {
		mock := newMockT()

		mapping := make(map[string]string)
		mapping["foo"] = "bar"

		expected := make([][]string, 1)
		expected[0] = []string{"foo", "baz"}

		CheckMapping(&mock, mapping, expected)

		if mock.helperCallCount != 1 {
			t.Error("t.Helper was not called")
		}

		if mock.errorCallCount != 1 {
			t.Errorf("t.Error should have been called once (got %d calls)", mock.errorCallCount)
		}
	})
	t.Run("errors if keys don't match", func(t *testing.T) {
		mock := newMockT()

		mapping := make(map[string]string)
		mapping["foo"] = "bar"

		expected := make([][]string, 1)
		expected[0] = []string{"hello", "world"}

		CheckMapping(&mock, mapping, expected)

		if mock.helperCallCount != 1 {
			t.Error("t.Helper was not called")
		}

		if mock.errorCallCount != 1 {
			t.Errorf("t.Error should have been called once (got %d calls)", mock.errorCallCount)
		}
	})
	t.Run("mapping/expected size mismatch", func(t *testing.T) {
		mock := newMockT()

		defer func() {
			if r := recover(); r != "fatal was called" {
				t.Error("t.Fatal was not called")
			}

			if mock.helperCallCount != 1 {
				t.Error("t.Helper was not called")
			}
		}()

		mapping := make(map[string]string)
		expected := make([][]string, 1)
		CheckMapping(&mock, mapping, expected)
	})
}

func TestNewTestReader(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		emptyString := ""
		reader := NewTestReader(&emptyString)
		line, _, err := reader.ReadLine()
		if len(line) != 0 {
			t.Errorf("Unexpected length of first line (got %d)", len(line))
		}

		if err == nil {
			t.Error("Expected error on second read, got none")
		}
	})
	t.Run("single line string", func(t *testing.T) {
		s := "Hello world!"
		reader := NewTestReader(&s)
		line, _, err := reader.ReadLine()
		if len(line) != len(s) {
			t.Errorf("Unexpected length of first line (got %d)", len(line))
		}

		if err != nil {
			t.Errorf("Unexpected error for first read (got '%s')", err)
		}

		_, _, err = reader.ReadLine()
		if err == nil {
			t.Error("Expected error on second read, got none")
		}
	})
	t.Run("multi-line string", func(t *testing.T) {
		s := "foo\nbar"
		reader := NewTestReader(&s)

		for i := 0; i < 2; i++ {
			_, _, err := reader.ReadLine()
			if err != nil {
				t.Errorf("Unexpected error for %dth read (got '%s')", (i + 1), err)
			}
		}

		_, _, err := reader.ReadLine()
		if err == nil {
			t.Error("Expected error on last read, got none")
		}
	})
}
