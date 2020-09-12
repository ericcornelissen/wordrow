package wordmaps

import "testing"

func TestAddToMapping(t *testing.T) {
	t.Run("add single mapping", func(t *testing.T) {
		from, to := "foo", "bar"
		target := make(map[string]string, 1)

		addToMapping(target, []string{from}, to)
		if len(target) != 1 {
			t.Fatalf("Unexpected size of map (got %d)", len(target))
		}

		if target[from] != to {
			t.Errorf("Unexpected value for '%s' (got '%s')", from, target[from])
		}
	})
	t.Run("add multiple mappings", func(t *testing.T) {
		from1, from2, to := "bar", "baz", "foo"
		target := make(map[string]string, 1)

		addToMapping(target, []string{from1, from2}, to)
		if len(target) != 2 {
			t.Fatalf("Unexpected size of map (got %d)", len(target))
		}

		if target[from1] != to {
			t.Errorf("Unexpected value for '%s' (got '%s')", from1, target[from1])
		}

		if target[from2] != to {
			t.Errorf("Unexpected value for '%s' (got '%s')", from1, target[from2])
		}
	})
	t.Run("override in target", func(t *testing.T) {
		from, to1, to2 := "foo", "bar", "baz"
		target := make(map[string]string, 1)
		target[from] = to1

		addToMapping(target, []string{from}, to2)
		if len(target) != 1 {
			t.Fatalf("Unexpected size of map (got %d)", len(target))
		}

		if target[from] != to2 {
			t.Errorf("Unexpected value for '%s' (got '%s')", from, target[from])
		}
	})
}
