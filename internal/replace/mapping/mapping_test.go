package mapping

import (
	"testing"
)

func TestEndsWithSuffixSymbol(t *testing.T) {
	t.Run("No suffix symbol", func(t *testing.T) {
		result := endsWithSuffixSymbol(`foobar`)
		if result == true {
			t.Error("Expected result to be false, was true")
		}
	})
	t.Run("A suffix symbol", func(t *testing.T) {
		result := endsWithSuffixSymbol(`foo-`)
		if result == false {
			t.Error("Expected result to be true, was false")
		}
	})
	t.Run("Escaped suffix symbol", func(t *testing.T) {
		result := endsWithSuffixSymbol(`foo\-`)
		if result == true {
			t.Error("Expected result to be false, was true")
		}
	})
}

func TestStartsWithPrefixSymbol(t *testing.T) {
	t.Run("No prefix symbol", func(t *testing.T) {
		result := startsWithPrefixSymbol(`foobar`)
		if result == true {
			t.Error("Expected result to be false, was true")
		}
	})
	t.Run("A prefix symbol", func(t *testing.T) {
		result := startsWithPrefixSymbol(`-foo`)
		if result == false {
			t.Error("Expected result to be true, was false")
		}
	})
	t.Run("Escaped prefix symbol", func(t *testing.T) {
		result := startsWithPrefixSymbol(`\-foo`)
		if result == true {
			t.Error("Expected result to be false, was true")
		}
	})
}

func TestMappingFrom(t *testing.T) {
	from, to := "hello", "hey"

	t.Run("no prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{from, to}

		result := mapping.From()
		if result != from {
			t.Errorf("Unexpected value  (got '%s')", result)
		}
	})
	t.Run("with prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{"-" + from, to}

		result := mapping.From()
		if result != from {
			t.Errorf("Unexpected value  (got '%s')", result)
		}
	})
	t.Run("no prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{from + "-", to}

		result := mapping.From()
		if result != from {
			t.Errorf("Unexpected value  (got '%s')", result)
		}
	})
	t.Run("with prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{"-" + from + "-", to}

		result := mapping.From()
		if result != from {
			t.Errorf("Unexpected value  (got '%s')", result)
		}
	})
}

func TestMappingTo(t *testing.T) {
	from, to := "hello", "hey"

	t.Run("no prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{from, to}

		result := mapping.To()
		if result != to {
			t.Errorf("Unexpected value  (got '%s')", result)
		}
	})
	t.Run("with prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{from, "-" + to}

		result := mapping.To()
		if result != to {
			t.Errorf("Unexpected value  (got '%s')", result)
		}
	})
	t.Run("no prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{from, to + "-"}

		result := mapping.To()
		if result != to {
			t.Errorf("Unexpected value  (got '%s')", result)
		}
	})
	t.Run("with prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{from, "-" + to + "-"}

		result := mapping.To()
		if result != to {
			t.Errorf("Unexpected value  (got '%s')", result)
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
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
		t.Run("-to", func(t *testing.T) {
			mapping := Mapping{from, "-" + to}

			result := mapping.String()
			if result != "[hello -> -hey]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
		t.Run("to-", func(t *testing.T) {
			mapping := Mapping{from, to + "-"}

			result := mapping.String()
			if result != "[hello -> hey-]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
		t.Run("-to-", func(t *testing.T) {
			mapping := Mapping{from, "-" + to + "-"}

			result := mapping.String()
			if result != "[hello -> -hey-]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
	})
	t.Run("-from -> to", func(t *testing.T) {
		t.Run("to", func(t *testing.T) {
			mapping := Mapping{"-" + from, to}

			result := mapping.String()
			if result != "[-hello -> hey]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
		t.Run("-to", func(t *testing.T) {
			mapping := Mapping{"-" + from, "-" + to}

			result := mapping.String()
			if result != "[-hello -> -hey]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
		t.Run("to-", func(t *testing.T) {
			mapping := Mapping{"-" + from, to + "-"}

			result := mapping.String()
			if result != "[-hello -> hey-]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
		t.Run("-to-", func(t *testing.T) {
			mapping := Mapping{"-" + from, "-" + to + "-"}

			result := mapping.String()
			if result != "[-hello -> -hey-]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
	})
	t.Run("from- -> to", func(t *testing.T) {
		t.Run("to", func(t *testing.T) {
			mapping := Mapping{from + "-", to}

			result := mapping.String()
			if result != "[hello- -> hey]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
		t.Run("-to", func(t *testing.T) {
			mapping := Mapping{from + "-", "-" + to}

			result := mapping.String()
			if result != "[hello- -> -hey]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
		t.Run("to-", func(t *testing.T) {
			mapping := Mapping{from + "-", to + "-"}

			result := mapping.String()
			if result != "[hello- -> hey-]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
		t.Run("-to-", func(t *testing.T) {
			mapping := Mapping{from + "-", "-" + to + "-"}

			result := mapping.String()
			if result != "[hello- -> -hey-]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
	})
	t.Run("-from- -> to", func(t *testing.T) {
		t.Run("to", func(t *testing.T) {
			mapping := Mapping{"-" + from + "-", to}

			result := mapping.String()
			if result != "[-hello- -> hey]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
		t.Run("-to", func(t *testing.T) {
			mapping := Mapping{"-" + from + "-", "-" + to}

			result := mapping.String()
			if result != "[-hello- -> -hey]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
		t.Run("to-", func(t *testing.T) {
			mapping := Mapping{"-" + from + "-", to + "-"}

			result := mapping.String()
			if result != "[-hello- -> hey-]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
		t.Run("-to-", func(t *testing.T) {
			mapping := Mapping{"-" + from + "-", "-" + to + "-"}

			result := mapping.String()
			if result != "[-hello- -> -hey-]" {
				t.Errorf("Unexpected value  (got '%s')", result)
			}
		})
	})
}
