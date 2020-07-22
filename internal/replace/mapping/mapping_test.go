package mapping

import (
	"testing"
)

func TestMappingGetFrom(t *testing.T) {
	from, to := "hello", "hey"

	t.Run("no prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{from, to}

		result := mapping.GetFrom()
		if result != from {
			t.Errorf("Unexpected value from GetFrom (got '%s')", result)
		}
	})
	t.Run("with prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{"-" + from, to}

		result := mapping.GetFrom()
		if result != from {
			t.Errorf("Unexpected value from GetFrom (got '%s')", result)
		}
	})
	t.Run("no prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{from + "-", to}

		result := mapping.GetFrom()
		if result != from {
			t.Errorf("Unexpected value from GetFrom (got '%s')", result)
		}
	})
	t.Run("with prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{"-" + from + "-", to}

		result := mapping.GetFrom()
		if result != from {
			t.Errorf("Unexpected value from GetFrom (got '%s')", result)
		}
	})
}

func TestMappingGetTo(t *testing.T) {
	from, to := "hello", "hey"

	t.Run("no prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{from, to}

		result := mapping.GetTo()
		if result != to {
			t.Errorf("Unexpected value from GetFrom (got '%s')", result)
		}
	})
	t.Run("with prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{from, "-" + to}

		result := mapping.GetTo()
		if result != to {
			t.Errorf("Unexpected value from GetFrom (got '%s')", result)
		}
	})
	t.Run("no prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{from, to + "-"}

		result := mapping.GetTo()
		if result != to {
			t.Errorf("Unexpected value from GetFrom (got '%s')", result)
		}
	})
	t.Run("with prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{from, "-" + to + "-"}

		result := mapping.GetTo()
		if result != to {
			t.Errorf("Unexpected value from GetFrom (got '%s')", result)
		}
	})
}

func TestGetReplacement(t *testing.T) {
	from, to := "hello", "hey"

	t.Run("no prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{from, to}

		result := mapping.getReplacement("", "")
		if result != to {
			t.Errorf("Unexpected replacement given no prefix or suffix (got '%s')", result)
		}

		result = mapping.getReplacement("foo", "")
		if result != to {
			t.Errorf("Unexpected replacement given a prefix but no suffix (got '%s')", result)
		}

		result = mapping.getReplacement("", "bar")
		if result != to {
			t.Errorf("Unexpected replacement given no prefix but a suffix (got '%s')", result)
		}

		result = mapping.getReplacement("foo", "bar")
		if result != to {
			t.Errorf("Unexpected replacement given a prefix and suffix (got '%s')", result)
		}
	})
	t.Run("with prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{from, "-" + to}

		result := mapping.getReplacement("", "")
		if result != to {
			t.Errorf("Unexpected replacement given no prefix or suffix (got '%s')", result)
		}

		result = mapping.getReplacement("foo", "")
		if result != "foo"+to {
			t.Errorf("Unexpected replacement given a prefix but no suffix (got '%s')", result)
		}

		result = mapping.getReplacement("", "bar")
		if result != to {
			t.Errorf("Unexpected replacement given no prefix but a suffix (got '%s')", result)
		}

		result = mapping.getReplacement("foo", "bar")
		if result != "foo"+to {
			t.Errorf("Unexpected replacement given a prefix and suffix (got '%s')", result)
		}
	})
	t.Run("no prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{from, to + "-"}

		result := mapping.getReplacement("", "")
		if result != to {
			t.Errorf("Unexpected replacement given no prefix or suffix (got '%s')", result)
		}

		result = mapping.getReplacement("foo", "")
		if result != to {
			t.Errorf("Unexpected replacement given a prefix but no suffix (got '%s')", result)
		}

		result = mapping.getReplacement("", "bar")
		if result != to+"bar" {
			t.Errorf("Unexpected replacement given no prefix but a suffix (got '%s')", result)
		}

		result = mapping.getReplacement("foo", "bar")
		if result != to+"bar" {
			t.Errorf("Unexpected replacement given a prefix and suffix (got '%s')", result)
		}
	})
	t.Run("with prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{from, "-" + to + "-"}

		result := mapping.getReplacement("", "")
		if result != to {
			t.Errorf("Unexpected replacement given no prefix or suffix (got '%s')", result)
		}

		result = mapping.getReplacement("foo", "")
		if result != "foo"+to {
			t.Errorf("Unexpected replacement given a prefix but no suffix (got '%s')", result)
		}

		result = mapping.getReplacement("", "bar")
		if result != to+"bar" {
			t.Errorf("Unexpected replacement given no prefix but a suffix (got '%s')", result)
		}

		result = mapping.getReplacement("foo", "bar")
		if result != "foo"+to+"bar" {
			t.Errorf("Unexpected replacement given a prefix and suffix (got '%s')", result)
		}
	})
}

