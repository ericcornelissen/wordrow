package replace

import (
	"fmt"
	"testing"
)

func TestDetectAffix(t *testing.T) {
	base := "foobar"

	t.Run("simple string", func(t *testing.T) {
		substr := base
		prefix, suffix := detectAffix(substr)

		if prefix == true {
			t.Error("Expected prefix value to be false")
		}

		if suffix == true {
			t.Error("Expected suffix value to be false")
		}
	})
	t.Run("escape leading '-'", func(t *testing.T) {
		substr := fmt.Sprintf(`\-%s`, base)
		prefix, suffix := detectAffix(substr)

		if prefix == true {
			t.Error("Expected prefix value to be false")
		}

		if suffix == true {
			t.Error("Expected suffix value to be false")
		}
	})
	t.Run("escape trailing '-'", func(t *testing.T) {
		substr := fmt.Sprintf(`%s\-`, base)
		prefix, suffix := detectAffix(substr)

		if prefix == true {
			t.Error("Expected prefix value to be false")
		}

		if suffix == true {
			t.Error("Expected suffix value to be false")
		}
	})
	t.Run("escape leading & trailing '-'", func(t *testing.T) {
		substr := fmt.Sprintf(`\-%s\-`, base)
		prefix, suffix := detectAffix(substr)

		if prefix == true {
			t.Error("Expected prefix value to be false")
		}

		if suffix == true {
			t.Error("Expected suffix value to be false")
		}
	})
	t.Run("leading '-'", func(t *testing.T) {
		substr := fmt.Sprintf(`-%s`, base)
		prefix, suffix := detectAffix(substr)

		if prefix == false {
			t.Error("Expected prefix value to be true")
		}

		if suffix == true {
			t.Error("Expected suffix value to be false")
		}
	})
	t.Run("trailing '-'", func(t *testing.T) {
		substr := fmt.Sprintf(`%s-`, base)
		prefix, suffix := detectAffix(substr)

		if prefix == true {
			t.Error("Expected prefix value to be false")
		}

		if suffix == false {
			t.Error("Expected suffix value to be true")
		}
	})
	t.Run("leading & trailing '-'", func(t *testing.T) {
		substr := fmt.Sprintf(`-%s-`, base)
		prefix, suffix := detectAffix(substr)

		if prefix == false {
			t.Error("Expected prefix value to be true")
		}

		if suffix == false {
			t.Error("Expected suffix value to be true")
		}
	})
}

func TestToSafeString(t *testing.T) {
	base := "foobar"

	t.Run("simple string", func(t *testing.T) {
		substr := base
		actual := toSafeString(substr)

		expected := base
		if actual != expected {
			t.Errorf("Unexpected regular expression (%s != %s)", actual, expected)
		}
	})
	t.Run("escape leading '-'", func(t *testing.T) {
		substr := fmt.Sprintf(`\-%s`, base)
		actual := toSafeString(substr)

		expected := fmt.Sprintf(`-%s`, base)
		if actual != expected {
			t.Errorf("Unexpected regular expression (%s != %s)", actual, expected)
		}
	})
	t.Run("escape trailing '-'", func(t *testing.T) {
		substr := fmt.Sprintf(`%s\-`, base)
		actual := toSafeString(substr)

		expected := fmt.Sprintf(`%s-`, base)
		if actual != expected {
			t.Errorf("Unexpected regular expression (%s != %s)", actual, expected)
		}
	})
	t.Run("escape leading & trailing '-'", func(t *testing.T) {
		substr := fmt.Sprintf(`\-%s\-`, base)
		actual := toSafeString(substr)

		expected := fmt.Sprintf(`-%s-`, base)
		if actual != expected {
			t.Errorf("Unexpected regular expression (%s != %s)", actual, expected)
		}
	})
	t.Run("leading '-'", func(t *testing.T) {
		substr := fmt.Sprintf(`-%s`, base)
		actual := toSafeString(substr)

		expected := base
		if actual != expected {
			t.Errorf("Unexpected regular expression (%s != %s)", actual, expected)
		}
	})
	t.Run("trailing '-'", func(t *testing.T) {
		substr := fmt.Sprintf(`%s-`, base)
		actual := toSafeString(substr)

		expected := base
		if actual != expected {
			t.Errorf("Unexpected regular expression (%s != %s)", actual, expected)
		}
	})
	t.Run("leading & trailing '-'", func(t *testing.T) {
		substr := fmt.Sprintf(`-%s-`, base)
		actual := toSafeString(substr)

		expected := base
		if actual != expected {
			t.Errorf("Unexpected regular expression (%s != %s)", actual, expected)
		}
	})
	t.Run("escape paranthesis", func(t *testing.T) {
		substr := fmt.Sprintf(`(%s`, base)
		actual := toSafeString(substr)

		expected := fmt.Sprintf(`\(%s`, base)
		if actual != expected {
			t.Errorf("Unexpected regular expression (%s != %s)", actual, expected)
		}

		substr = fmt.Sprintf(`%s)`, base)
		actual = toSafeString(substr)

		expected = fmt.Sprintf(`%s\)`, base)
		if actual != expected {
			t.Errorf("Unexpected regular expression (%s != %s)", actual, expected)
		}

		substr = fmt.Sprintf(`(%s)`, base)
		actual = toSafeString(substr)

		expected = fmt.Sprintf(`\(%s\)`, base)
		if actual != expected {
			t.Errorf("Unexpected regular expression (%s != %s)", actual, expected)
		}
	})
	t.Run("escape backslashes", func(t *testing.T) {
		substr := fmt.Sprintf(`\%s`, base)
		actual := toSafeString(substr)

		expected := fmt.Sprintf(`\\%s`, base)
		if actual != expected {
			t.Errorf("Unexpected regular expression (%s != %s)", actual, expected)
		}
	})
}

