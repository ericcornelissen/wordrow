package replace

import "testing"

// Drain empties a channel and returns all values as a slice.
func drain(ch chan *match) (matches []match) {
	for match := range ch {
		matches = append(matches, *match)
	}

	return matches
}

// ReportIncorrectReplacement pretty prints an error by the replacement
// functionality.
func reportIncorrectReplacement(t *testing.T, expected, actual string) {
	t.Helper()
	t.Errorf(`Replacement did not work as intended
		expected : '%s'
		got      : '%s'
	`, expected, actual)
}
