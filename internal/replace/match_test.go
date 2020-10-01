package replace

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEmptyChannel(t *testing.T) {
	i := 0

	ch := emptyChannel()
	for range ch {
		i++
	}

	if i != 0 {
		t.Errorf("Empty channel outputted %d values", i)
	}
}

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

func TestIsValidFor(t *testing.T) {
	t.Run("match w/o prefix or suffix, query w/o prefix or suffix", func(t *testing.T) {
		query := "foobar"
		full := query
		m := &match{
			full:   []byte(full),
			word:   []byte(query),
			prefix: []byte{},
			suffix: []byte{},
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match with prefix w/o suffix, query w/o prefix or suffix", func(t *testing.T) {
		query := "bar"
		full := fmt.Sprintf("foo%s", query)
		m := &match{
			full:   []byte(full),
			word:   []byte(query),
			prefix: []byte("foo"),
			suffix: []byte{},
			start:  3,
			end:    3 + len(full),
		}

		if isValidFor(m, query) {
			t.Error("Expected match to be invalid for query")
		}
	})
	t.Run("match w/o prefix with suffix, query w/o prefix or suffix", func(t *testing.T) {
		query := "foo"
		full := fmt.Sprintf("%sbar", query)
		m := &match{
			full:   []byte(full),
			word:   []byte(query),
			prefix: []byte{},
			suffix: []byte("bar"),
			start:  3,
			end:    3 + len(full),
		}

		if isValidFor(m, query) {
			t.Error("Expected match to be invalid for query")
		}
	})
	t.Run("match with prefix and suffix, query w/o prefix or suffix", func(t *testing.T) {
		query := "freaking"
		full := fmt.Sprintf("foo%sbar", query)
		m := &match{
			full:   []byte(full),
			word:   []byte(query),
			prefix: []byte("foo"),
			suffix: []byte("bar"),
			start:  3,
			end:    3 + len(full),
		}

		if isValidFor(m, query) {
			t.Error("Expected match to be invalid for query")
		}
	})
	t.Run("match w/o prefix or suffix, query with prefix w/o suffix", func(t *testing.T) {
		query := "-bar"
		full := query[1:]
		m := &match{
			full:   []byte(full),
			word:   []byte(query[1:]),
			prefix: []byte{},
			suffix: []byte{},
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match with prefix w/o suffix, query with prefix w/o suffix", func(t *testing.T) {
		query := "-bar"
		full := fmt.Sprintf("foo%s", query[1:])
		m := &match{
			full:   []byte(full),
			word:   []byte(query[1:]),
			prefix: []byte("foo"),
			suffix: []byte{},
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match w/o prefix with suffix, query with prefix w/o suffix", func(t *testing.T) {
		query := "-bar"
		full := fmt.Sprintf("%sbar", query[1:])
		m := &match{
			full:   []byte(full),
			word:   []byte(query[1:]),
			prefix: []byte{},
			suffix: []byte("bar"),
			start:  3,
			end:    3 + len(full),
		}

		if isValidFor(m, query) {
			t.Error("Expected match to be invalid for query")
		}
	})
	t.Run("match with prefix and suffix, query with prefix w/o suffix", func(t *testing.T) {
		query := "-freaking"
		full := fmt.Sprintf("foo%sbar", query[1:])
		m := &match{
			full:   []byte(full),
			word:   []byte(query[1:]),
			prefix: []byte("foo"),
			suffix: []byte("bar"),
			start:  3,
			end:    3 + len(full),
		}

		if isValidFor(m, query) {
			t.Error("Expected match to be invalid for query")
		}
	})
	t.Run("match w/o prefix or suffix, query w/o prefix with suffix", func(t *testing.T) {
		query := "foo-"
		full := query[:len(query)-1]
		m := &match{
			full:   []byte(full),
			word:   []byte(query[:len(query)-1]),
			prefix: []byte{},
			suffix: []byte{},
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match with prefix w/o suffix, query w/o prefix with suffix", func(t *testing.T) {
		query := "foo-"
		full := fmt.Sprintf("foo%s", query[:len(query)-1])
		m := &match{
			full:   []byte(full),
			word:   []byte(query[:len(query)-1]),
			prefix: []byte("foo"),
			suffix: []byte{},
			start:  3,
			end:    3 + len(full),
		}

		if isValidFor(m, query) {
			t.Error("Expected match to be invalid for query")
		}
	})
	t.Run("match w/o prefix with suffix, query w/o prefix with suffix", func(t *testing.T) {
		query := "foo-"
		full := fmt.Sprintf("%sbar", query[:len(query)-1])
		m := &match{
			full:   []byte(full),
			word:   []byte(query[:len(query)-1]),
			prefix: []byte{},
			suffix: []byte("bar"),
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match with prefix and suffix, query w/o prefix with suffix", func(t *testing.T) {
		query := "freaking-"
		full := fmt.Sprintf("foo%sbar", query[:len(query)-1])
		m := &match{
			full:   []byte(full),
			word:   []byte(query[:len(query)-1]),
			prefix: []byte("foo"),
			suffix: []byte("bar"),
			start:  3,
			end:    3 + len(full),
		}

		if isValidFor(m, query) {
			t.Error("Expected match to be invalid for query")
		}
	})
	t.Run("match w/o prefix or suffix, query with prefix and suffix", func(t *testing.T) {
		query := "-freaking-"
		full := query[1 : len(query)-1]
		m := &match{
			full:   []byte(full),
			word:   []byte(query[1 : len(query)-1]),
			prefix: []byte{},
			suffix: []byte{},
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match with prefix w/o suffix, query w/o prefix with suffix", func(t *testing.T) {
		query := "-freaking-"
		full := fmt.Sprintf("foo%s", query[1:len(query)-1])
		m := &match{
			full:   []byte(full),
			word:   []byte(query[1 : len(query)-1]),
			prefix: []byte("foo"),
			suffix: []byte{},
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match w/o prefix with suffix, query w/o prefix with suffix", func(t *testing.T) {
		query := "-freaking-"
		full := fmt.Sprintf("%sbar", query[1:len(query)-1])
		m := &match{
			full:   []byte(full),
			word:   []byte(query[1 : len(query)-1]),
			prefix: []byte{},
			suffix: []byte("bar"),
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match with prefix and suffix, query w/o prefix with suffix", func(t *testing.T) {
		query := "-freaking-"
		full := fmt.Sprintf("foo%sbar", query[1:len(query)-1])
		m := &match{
			full:   []byte(full),
			word:   []byte(query[1 : len(query)-1]),
			prefix: []byte("foo"),
			suffix: []byte("bar"),
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
}

func TestIndicesToMatch(t *testing.T) {
	s := []byte("Lorem ipsum dolor sit amet")
	t.Run("start at 0", func(t *testing.T) {
		if len(s) < 11 {
			t.Fatal("Test string is too short")
		}

		indices := []int{
			0,
			11,
			0,
			2,
			2,
			8,
			9,
			11,
		}

		result := indicesToMatch(s, indices)
		if !bytes.Equal(result.full, s[indices[0]:indices[1]]) {
			t.Errorf("Full match incorrect (got '%s')", result.full)
		}

		if !bytes.Equal(result.word, s[indices[4]:indices[5]]) {
			t.Errorf("Word match incorrect (got '%s')", result.word)
		}

		if !bytes.Equal(result.prefix, s[indices[2]:indices[3]]) {
			t.Errorf("Prefix match incorrect (got '%s')", result.prefix)
		}

		if !bytes.Equal(result.suffix, s[indices[6]:indices[7]]) {
			t.Errorf("Suffix match incorrect (got '%s')", result.suffix)
		}

		if result.start != indices[0] {
			t.Errorf("Match start incorrect (got '%d')", result.start)
		}

		if result.end != indices[7] {
			t.Errorf("Match End incorrect (got '%d')", result.end)
		}
	})
	t.Run("in the middle", func(t *testing.T) {
		if len(s) < 18 {
			t.Fatal("Test string is too short")
		}

		indices := []int{
			6,
			17,
			6,
			8,
			8,
			14,
			15,
			17,
		}

		result := indicesToMatch(s, indices)
		if !bytes.Equal(result.full, s[indices[0]:indices[1]]) {
			t.Errorf("Full match incorrect (got '%s')", result.full)
		}

		if !bytes.Equal(result.word, s[indices[4]:indices[5]]) {
			t.Errorf("Word match incorrect (got '%s')", result.word)
		}

		if !bytes.Equal(result.prefix, s[indices[2]:indices[3]]) {
			t.Errorf("Prefix match incorrect (got '%s')", result.prefix)
		}

		if !bytes.Equal(result.suffix, s[indices[6]:indices[7]]) {
			t.Errorf("Suffix match incorrect (got '%s')", result.suffix)
		}

		if result.start != indices[0] {
			t.Errorf("Match start incorrect (got '%d')", result.start)
		}

		if result.end != indices[7] {
			t.Errorf("Match End incorrect (got '%d')", result.end)
		}
	})
	t.Run("end at `len(s)`", func(t *testing.T) {
		if len(s) < 14 {
			t.Fatal("Test string is too short")
		}

		indices := []int{
			len(s) - 14,
			len(s),
			len(s) - 14,
			len(s) - 10,
			len(s) - 10,
			len(s) - 3,
			len(s) - 2,
			len(s),
		}

		result := indicesToMatch(s, indices)
		if !bytes.Equal(result.full, s[indices[0]:indices[1]]) {
			t.Errorf("Full match incorrect (got '%s')", result.full)
		}

		if !bytes.Equal(result.word, s[indices[4]:indices[5]]) {
			t.Errorf("Word match incorrect (got '%s')", result.word)
		}

		if !bytes.Equal(result.prefix, s[indices[2]:indices[3]]) {
			t.Errorf("Prefix match incorrect (got '%s')", result.prefix)
		}

		if !bytes.Equal(result.suffix, s[indices[6]:indices[7]]) {
			t.Errorf("Suffix match incorrect (got '%s')", result.suffix)
		}

		if result.start != indices[0] {
			t.Errorf("Match start incorrect (got '%d')", result.start)
		}

		if result.end != indices[7] {
			t.Errorf("Match End incorrect (got '%d')", result.end)
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

func TestMatches(t *testing.T) {
	t.Run("empty search string", func(t *testing.T) {
		s := []byte{}
		for match := range matches(s, "bar") {
			t.Fatalf("Expected no matches (got '%+v')", match)
		}
	})
	t.Run("search string not containing substring", func(t *testing.T) {
		s := []byte("hello world!")
		t.Run("not at all", func(t *testing.T) {
			for match := range matches(s, "bar") {
				t.Fatalf("Expected no matches (got '%+v')", match)
			}
		})
		t.Run("present, but with prefix", func(t *testing.T) {
			for match := range matches(s, "ello") {
				t.Fatalf("Expected no matches (got '%+v')", match)
			}
		})
		t.Run("present, but with suffix", func(t *testing.T) {
			for match := range matches(s, "hell") {
				t.Fatalf("Expected no matches (got '%+v')", match)
			}
		})
		t.Run("present, but with prefix & suffix", func(t *testing.T) {
			for match := range matches(s, "ell") {
				t.Fatalf("Expected no matches (got '%+v')", match)
			}
		})
	})
	t.Run("search string containing substring once", func(t *testing.T) {
		s := []byte("hello world!")
		t.Run("match with no prefix or suffix", func(t *testing.T) {
			for actualMatch := range matches(s, "hello") {
				expectedMatch := match{
					full:   []byte("hello"),
					word:   []byte("hello"),
					prefix: []byte{},
					suffix: []byte{},
					start:  0,
					end:    5,
				}

				checkMatch(t, actualMatch, &expectedMatch)
			}
		})
		t.Run("match with prefix", func(t *testing.T) {
			for actualMatch := range matches(s, "-ello") {
				expectedMatch := match{
					full:   []byte("hello"),
					word:   []byte("ello"),
					prefix: []byte("h"),
					suffix: []byte{},
					start:  0,
					end:    5,
				}

				checkMatch(t, actualMatch, &expectedMatch)
			}
		})
		t.Run("match with suffix", func(t *testing.T) {
			for actualMatch := range matches(s, "hell-") {
				expectedMatch := match{
					full:   []byte("hello"),
					word:   []byte("hell"),
					prefix: []byte{},
					suffix: []byte("o"),
					start:  0,
					end:    5,
				}

				checkMatch(t, actualMatch, &expectedMatch)
			}
		})
		t.Run("match with prefix & suffix", func(t *testing.T) {
			for actualMatch := range matches(s, "-ell-") {
				expectedMatch := match{
					full:   []byte("hello"),
					word:   []byte("ell"),
					prefix: []byte("h"),
					suffix: []byte("o"),
					start:  0,
					end:    5,
				}

				checkMatch(t, actualMatch, &expectedMatch)
			}
		})
		t.Run("match with prefix, keep prefix", func(t *testing.T) {
			for actualMatch := range matches(s, "-ello") {
				expectedMatch := match{
					full:   []byte("hello"),
					word:   []byte("ello"),
					prefix: []byte("h"),
					suffix: []byte{},
					start:  0,
					end:    5,
				}

				checkMatch(t, actualMatch, &expectedMatch)
			}
		})
		t.Run("match with suffix, keep suffix", func(t *testing.T) {
			for actualMatch := range matches(s, "hell-") {
				expectedMatch := match{
					full:   []byte("hello"),
					word:   []byte("hell"),
					prefix: []byte{},
					suffix: []byte("o"),
					start:  0,
					end:    5,
				}

				checkMatch(t, actualMatch, &expectedMatch)
			}
		})
		t.Run("match with prefix & suffix, keep prefix", func(t *testing.T) {
			for actualMatch := range matches(s, "-ell-") {
				expectedMatch := match{
					full:   []byte("hello"),
					word:   []byte("ell"),
					prefix: []byte("h"),
					suffix: []byte("o"),
					start:  0,
					end:    5,
				}

				checkMatch(t, actualMatch, &expectedMatch)
			}
		})
		t.Run("match with prefix & suffix, keep suffix", func(t *testing.T) {
			for actualMatch := range matches(s, "-ell-") {
				expectedMatch := match{
					full:   []byte("hello"),
					word:   []byte("ell"),
					prefix: []byte("h"),
					suffix: []byte("o"),
					start:  0,
					end:    5,
				}

				checkMatch(t, actualMatch, &expectedMatch)
			}
		})
		t.Run("match with prefix & suffix, keep both", func(t *testing.T) {
			for actualMatch := range matches(s, "-ell-") {
				expectedMatch := match{
					full:   []byte("hello"),
					word:   []byte("ell"),
					prefix: []byte("h"),
					suffix: []byte("o"),
					start:  0,
					end:    5,
				}

				checkMatch(t, actualMatch, &expectedMatch)
			}
		})
	})
	t.Run("search string contains UTF-8 character", func(t *testing.T) {
		s := []byte("foobar")
		for match := range matches(s, "\xbf") {
			t.Fatalf("Expected no matches (got '%+v')", match)
		}
	})
}