func TestString(t *testing.T) {
	from, to := "hello", "hey"

	t.Run("from", func(t *testing.T) {
		t.Run("to", func(t *testing.T) {
			mapping := Mapping{from, to}

			result := mapping.String()
			if result != "[hello -> hey]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
		t.Run("-to", func(t *testing.T) {
			mapping := Mapping{from, "-" + to}

			result := mapping.String()
			if result != "[hello -> -hey]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
		t.Run("to-", func(t *testing.T) {
			mapping := Mapping{from, to + "-"}

			result := mapping.String()
			if result != "[hello -> hey-]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
		t.Run("-to-", func(t *testing.T) {
			mapping := Mapping{from, "-" + to + "-"}

			result := mapping.String()
			if result != "[hello -> -hey-]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
	})
	t.Run("-from -> to", func(t *testing.T) {
		t.Run("to", func(t *testing.T) {
			mapping := Mapping{"-" + from, to}

			result := mapping.String()
			if result != "[-hello -> hey]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
		t.Run("-to", func(t *testing.T) {
			mapping := Mapping{"-" + from, "-" + to}

			result := mapping.String()
			if result != "[-hello -> -hey]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
		t.Run("to-", func(t *testing.T) {
			mapping := Mapping{"-" + from, to + "-"}

			result := mapping.String()
			if result != "[-hello -> hey-]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
		t.Run("-to-", func(t *testing.T) {
			mapping := Mapping{"-" + from, "-" + to + "-"}

			result := mapping.String()
			if result != "[-hello -> -hey-]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
	})
	t.Run("from- -> to", func(t *testing.T) {
		t.Run("to", func(t *testing.T) {
			mapping := Mapping{from + "-", to}

			result := mapping.String()
			if result != "[hello- -> hey]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
		t.Run("-to", func(t *testing.T) {
			mapping := Mapping{from + "-", "-" + to}

			result := mapping.String()
			if result != "[hello- -> -hey]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
		t.Run("to-", func(t *testing.T) {
			mapping := Mapping{from + "-", to + "-"}

			result := mapping.String()
			if result != "[hello- -> hey-]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
		t.Run("-to-", func(t *testing.T) {
			mapping := Mapping{from + "-", "-" + to + "-"}

			result := mapping.String()
			if result != "[hello- -> -hey-]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
	})
	t.Run("-from- -> to", func(t *testing.T) {
		t.Run("to", func(t *testing.T) {
			mapping := Mapping{"-" + from + "-", to}

			result := mapping.String()
			if result != "[-hello- -> hey]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
		t.Run("-to", func(t *testing.T) {
			mapping := Mapping{"-" + from + "-", "-" + to}

			result := mapping.String()
			if result != "[-hello- -> -hey]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
		t.Run("to-", func(t *testing.T) {
			mapping := Mapping{"-" + from + "-", to + "-"}

			result := mapping.String()
			if result != "[-hello- -> hey-]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
		t.Run("-to-", func(t *testing.T) {
			mapping := Mapping{"-" + from + "-", "-" + to + "-"}

			result := mapping.String()
			if result != "[-hello- -> -hey-]" {
				t.Errorf("Unexpected value from GetFrom (got '%s')", result)
			}
		})
	})
}
