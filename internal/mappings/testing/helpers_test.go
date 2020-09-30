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
