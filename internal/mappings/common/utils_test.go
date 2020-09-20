package common

import "testing"

func TestMapFrom(t *testing.T) {
	t.Run("add single mapping", func(t *testing.T) {
		from, to := "foo", "bar"

		mapping := MapFrom([]string{from}, to)
		if len(mapping) != 1 {
			t.Fatalf("Unexpected size of map (got %d)", len(mapping))
		}

		if mapping[from] != to {
			t.Errorf("Unexpected value for '%s' (got '%s')", from, mapping[from])
		}
	})
	t.Run("add multiple mappings", func(t *testing.T) {
		from1, from2, to := "bar", "baz", "foo"

		mapping := MapFrom([]string{from1, from2}, to)
		if len(mapping) != 2 {
			t.Fatalf("Unexpected size of map (got %d)", len(mapping))
		}

		if mapping[from1] != to {
			t.Errorf("Unexpected value for '%s' (got '%s')", from1, mapping[from1])
		}

		if mapping[from2] != to {
			t.Errorf("Unexpected value for '%s' (got '%s')", from1, mapping[from2])
		}
	})
}

func TestMergeMaps(t *testing.T) {
	t.Run("merge disjoint maps", func(t *testing.T) {
		target := make(map[string]string, 1)
		other := make(map[string]string, 1)
		target["foo"] = "bar"
		other["hello"] = "world"

		MergeMaps(target, other)
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

		MergeMaps(target, other)
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

		MergeMaps(target, other)
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

		MergeMaps(target, other)
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

		MergeMaps(target, other)
		if len(target) != 1 {
			t.Errorf("Unexpected size of target map (got %d)", len(target))
		}

		if target["foo"] != "bar" {
			t.Errorf("Unexpected value for key 'foo' (got '%s')", target["foo"])
		}
	})
}
