package replacer

import "testing"

// ReportIncorrectReplacement pretty prints an error by the replacement
// functionality.
func reportIncorrectReplacement(t *testing.T, expected, actual string) {
	t.Helper()
	t.Errorf(`Replacement did not work as intended
		expected : '%s'
		got      : '%s'
	`, expected, actual)
}
