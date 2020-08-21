package wordmaps

import "testing"

// Helper function to check if a WordMap is of the correct size and contains the
// correct values.
func checkWordMap(t *testing.T, wm StringMap, expected [][]string) {
	t.Helper()

	if len(wm) != len(expected) {
		t.Fatalf("The WordMap size should be %d (got %d)", len(expected), len(wm))
	}

	for _, expectedI := range expected {
		from, to := expectedI[0], expectedI[1]

		actualTo, ok := wm[from]
		if !ok {
			t.Errorf("Missing from-value '%s'", from)
		}

		if actualTo != to {
			t.Errorf("Incorrect to value for '%s' (got '%s')", from, actualTo)
		}
	}
}
