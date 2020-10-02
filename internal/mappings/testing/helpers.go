package testing

type testingT interface {
	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Helper()
}

// CheckMapping checks if a mapping is of the correct size and contains the
// correct values. This is a test helper (i.e. it will call t.Helper()).
func CheckMapping(
	t testingT,
	mapping map[string]string,
	expected [][]string,
) {
	t.Helper()

	if len(mapping) != len(expected) {
		t.Fatalf("The mapping size should be %d (got %d)", len(expected), len(mapping))
	}

	for _, expectedI := range expected {
		from, to := expectedI[0], expectedI[1]

		actualTo, ok := mapping[from]
		if !ok {
			t.Errorf("Missing from-value '%s'", from)
			continue
		}

		if actualTo != to {
			t.Errorf("Incorrect to value for '%s' (got '%s')", from, actualTo)
		}
	}
}
