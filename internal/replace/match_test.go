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

func TestIsValidForNoPrefixNoSuffix(t *testing.T) {
	t.Run("match w/o prefix or suffix, query w/o prefix or suffix", func(t *testing.T) {
		query := "foobar"
		full := query
		m := &match{
			full:   full,
			word:   query,
			prefix: "",
			suffix: "",
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
			full:   full,
			word:   query,
			prefix: "foo",
			suffix: "",
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
			full:   full,
			word:   query,
			prefix: "",
			suffix: "bar",
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
			full:   full,
			word:   query,
			prefix: "foo",
			suffix: "bar",
			start:  3,
			end:    3 + len(full),
		}

		if isValidFor(m, query) {
			t.Error("Expected match to be invalid for query")
		}
	})
}

func TestIsValidForWithPrefixNoSuffix(t *testing.T) {
	t.Run("match w/o prefix or suffix", func(t *testing.T) {
		query := "-bar"
		full := query[1:]
		m := &match{
			full:   full,
			word:   query[1:],
			prefix: "",
			suffix: "",
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match with prefix w/o suffix", func(t *testing.T) {
		query := "-bar"
		full := fmt.Sprintf("foo%s", query[1:])
		m := &match{
			full:   full,
			word:   query[1:],
			prefix: "foo",
			suffix: "",
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match w/o prefix with suffix", func(t *testing.T) {
		query := "-bar"
		full := fmt.Sprintf("%sbar", query[1:])
		m := &match{
			full:   full,
			word:   query[1:],
			prefix: "",
			suffix: "bar",
			start:  3,
			end:    3 + len(full),
		}

		if isValidFor(m, query) {
			t.Error("Expected match to be invalid for query")
		}
	})
	t.Run("match with prefix and suffix", func(t *testing.T) {
		query := "-freaking"
		full := fmt.Sprintf("foo%sbar", query[1:])
		m := &match{
			full:   full,
			word:   query[1:],
			prefix: "foo",
			suffix: "bar",
			start:  3,
			end:    3 + len(full),
		}

		if isValidFor(m, query) {
			t.Error("Expected match to be invalid for query")
		}
	})
}

func TestIsValidForNoPrefixWithSuffix(t *testing.T) {
	t.Run("match w/o prefix or suffix", func(t *testing.T) {
		query := "foo-"
		full := query[:len(query)-1]
		m := &match{
			full:   full,
			word:   query[:len(query)-1],
			prefix: "",
			suffix: "",
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match with prefix w/o suffix", func(t *testing.T) {
		query := "foo-"
		full := fmt.Sprintf("foo%s", query[:len(query)-1])
		m := &match{
			full:   full,
			word:   query[:len(query)-1],
			prefix: "foo",
			suffix: "",
			start:  3,
			end:    3 + len(full),
		}

		if isValidFor(m, query) {
			t.Error("Expected match to be invalid for query")
		}
	})
	t.Run("match w/o prefix with suffix", func(t *testing.T) {
		query := "foo-"
		full := fmt.Sprintf("%sbar", query[:len(query)-1])
		m := &match{
			full:   full,
			word:   query[:len(query)-1],
			prefix: "",
			suffix: "bar",
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match with prefix and suffix", func(t *testing.T) {
		query := "freaking-"
		full := fmt.Sprintf("foo%sbar", query[:len(query)-1])
		m := &match{
			full:   full,
			word:   query[:len(query)-1],
			prefix: "foo",
			suffix: "bar",
			start:  3,
			end:    3 + len(full),
		}

		if isValidFor(m, query) {
			t.Error("Expected match to be invalid for query")
		}
	})
}

func TestIsValidForWithPrefixWithSuffix(t *testing.T) {
	t.Run("match w/o prefix or suffix", func(t *testing.T) {
		query := "-freaking-"
		full := query[1 : len(query)-1]
		m := &match{
			full:   full,
			word:   query[1 : len(query)-1],
			prefix: "",
			suffix: "",
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match with prefix w/o suffix", func(t *testing.T) {
		query := "-freaking-"
		full := fmt.Sprintf("foo%s", query[1:len(query)-1])
		m := &match{
			full:   full,
			word:   query[1 : len(query)-1],
			prefix: "foo",
			suffix: "",
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match w/o prefix with suffix", func(t *testing.T) {
		query := "-freaking-"
		full := fmt.Sprintf("%sbar", query[1:len(query)-1])
		m := &match{
			full:   full,
			word:   query[1 : len(query)-1],
			prefix: "",
			suffix: "bar",
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
	t.Run("match with prefix and suffix", func(t *testing.T) {
		query := "-freaking-"
		full := fmt.Sprintf("foo%sbar", query[1:len(query)-1])
		m := &match{
			full:   full,
			word:   query[1 : len(query)-1],
			prefix: "foo",
			suffix: "bar",
			start:  3,
			end:    3 + len(full),
		}

		if !isValidFor(m, query) {
			t.Error("Expected match to be valid for query")
		}
	})
}

func TestIndicesToMatchAtStartOfString(t *testing.T) {
	s := "Lorem ipsum dolor sit amet"
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
	if result.full != s[indices[0]:indices[1]] {
		t.Errorf("Full match incorrect (got '%s')", result.full)
	}

	if result.word != s[indices[4]:indices[5]] {
		t.Errorf("Full match incorrect (got '%s')", result.word)
	}

	if result.prefix != s[indices[2]:indices[3]] {
		t.Errorf("Full match incorrect (got '%s')", result.prefix)
	}

	if result.suffix != s[indices[6]:indices[7]] {
		t.Errorf("Full match incorrect (got '%s')", result.suffix)
	}

	if result.start != indices[0] {
		t.Errorf("Full match incorrect (got '%s')", result.suffix)
	}

	if result.end != indices[7] {
		t.Errorf("Full match incorrect (got '%s')", result.suffix)
	}
}

func TestIndicesToMatchInMiddleOfString(t *testing.T) {
	s := "Lorem ipsum dolor sit amet"
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
	if result.full != s[indices[0]:indices[1]] {
		t.Errorf("Full match incorrect (got '%s')", result.full)
	}

	if result.word != s[indices[4]:indices[5]] {
		t.Errorf("Full match incorrect (got '%s')", result.word)
	}

	if result.prefix != s[indices[2]:indices[3]] {
		t.Errorf("Full match incorrect (got '%s')", result.prefix)
	}

	if result.suffix != s[indices[6]:indices[7]] {
		t.Errorf("Full match incorrect (got '%s')", result.suffix)
	}

	if result.start != indices[0] {
		t.Errorf("Full match incorrect (got '%s')", result.suffix)
	}

	if result.end != indices[7] {
		t.Errorf("Full match incorrect (got '%s')", result.suffix)
	}
}

func TestIndicesToMatchAtEndOfString(t *testing.T) {
	s := "Lorem ipsum dolor sit amet"
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
	if result.full != s[indices[0]:indices[1]] {
		t.Errorf("Full match incorrect (got '%s')", result.full)
	}

	if result.word != s[indices[4]:indices[5]] {
		t.Errorf("Full match incorrect (got '%s')", result.word)
	}

	if result.prefix != s[indices[2]:indices[3]] {
		t.Errorf("Full match incorrect (got '%s')", result.prefix)
	}

	if result.suffix != s[indices[6]:indices[7]] {
		t.Errorf("Full match incorrect (got '%s')", result.suffix)
	}

	if result.start != indices[0] {
		t.Errorf("Full match incorrect (got '%s')", result.suffix)
	}

	if result.end != indices[7] {
		t.Errorf("Full match incorrect (got '%s')", result.suffix)
	}
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
	t.Run("escape parenthesis", func(t *testing.T) {
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

func TestMatchesFindNothing(t *testing.T) {
	t.Run("not at all", func(t *testing.T) {
		ch := matches("hello world!", "bar")
		result := drain(ch)
		if len(result) > 0 {
			t.Fatalf("Expected no matches (got %d)", len(result))
		}
	})
	t.Run("present, but with prefix", func(t *testing.T) {
		ch := matches("hello world!", "ello")
		result := drain(ch)
		if len(result) > 0 {
			t.Fatalf("Expected no matches (got %d)", len(result))
		}
	})
	t.Run("present, but with suffix", func(t *testing.T) {
		ch := matches("hello world!", "hell")
		result := drain(ch)
		if len(result) > 0 {
			t.Fatalf("Expected no matches (got %d)", len(result))
		}
	})
	t.Run("present, but with prefix & suffix", func(t *testing.T) {
		ch := matches("hello world!", "ell")
		result := drain(ch)
		if len(result) > 0 {
			t.Fatalf("Expected no matches (got %d)", len(result))
		}
	})
}

func TestMatchesFindSomething(t *testing.T) {
	t.Run("match without prefix or suffix", func(t *testing.T) {
		ch := matches("hello world!", "hello")
		result := drain(ch)
		if len(result) != 1 {
			t.Fatalf("Expected one match (got %d)", len(result))
		}

		actualMatch := result[0]
		expectedMatch := match{
			full:   "hello",
			word:   "hello",
			prefix: "",
			suffix: "",
			start:  0,
			end:    5,
		}

		if actualMatch != expectedMatch {
			t.Errorf("Unexpected match (got '%+v')", actualMatch)
		}
	})
	t.Run("match with prefix", func(t *testing.T) {
		ch := matches("hello world!", "-ello")
		result := drain(ch)
		if len(result) != 1 {
			t.Fatalf("Expected one match (got %d)", len(result))
		}

		actualMatch := result[0]
		expectedMatch := match{
			full:   "hello",
			word:   "ello",
			prefix: "h",
			suffix: "",
			start:  0,
			end:    5,
		}

		if actualMatch != expectedMatch {
			t.Errorf("Unexpected match (got '%+v')", actualMatch)
		}
	})
	t.Run("match with suffix", func(t *testing.T) {
		ch := matches("hello world!", "hell-")
		result := drain(ch)
		if len(result) != 1 {
			t.Fatalf("Expected one match (got %d)", len(result))
		}

		actualMatch := result[0]
		expectedMatch := match{
			full:   "hello",
			word:   "hell",
			prefix: "",
			suffix: "o",
			start:  0,
			end:    5,
		}

		if actualMatch != expectedMatch {
			t.Errorf("Unexpected match (got '%+v')", actualMatch)
		}
	})
	t.Run("match with prefix & suffix", func(t *testing.T) {
		ch := matches("hello world!", "-ell-")
		result := drain(ch)
		if len(result) != 1 {
			t.Fatalf("Expected one match (got %d)", len(result))
		}

		actualMatch := result[0]
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
	})
}
