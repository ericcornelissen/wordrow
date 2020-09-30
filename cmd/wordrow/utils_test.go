package main

import "testing"

func TestInvert(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		mapping := make(map[string]string)

		result := invert(mapping)
		if len(result) != 0 {
			t.Errorf("Unexpected size of inverted map (got %d)", len(result))
		}
	})
	t.Run("inverts the map", func(t *testing.T) {
		from0, to0 := "foo", "bar"
		from1, to1 := "hello", "world"

		mapping := make(map[string]string, 2)
		mapping[from0] = to0
		mapping[from1] = to1

		result := invert(mapping)
		if len(result) != len(mapping) {
			t.Errorf("Unexpected size of inverted map (got %d)", len(result))
		}

		if result[to0] != from0 {
			t.Errorf("Unexpected value for first key (got '%s')", result[to0])
		}

		if result[to1] != from1 {
			t.Errorf("Unexpected value for second key (got '%s')", result[to1])
		}
	})
	t.Run("works with mirrored mapping", func(t *testing.T) {
		from0, to0 := "foo", "bar"
		from1, to1 := "bar", "foo"

		mapping := make(map[string]string, 2)
		mapping[from0] = to0
		mapping[from1] = to1

		result := invert(mapping)
		if len(result) != len(mapping) {
			t.Errorf("Unexpected size of inverted map (got %d)", len(result))
		}

		if result[to0] != from0 {
			t.Errorf("Unexpected value for first key (got '%s')", result[to0])
		}

		if result[to1] != from1 {
			t.Errorf("Unexpected value for second key (got '%s')", result[to1])
		}
	})
}
