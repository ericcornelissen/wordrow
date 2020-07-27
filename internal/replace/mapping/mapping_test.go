package mapping

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
				{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       0,
					End:         len(from),
				},
				{
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
				{
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
				{
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
				{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       10,
					End:         10 + len(from),
				},
				{
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
				{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       10,
					End:         10 + len(from),
				},
				{
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
				{
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
				{
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
				{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       10,
					End:         10 + len(from),
				},
				{
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
				{
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
				{
					Full:        from + "suf",
					Word:        from,
					Replacement: to + "suf",
					Prefix:      "",
					Suffix:      "suf",
					Start:       10,
					End:         13 + len(from),
				},
				{
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
				{
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
				{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       10,
					End:         10 + len(from),
				},
				{
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
				{
					Full:        from,
					Word:        from,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       10,
					End:         10 + len(from),
				},
				{
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
				{
					Full:        from + "suf",
					Word:        from,
					Replacement: to + "suf",
					Prefix:      "",
					Suffix:      "suf",
					Start:       10,
					End:         13 + len(from),
				},
				{
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
				{
					Full:        "pre" + from,
					Word:        from,
					Replacement: "pre" + to,
					Prefix:      "pre",
					Suffix:      "",
					Start:       10,
					End:         13 + len(from),
				},
				{
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
	t.Run("escaped prefix", func(t *testing.T) {
		from, unescapedFrom, to := `\-bar`, `-bar`, `bar`
		mapping := Mapping{from, to}

		t.Run("a match", func(t *testing.T) {
			rawSource := "foo %s"
			source := fmt.Sprintf(rawSource, unescapedFrom)
			fmt.Printf(source)

			expectedMatches := []Match{
				{
					Full:        unescapedFrom,
					Word:        unescapedFrom,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       4,
					End:         4 + len(unescapedFrom),
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
		t.Run("from value with a prefix", func(t *testing.T) {
			for match := range mapping.Match("foobar") {
				t.Errorf("There shouldn't be any matches (got %+v)", match)
			}
		})
	})
	t.Run("escaped suffix", func(t *testing.T) {
		from, unescapedFrom, to := `foo\-`, `foo-`, `foo`
		mapping := Mapping{from, to}

		t.Run("a match", func(t *testing.T) {
			rawSource := "%s bar"
			source := fmt.Sprintf(rawSource, unescapedFrom)
			fmt.Printf(source)

			expectedMatches := []Match{
				{
					Full:        unescapedFrom,
					Word:        unescapedFrom,
					Replacement: to,
					Prefix:      "",
					Suffix:      "",
					Start:       0,
					End:         0 + len(unescapedFrom),
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
		t.Run("from value with a prefix", func(t *testing.T) {
			for match := range mapping.Match("foobar") {
				t.Errorf("There shouldn't be any matches (got %+v)", match)
			}
		})
	})
}

func TestUnconventionalMappings(t *testing.T) {
	t.Run("non UTF-8 characters", func(t *testing.T) {
		mapping := Mapping{"\xbf", "a"}

		for match := range mapping.Match("Hello world!") {
			t.Errorf("Expected no matches (matched '%s')", match.Full)
		}

		for match := range mapping.Match("Hello \xbf!") {
			t.Errorf("Expected no matches (matched '%s')", match.Full)
		}
	})
	t.Run("Invalid Regular Expression", func(t *testing.T) {
		from, to := "(foo", "(bar"
		mapping := Mapping{from, to}

		rawSource := "Foo bar %s bar)"
		source := fmt.Sprintf(rawSource, from)

		expectedMatches := []Match{
			{
				Full:        from,
				Word:        from,
				Replacement: to,
				Prefix:      "",
				Suffix:      "",
				Start:       8,
				End:         8 + len(from),
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