package mapping

import (
	"testing"
)

func TestCleanStringToMatch(t *testing.T) {
	t.Run("Unescapes hyphens", func(t *testing.T) {
		in, expected := `\-`, `-`
		actual := cleanStringToMatch(in)
		if actual != expected {
			t.Errorf("Unexpected result (got '%s')", actual)
		}
	})
	t.Run("Escapes paranthesis", func(t *testing.T) {
		in, expected := `(`, `\(`
		actual := cleanStringToMatch(in)
		if actual != expected {
			t.Errorf("Unexpected result (got '%s')", actual)
		}

		in, expected = `)`, `\)`
		actual = cleanStringToMatch(in)
		if actual != expected {
			t.Errorf("Unexpected result (got '%s')", actual)
		}
	})
	t.Run("Escapes backslash", func(t *testing.T) {
		in, expected := `\`, `\\`
		actual := cleanStringToMatch(in)
		if actual != expected {
			t.Errorf("Unexpected result (got '%s')", actual)
		}
	})
}

func TestGetAllMatches(t *testing.T) {
	t.Run("Empty search string", func(t *testing.T) {
		for match := range getAllMatches("", "foo") {
			t.Fatalf("Expected no matches (got '%+v')", match)
		}
	})
	t.Run("search string, not found", func(t *testing.T) {
		for match := range getAllMatches("Hello world!", "foobar") {
			t.Fatalf("Expected no matches (got '%+v')", match)
		}
	})
	t.Run("search string, a match", func(t *testing.T) {
		for actualMatch := range getAllMatches("Hello world!", "ell") {
			expectedMatch := Match{
				Full:   "Hello",
				Word:   "ell",
				Prefix: "H",
				Suffix: "o",
				Start:  0,
				End:    5,
			}

			if actualMatch != expectedMatch {
				t.Errorf("Unexpected match (got '%+v')", actualMatch)
			}
		}
	})
	t.Run("Search non UTF-8 character", func(t *testing.T) {
		for match := range getAllMatches("foobar", "\xbf") {
			t.Fatalf("Expected no matches (got '%+v')", match)
		}
	})
}
