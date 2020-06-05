package wordmaps

import (
	"strings"
	"testing"
)

func TestTrimAll(t *testing.T) {
	subject := []string{
		"a",
		" b",
		"c ",
		" d ",
	}

	trimSpaceAll(subject)

	if subject[0] != "a" {
		t.Errorf("Incorrect first value (got '%s')", subject[0])
	}

	if subject[1] != "b" {
		t.Errorf("Incorrect second value (got '%s')", subject[1])
	}

	if subject[2] != "c" {
		t.Errorf("Incorrect third value (got '%s')", subject[2])
	}

	if subject[3] != "d" {
		t.Errorf("Incorrect fourth value (got '%s')", subject[3])
	}
}

func TestAnyString(t *testing.T) {
	condition := func(s string) bool {
		return s == ""
	}

	t.Run("empty list", func(t *testing.T) {
		result := anyString([]string{}, condition)

		if result == true {
			t.Error("Unexpected result `true` for empty list")
		}
	})
	t.Run("condition never holds", func(t *testing.T) {
		result := anyString([]string{"a", "b", "c"}, condition)

		if result == true {
			t.Error("Unexpected result `true` for list")
		}
	})
	t.Run("condition holds once", func(t *testing.T) {
		result := anyString([]string{"a", "", "c"}, condition)

		if result == false {
			t.Error("Unexpected result `false` for list")
		}
	})
	t.Run("condition holds often", func(t *testing.T) {
		result := anyString([]string{"", "", "c"}, condition)

		if result == false {
			t.Error("Unexpected result `false` for list")
		}
	})
}

func TestIsEmptyString(t *testing.T) {
	result := isEmptyString("")
	if result == false {
		t.Error("Unexpected result `false` for empty string")
	}

	result = isEmptyString("Hello world!")
	if result == true {
		t.Error("Unexpected result `true` for non-empty string")
	}
}

func TestCsvOneRow(t *testing.T) {
	csv := `cat,dog`
	wm, err := parseCsvFile(&csv)

	if err != nil {
		t.Fatalf("Error should be nil for this test (Error: %s)", err)
	}

	if wm.Size() != 1 {
		t.Fatalf("The WordMap size should be 1 (was %d)", wm.Size())
	}

	actual, expected := wm.GetFrom(0), "cat"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
	}

	actual, expected = wm.GetTo(0), "dog"
	if actual != expected {
		t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
	}
}

func TestCsvMultipleRows(t *testing.T) {
	csv := `
		cat,dog
		horse,zebra
	`
	wm, err := parseCsvFile(&csv)

	if err != nil {
		t.Fatalf("Error should be nil for this test (Error: %s)", err)
	}

	if wm.Size() != 2 {
		t.Fatalf("The WordMap size should be 2 (was %d)", wm.Size())
	}

	actual, expected := wm.GetFrom(0), "cat"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
	}

	actual, expected = wm.GetTo(0), "dog"
	if actual != expected {
		t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
	}

	actual, expected = wm.GetFrom(1), "horse"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
	}

	actual, expected = wm.GetTo(1), "zebra"
	if actual != expected {
		t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
	}
}

func TestCsvManyColumns(t *testing.T) {
	csv := `cat,dog,horse`
	wm, err := parseCsvFile(&csv)

	if err != nil {
		t.Fatalf("Error should be nil for this test (Error: %s)", err)
	}

	if wm.Size() != 2 {
		t.Fatalf("The WordMap size should be 2 (was %d)", wm.Size())
	}

	actual, expected := wm.GetFrom(0), "cat"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
	}

	actual, expected = wm.GetTo(0), "horse"
	if actual != expected {
		t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
	}

	actual, expected = wm.GetFrom(1), "dog"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
	}

	actual, expected = wm.GetTo(1), "horse"
	if actual != expected {
		t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
	}
}

func TestCsvEmptyColumnValues(t *testing.T) {
	t.Run("Empty from value", func(t *testing.T) {
		csv := `,bar`

		_, err := parseCsvFile(&csv)

		if err == nil {
			t.Errorf("Error should be set if the from value is empty")
		}
	})
	t.Run("Empty to value", func(t *testing.T) {
		csv := `foo,`

		_, err := parseCsvFile(&csv)

		if err == nil {
			t.Errorf("Error should be set if the to value is empty")
		}
	})
}

func TestCsvIgnoreEmptyLines(t *testing.T) {
	csv := `
		cat,dog

		horse,zebra
	`
	wm, err := parseCsvFile(&csv)

	if err != nil {
		t.Fatalf("Error should be nil for this test (Error: %s)", err)
	}

	if wm.Size() != 2 {
		t.Fatalf("The WordMap size should be 2 (was %d)", wm.Size())
	}

	actual, expected := wm.GetFrom(0), "cat"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
	}

	actual, expected = wm.GetTo(0), "dog"
	if actual != expected {
		t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
	}

	actual, expected = wm.GetFrom(1), "horse"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
	}

	actual, expected = wm.GetTo(1), "zebra"
	if actual != expected {
		t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
	}
}

func TestCsvIgnoresWhitespaceInRow(t *testing.T) {
	csv := `
		cat, dog
		horse  , zebra
	`
	wm, err := parseCsvFile(&csv)

	if err != nil {
		t.Fatalf("Error should be nil for this test (Error: %s)", err)
	}

	if wm.Size() != 2 {
		t.Fatalf("The WordMap size should be 2 (was %d)", wm.Size())
	}

	actual, expected := wm.GetFrom(0), "cat"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
	}

	actual, expected = wm.GetTo(0), "dog"
	if actual != expected {
		t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
	}

	actual, expected = wm.GetFrom(1), "horse"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
	}

	actual, expected = wm.GetTo(1), "zebra"
	if actual != expected {
		t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
	}
}

func TestCsvToFewColumns(t *testing.T) {
	csv := `zebra`
	_, err := parseCsvFile(&csv)

	if err == nil {
		t.Fatal("Error should be set for incorrect CSV file")
	}

	if !strings.Contains(err.Error(), "Unexpected row") {
		t.Errorf("Incorrect error message for (actual '%s')", err)
	}
}
