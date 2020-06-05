package wordmaps

import "testing"

// Helper function to check if a WordMap is of the correct size and contains the
// correct values.
func checkWordMap(t *testing.T, wm WordMap, expected [][]string) {
	t.Helper()

	if wm.Size() != len(expected) {
		t.Fatalf("The WordMap size should be %d (got %d)", len(expected), wm.Size())
	}

	for i, expectedI := range expected {
		from, to := expectedI[0], expectedI[1]

		actual := wm.GetFrom(i)
		if actual != from {
			t.Errorf("Incorrect from-value at index %d (got '%s')", i, actual)
		}

		actual = wm.GetTo(i)
		if actual != to {
			t.Errorf("Incorrect to-value at index %d (got '%s')", i, actual)
		}

		i++
	}
}
