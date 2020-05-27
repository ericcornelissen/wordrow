package wordmaps

import (
	"fmt"
	"testing"
)

func ExampleMatch() {
	s := "Hello world!"
	mapping := Mapping{"hello", "hey"}

	for match := range mapping.Match(s) {
		fmt.Println(match.Word)
		// Output: Hello
	}
}

func TestMappingGetFrom(t *testing.T) {
	from, to := "hello", "hey"

	t.Run("no prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{from, to}

		result := mapping.GetFrom()
		if result != from {
			t.Errorf("Unexpected value from GetFrom (was '%s')", result)
		}
	})
	t.Run("with prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{"-" + from, to}

		result := mapping.GetFrom()
		if result != from {
			t.Errorf("Unexpected value from GetFrom (was '%s')", result)
		}
	})
	t.Run("no prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{from + "-", to}

		result := mapping.GetFrom()
		if result != from {
			t.Errorf("Unexpected value from GetFrom (was '%s')", result)
		}
	})
	t.Run("with prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{"-" + from + "-", to}

		result := mapping.GetFrom()
		if result != from {
			t.Errorf("Unexpected value from GetFrom (was '%s')", result)
		}
	})
}

func TestMappingGetTo(t *testing.T) {
	from, to := "hello", "hey"

	t.Run("no prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{from, to}

		result := mapping.GetTo()
		if result != to {
			t.Errorf("Unexpected value from GetFrom (was '%s')", result)
		}
	})
	t.Run("with prefix, no suffix", func(t *testing.T) {
		mapping := Mapping{from, "-" + to}

		result := mapping.GetTo()
		if result != to {
			t.Errorf("Unexpected value from GetFrom (was '%s')", result)
		}
	})
	t.Run("no prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{from, to + "-"}

		result := mapping.GetTo()
		if result != to {
			t.Errorf("Unexpected value from GetFrom (was '%s')", result)
		}
	})
	t.Run("with prefix, with suffix", func(t *testing.T) {
		mapping := Mapping{from, "-" + to + "-"}

		result := mapping.GetTo()
		if result != to {
			t.Errorf("Unexpected value from GetFrom (was '%s')", result)
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

func TestMatch(t *testing.T) {
	t.Run("no prefix, no suffix", func(t *testing.T) {
		from, to := "hello", "hey"
		mapping := Mapping{from, to}

		t.Run("Empty input string", func(t *testing.T) {
			for match := range mapping.Match("") {
				t.Errorf("There shouldn't be any matches (got %+v)", match)
			}
		})
		t.Run("No matches", func(t *testing.T) {
			source := "This string should not contain the from"
			for match := range mapping.Match(source) {
				t.Errorf("There shouldn't be any matches (got %+v)", match)
			}
		})
		t.Run("No prefix and no suffix", func(t *testing.T) {
			rawSource := "%s there! %s, how are you?"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       0,
					End:         len(from),
				},
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       len(from) + 8,
					End:         len(from) + 8 + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
		t.Run("With prefix and no suffix", func(t *testing.T) {
			rawSource := "%s there! foo%s, how are you?"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       0,
					End:         len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
		t.Run("No prefix and with suffix", func(t *testing.T) {
			rawSource := "%sbar there! %s, how are you?"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       len(from) + 11,
					End:         len(from) + 11 + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
		t.Run("With prefix and with suffix", func(t *testing.T) {
			rawSource := "foo%s there! %sbar, how are you?"
			source := fmt.Sprintf(rawSource, from, from)

			for match := range mapping.Match(source) {
				t.Errorf("There shouldn't be any matches (got %+v)", match)
			}
		})
	})
	t.Run("with prefix, no suffix", func(t *testing.T) {
		from, to := "bar", "foo"
		mapping := Mapping{"-" + from, "-" + to}

		t.Run("Empty input string", func(t *testing.T) {
			for match := range mapping.Match("") {
				t.Errorf("There shouldn't be any matches (got %+v)", match)
			}
		})
		t.Run("No matches", func(t *testing.T) {
			source := "This string should not contain the from"
			for match := range mapping.Match(source) {
				t.Errorf("There shouldn't be any matches (got %+v)", match)
			}
		})
		t.Run("No prefix and no suffix", func(t *testing.T) {
			rawSource := "Here is a %s and there is another %s"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       10,
					End:         10 + len(from),
				},
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       len(from) + 32,
					End:         len(from) + 32 + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
		t.Run("With prefix and no suffix", func(t *testing.T) {
			rawSource := "Here is a %s and there is another pre%s"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       10,
					End:         10 + len(from),
				},
				Match{
					Full:        "pre" + from,
					Word:        from,
					Replacement: "pre" + to,
					Prefix:      "pre",
					Suffix:      "",
					Start:       len(from) + 32,
					End:         len(from) + 35 + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
		t.Run("No prefix and with suffix", func(t *testing.T) {
			rawSource := "Here is a %ssuf and there is another %s"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       len(from) + 35,
					End:         len(from) + 35 + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
		t.Run("With prefix and with suffix", func(t *testing.T) {
			rawSource := "Here is a pre%s and there is another %ssuf"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        "pre" + from,
					Word:        from,
					Replacement: "pre" + to,
					Prefix:      "pre",
					Suffix:      "",
					Start:       10,
					End:         13 + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
	})
	t.Run("no prefix, with suffix", func(t *testing.T) {
		from, to := "foo", "bar"
		mapping := Mapping{from + "-", to + "-"}

		t.Run("Empty input string", func(t *testing.T) {
			for match := range mapping.Match("") {
				t.Errorf("There shouldn't be any matches (got %+v)", match)
			}
		})
		t.Run("No matches", func(t *testing.T) {
			source := "This string should not contain the from"
			for match := range mapping.Match(source) {
				t.Errorf("There shouldn't be any matches (got %+v)", match)
			}
		})
		t.Run("No prefix and no suffix", func(t *testing.T) {
			rawSource := "Here is a %s and there is another %s"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       10,
					End:         10 + len(from),
				},
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       len(from) + 32,
					End:         len(from) + 32 + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
		t.Run("With prefix and no suffix", func(t *testing.T) {
			rawSource := "Here is a %s and there is another pre%s"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       10,
					End:         10 + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
		t.Run("No prefix and with suffix", func(t *testing.T) {
			rawSource := "Here is a %ssuf and there is another %s"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        from + "suf",
					Word:        from,
					Replacement: to + "suf",
					Prefix:      "",
					Suffix:      "suf",
					Start:       10,
					End:         13 + len(from),
				},
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       len(from) + 35,
					End:         len(from) + 35 + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
		t.Run("With prefix and with suffix", func(t *testing.T) {
			rawSource := "Here is a pre%s and there is another %ssuf"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        from + "suf",
					Word:        from,
					Replacement: to + "suf",
					Prefix:      "",
					Suffix:      "suf",
					Start:       35 + len(from),
					End:         38 + len(from) + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
	})
	t.Run("with prefix, with suffix", func(t *testing.T) {
		from, to := "foobar", "lorem"
		mapping := Mapping{"-" + from + "-", "-" + to + "-"}

		t.Run("Empty input string", func(t *testing.T) {
			for match := range mapping.Match("") {
				t.Errorf("There shouldn't be any matches (got %+v)", match)
			}
		})
		t.Run("No matches", func(t *testing.T) {
			source := "This string should not contain the from"
			for match := range mapping.Match(source) {
				t.Errorf("There shouldn't be any matches (got %+v)", match)
			}
		})
		t.Run("No prefix and no suffix", func(t *testing.T) {
			rawSource := "Here is a %s and there is another %s"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       10,
					End:         10 + len(from),
				},
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       len(from) + 32,
					End:         len(from) + 32 + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
		t.Run("With prefix and no suffix", func(t *testing.T) {
			rawSource := "Here is a %s and there is another pre%s"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       10,
					End:         10 + len(from),
				},
				Match{
					Full:        "pre" + from,
					Word:        from,
					Replacement: "pre" + to,
					Prefix:      "pre",
					Suffix:      "",
					Start:       len(from) + 32,
					End:         len(from) + 35 + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
		t.Run("No prefix and with suffix", func(t *testing.T) {
			rawSource := "Here is a %ssuf and there is another %s"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        from + "suf",
					Word:        from,
					Replacement: to + "suf",
					Prefix:      "",
					Suffix:      "suf",
					Start:       10,
					End:         13 + len(from),
				},
				Match{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       len(from) + 35,
					End:         len(from) + 35 + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
		t.Run("With prefix and with suffix", func(t *testing.T) {
			rawSource := "Here is a pre%s and there is another %ssuf"
			source := fmt.Sprintf(rawSource, from, from)

			expectedMatches := []Match{
				Match{
					Full:        "pre" + from,
					Word:        from,
					Replacement: "pre" + to,
					Prefix:      "pre",
					Suffix:      "",
					Start:       10,
					End:         13 + len(from),
				},
				Match{
					Full:        from + "suf",
					Word:        from,
					Replacement: to + "suf",
					Prefix:      "",
					Suffix:      "suf",
					Start:       35 + len(from),
					End:         38 + len(from) + len(from),
				},
			}

			i := 0
			for match := range mapping.Match(source) {
				if i >= len(expectedMatches) {
					t.Fatal("Too many matches found")
				}

				if match != expectedMatches[i] {
					t.Errorf("Unexpected match at index %d (was %+v)", i, match)
				}

				i++
			}

			if i != len(expectedMatches) {
				t.Errorf("not enough matches (got %d)", i)
			}
		})
	})
}

func TestString(t *testing.T) {
	from, to := "hello", "hey"

	t.Run("from", func(t *testing.T) {
		t.Run("to", func(t *testing.T) {
			mapping := Mapping{from, to}

			result := mapping.String()
			if result != "[hello -> hey]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
		t.Run("-to", func(t *testing.T) {
			mapping := Mapping{from, "-" + to}

			result := mapping.String()
			if result != "[hello -> -hey]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
		t.Run("to-", func(t *testing.T) {
			mapping := Mapping{from, to + "-"}

			result := mapping.String()
			if result != "[hello -> hey-]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
		t.Run("-to-", func(t *testing.T) {
			mapping := Mapping{from, "-" + to + "-"}

			result := mapping.String()
			if result != "[hello -> -hey-]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
	})
	t.Run("-from -> to", func(t *testing.T) {
		t.Run("to", func(t *testing.T) {
			mapping := Mapping{"-" + from, to}

			result := mapping.String()
			if result != "[-hello -> hey]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
		t.Run("-to", func(t *testing.T) {
			mapping := Mapping{"-" + from, "-" + to}

			result := mapping.String()
			if result != "[-hello -> -hey]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
		t.Run("to-", func(t *testing.T) {
			mapping := Mapping{"-" + from, to + "-"}

			result := mapping.String()
			if result != "[-hello -> hey-]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
		t.Run("-to-", func(t *testing.T) {
			mapping := Mapping{"-" + from, "-" + to + "-"}

			result := mapping.String()
			if result != "[-hello -> -hey-]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
	})
	t.Run("from- -> to", func(t *testing.T) {
		t.Run("to", func(t *testing.T) {
			mapping := Mapping{from + "-", to}

			result := mapping.String()
			if result != "[hello- -> hey]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
		t.Run("-to", func(t *testing.T) {
			mapping := Mapping{from + "-", "-" + to}

			result := mapping.String()
			if result != "[hello- -> -hey]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
		t.Run("to-", func(t *testing.T) {
			mapping := Mapping{from + "-", to + "-"}

			result := mapping.String()
			if result != "[hello- -> hey-]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
		t.Run("-to-", func(t *testing.T) {
			mapping := Mapping{from + "-", "-" + to + "-"}

			result := mapping.String()
			if result != "[hello- -> -hey-]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
	})
	t.Run("-from- -> to", func(t *testing.T) {
		t.Run("to", func(t *testing.T) {
			mapping := Mapping{"-" + from + "-", to}

			result := mapping.String()
			if result != "[-hello- -> hey]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
		t.Run("-to", func(t *testing.T) {
			mapping := Mapping{"-" + from + "-", "-" + to}

			result := mapping.String()
			if result != "[-hello- -> -hey]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
		t.Run("to-", func(t *testing.T) {
			mapping := Mapping{"-" + from + "-", to + "-"}

			result := mapping.String()
			if result != "[-hello- -> hey-]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
		t.Run("-to-", func(t *testing.T) {
			mapping := Mapping{"-" + from + "-", "-" + to + "-"}

			result := mapping.String()
			if result != "[-hello- -> -hey-]" {
				t.Errorf("Unexpected value from GetFrom (was '%s')", result)
			}
		})
	})
}
