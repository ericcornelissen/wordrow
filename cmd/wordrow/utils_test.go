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

func TestMerge(t *testing.T) {
	t.Run("merge disjoint maps", func(t *testing.T) {
		target := make(map[string]string, 1)
		other := make(map[string]string, 1)
		target["foo"] = "bar"
		other["hello"] = "world"

		merge(target, other)
		if len(target) != 2 {
			t.Errorf("Unexpected size of target map (got %d)", len(target))
		}

		if target["foo"] != "bar" {
			t.Errorf("Unexpected value for key 'foo' (got '%s')", target["foo"])
		}

		if target["hello"] != "world" {
			t.Errorf("Unexpected value for key 'hello' (got '%s')", target["hello"])
		}
	})
	t.Run("other overrides in target", func(t *testing.T) {
		target := make(map[string]string, 1)
		other := make(map[string]string, 1)
		target["foo"] = "bar"
		other["foo"] = "baz"

		merge(target, other)
		if len(target) != 1 {
			t.Errorf("Unexpected size of target map (got %d)", len(target))
		}

		if target["foo"] != "baz" {
			t.Errorf("Unexpected value for key 'foo' (got '%s')", target["foo"])
		}
	})
	t.Run("other is not changed", func(t *testing.T) {
		target := make(map[string]string, 1)
		other := make(map[string]string, 1)
		target["foo"] = "bar"
		other["hello"] = "world"

		merge(target, other)
		if len(other) != 1 {
			t.Errorf("Unexpected size of other map (got %d)", len(other))
		}

		if other["hello"] != "world" {
			t.Errorf("Unexpected value for key 'hello' (got '%s')", other["hello"])
		}
	})
	t.Run("target is empty", func(t *testing.T) {
		target := make(map[string]string)
		other := make(map[string]string, 1)
		other["hello"] = "world"

		merge(target, other)
		if len(target) != 1 {
			t.Errorf("Unexpected size of target map (got %d)", len(target))
		}

		if target["hello"] != "world" {
			t.Errorf("Unexpected value for key 'hello' (got '%s')", target["hello"])
		}
	})
	t.Run("other is empty", func(t *testing.T) {
		target := make(map[string]string, 1)
		other := make(map[string]string)
		target["foo"] = "bar"

		merge(target, other)
		if len(target) != 1 {
			t.Errorf("Unexpected size of target map (got %d)", len(target))
		}

		if target["foo"] != "bar" {
			t.Errorf("Unexpected value for key 'foo' (got '%s')", target["foo"])
		}
	})
}
