package common

import (
	"bytes"
	"testing"
)

func TestAddValuesToMap(t *testing.T) {
	t.Run("two values", func(t *testing.T) {
		from, to := "baz", "bar"

		mapping := make(map[string]string, 1)
		values := [][]byte{
			[]byte(from),
			[]byte(to),
		}

		AddValuesToMap(mapping, values)

		if mapping[from] != to {
			t.Errorf("Unexpected value for %s (got '%s')", from, to)
		}
	})
	t.Run("many values", func(t *testing.T) {
		from1, from2, to := "hello", "hey", "howdy"

		mapping := make(map[string]string, 1)
		values := [][]byte{
			[]byte(from1),
			[]byte(from2),
			[]byte(to),
		}

		AddValuesToMap(mapping, values)

		if mapping[from1] != to {
			t.Errorf("Unexpected value for %s (got '%s')", from1, to)
		}

		if mapping[from2] != to {
			t.Errorf("Unexpected value for %s (got '%s')", from2, to)
		}
	})
}

func TestTrimValues(t *testing.T) {
	t.Run("no values", func(t *testing.T) {
		inp := [][]byte{}

		out, err := TrimValues(inp)
		if err != nil {
			t.Errorf("Unexpected error (got '%s')", err)
		}

		if len(out) != 0 {
			t.Errorf("Expected output to be empty (got %d)", len(out))
		}
	})
	t.Run("many non-empty values", func(t *testing.T) {
		inp0, inp1, inp2 := "hello ", " hey", " howdy "

		inp := [][]byte{
			[]byte(inp0),
			[]byte(inp1),
			[]byte(inp2),
		}

		out, err := TrimValues(inp)
		if err != nil {
			t.Errorf("Unexpected error (got '%s')", err)
		}

		if !bytes.Equal(out[0], bytes.TrimSpace([]byte(inp0))) {
			t.Errorf("Unexpected output value (got '%s')", out[0])
		}

		if !bytes.Equal(out[1], bytes.TrimSpace([]byte(inp1))) {
			t.Errorf("Unexpected output value (got '%s')", out[1])
		}

		if !bytes.Equal(out[2], bytes.TrimSpace([]byte(inp2))) {
			t.Errorf("Unexpected output value (got '%s')", out[2])
		}
	})
	t.Run("many values, one is empty", func(t *testing.T) {
		inp0, inp1, inp2 := "hello ", "  ", " howdy "

		inp := [][]byte{
			[]byte(inp0),
			[]byte(inp1),
			[]byte(inp2),
		}

		_, err := TrimValues(inp)
		if err == nil {
			t.Error("Expected an error but got none")
		}
	})
}
