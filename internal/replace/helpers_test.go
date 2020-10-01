package replace

import (
	"bytes"
	"testing"
)

// CheckMatch checks if two matchers are equal.
func checkMatch(t *testing.T, actualMatch, expectedMatch *match) {
	if !bytes.Equal(actualMatch.full, expectedMatch.full) {
		t.Errorf("Full match incorrect (got '%s')", actualMatch.full)
	}

	if !bytes.Equal(actualMatch.word, expectedMatch.word) {
		t.Errorf("Word match incorrect (got '%s')", actualMatch.word)
	}

	if !bytes.Equal(actualMatch.prefix, expectedMatch.prefix) {
		t.Errorf("Prefix match incorrect (got '%s')", actualMatch.prefix)
	}

	if !bytes.Equal(actualMatch.suffix, expectedMatch.suffix) {
		t.Errorf("Suffix match incorrect (got '%s')", actualMatch.suffix)
	}

	if actualMatch.start != expectedMatch.start {
		t.Errorf("Match start incorrect (got '%d')", actualMatch.start)
	}

	if actualMatch.end != expectedMatch.end {
		t.Errorf("Match End incorrect (got '%d')", actualMatch.end)
	}
}

// ReportIncorrectReplacement pretty prints an error by the replacement
// functionality.
func reportIncorrectReplacement(t *testing.T, expected, actual []byte) {
	t.Helper()
	t.Errorf(`Replacement did not work as intended
		expected : '%s'
		got      : '%s'
	`, expected, actual)
}