func TestAllMatches(t *testing.T) {
	t.Run("empty search string", func(t *testing.T) {
		for match := range findAllMatches("", "foo") {
			t.Fatalf("Expected no matches (got '%+v')", match)
		}
	})
	t.Run("search string not containing substring", func(t *testing.T) {
		for match := range findAllMatches("hello world!", "foobar") {
			t.Fatalf("Expected no matches (got '%+v')", match)
		}
	})
	t.Run("search string containing substring once", func(t *testing.T) {
		i := 0
		for actualMatch := range findAllMatches("hello world!", "ell") {
			i++

			expectedMatch := match{
				full:   "hello",
				word:   "ell",
				prefix: "h",
				suffix: "o",
				start:  0,
				end:    5,
			}

			if actualMatch != expectedMatch {
				t.Errorf("Unexpected match (got '%+v')", actualMatch)
			}
		}

		if i != 1 {
			t.Errorf("Incorrect number of matches (got %d)", i)
		}
	})
	t.Run("search string with substring multiple times", func(t *testing.T) {
		i := 0
		for range findAllMatches("foo foobar bar", "foo") {
			i++
		}

		if i != 2 {
			t.Errorf("Incorrect number of matches (got %d)", i)
		}
	})
}

func TestMatches(t *testing.T) {
	t.Run("empty search string", func(t *testing.T) {
		for match := range matches("", "bar", "baz") {
			t.Fatalf("Expected no matches (got '%+v')", match)
		}
	})
	t.Run("search string not containing substring", func(t *testing.T) {
		t.Run("not at all", func(t *testing.T) {
			for match := range matches("hello world!", "bar", "baz") {
				t.Fatalf("Expected no matches (got '%+v')", match)
			}
		})
		t.Run("present, but with prefix", func(t *testing.T) {
			for match := range matches("hello world!", "ello", "hey") {
				t.Fatalf("Expected no matches (got '%+v')", match)
			}
		})
		t.Run("present, but with suffix", func(t *testing.T) {
			for match := range matches("hello world!", "hell", "hey") {
				t.Fatalf("Expected no matches (got '%+v')", match)
			}
		})
		t.Run("present, but with prefix & suffix", func(t *testing.T) {
			for match := range matches("hello world!", "ell", "hey") {
				t.Fatalf("Expected no matches (got '%+v')", match)
			}
		})
	})
	t.Run("search string containing substring once", func(t *testing.T) {
		t.Run("match with no prefix or suffix", func(t *testing.T) {
			for actualMatch := range matches("hello world!", "hello", "hey") {
				expectedMatch := match{
					full:        "hello",
					word:        "hello",
					replacement: "hey",
					prefix:      "",
					suffix:      "",
					start:       0,
					end:         5,
				}

				if actualMatch != expectedMatch {
					t.Errorf("Unexpected match (got '%+v')", actualMatch)
				}
			}
		})
		t.Run("match with prefix", func(t *testing.T) {
			for actualMatch := range matches("hello world!", "-ello", "hey") {
				expectedMatch := match{
					full:        "hello",
					word:        "ello",
					replacement: "hey",
					prefix:      "h",
					suffix:      "",
					start:       0,
					end:         5,
				}

				if actualMatch != expectedMatch {
					t.Errorf("Unexpected match (got '%+v')", actualMatch)
				}
			}
		})
		t.Run("match with suffix", func(t *testing.T) {
			for actualMatch := range matches("hello world!", "hell-", "hey") {
				expectedMatch := match{
					full:        "hello",
					word:        "hell",
					replacement: "hey",
					prefix:      "",
					suffix:      "o",
					start:       0,
					end:         5,
				}

				if actualMatch != expectedMatch {
					t.Errorf("Unexpected match (got '%+v')", actualMatch)
				}
			}
		})
		t.Run("match with prefix & suffix", func(t *testing.T) {
			for actualMatch := range matches("hello world!", "-ell-", "hey") {
				expectedMatch := match{
					full:        "hello",
					word:        "ell",
					replacement: "hey",
					prefix:      "h",
					suffix:      "o",
					start:       0,
					end:         5,
				}

				if actualMatch != expectedMatch {
					t.Errorf("Unexpected match (got '%+v')", actualMatch)
				}
			}
		})
		t.Run("match with prefix, keep prefix", func(t *testing.T) {
			for actualMatch := range matches("hello world!", "-ello", "-ey") {
				expectedMatch := match{
					full:        "hello",
					word:        "ello",
					replacement: "hey",
					prefix:      "h",
					suffix:      "",
					start:       0,
					end:         5,
				}

				if actualMatch != expectedMatch {
					t.Errorf("Unexpected match (got '%+v')", actualMatch)
				}
			}
		})
		t.Run("match with suffix, keep suffix", func(t *testing.T) {
			for actualMatch := range matches("hello world!", "hell-", "hey-") {
				expectedMatch := match{
					full:        "hello",
					word:        "hell",
					replacement: "heyo",
					prefix:      "",
					suffix:      "o",
					start:       0,
					end:         5,
				}

				if actualMatch != expectedMatch {
					t.Errorf("Unexpected match (got '%+v')", actualMatch)
				}
			}
		})
		t.Run("match with prefix & suffix, keep prefix", func(t *testing.T) {
			for actualMatch := range matches("hello world!", "-ell-", "-ey") {
				expectedMatch := match{
					full:        "hello",
					word:        "ell",
					replacement: "hey",
					prefix:      "h",
					suffix:      "o",
					start:       0,
					end:         5,
				}

				if actualMatch != expectedMatch {
					t.Errorf("Unexpected match (got '%+v')", actualMatch)
				}
			}
		})
		t.Run("match with prefix & suffix, keep suffix", func(t *testing.T) {
			for actualMatch := range matches("hello world!", "-ell-", "hey-") {
				expectedMatch := match{
					full:        "hello",
					word:        "ell",
					replacement: "heyo",
					prefix:      "h",
					suffix:      "o",
					start:       0,
					end:         5,
				}

				if actualMatch != expectedMatch {
					t.Errorf("Unexpected match (got '%+v')", actualMatch)
				}
			}
		})
		t.Run("match with prefix & suffix, keep both", func(t *testing.T) {
			for actualMatch := range matches("hello world!", "-ell-", "-ey-") {
				expectedMatch := match{
					full:        "hello",
					word:        "ell",
					replacement: "heyo",
					prefix:      "h",
					suffix:      "o",
					start:       0,
					end:         5,
				}

				if actualMatch != expectedMatch {
					t.Errorf("Unexpected match (got '%+v')", actualMatch)
				}
			}
		})
	})
	t.Run("search string contains UTF-8 character", func(t *testing.T) {
		for match := range matches("foobar", "\xbf", "baz") {
			t.Fatalf("Expected no matches (got '%+v')", match)
		}
	})
}
