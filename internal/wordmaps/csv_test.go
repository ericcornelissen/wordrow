package wordmaps

import (
	"fmt"
	"strings"
	"testing"
)

func TestCsvOneRow(t *testing.T) {
	from, to := "cat", "dog"
	csv := fmt.Sprintf("%s,%s", from, to)

	wm, err := parseCsvFile(&csv)
	if err != nil {
		t.Fatalf("Error should be nil for this test (got '%s')", err)
	}

	if wm.Size() != 1 {
		t.Fatalf("The WordMap size should be 1 (got %d)", wm.Size())
	}

	actual := wm.GetFrom(0)
	if actual != from {
		t.Errorf("Incorrect from-value at index 0 (got '%s')", actual)
	}

	actual = wm.GetTo(0)
	if actual != to {
		t.Errorf("Incorrect to-value at index 0 (got '%s')", actual)
	}
}

func TestCsvMultipleRows(t *testing.T) {
	from0, to0 := "cat", "dog"
	from1, to1 := "horse", "zebra"
	csv := fmt.Sprintf(`
		%s,%s
		%s,%s
	`, from0, to0, from1, to1)

	wm, err := parseCsvFile(&csv)
	if err != nil {
		t.Fatalf("Error should be nil for this test (got '%s')", err)
	}

	if wm.Size() != 2 {
		t.Fatalf("The WordMap size should be 2 (got %d)", wm.Size())
	}

	actual := wm.GetFrom(0)
	if actual != from0 {
		t.Errorf("Incorrect from-value at index 0 (got '%s')", actual)
	}

	actual = wm.GetTo(0)
	if actual != to0 {
		t.Errorf("Incorrect to-value at index 0 (got '%s')", actual)
	}

	actual = wm.GetFrom(1)
	if actual != from1 {
		t.Errorf("Incorrect from-value at index 1 (got '%s')", actual)
	}

	actual = wm.GetTo(1)
	if actual != to1 {
		t.Errorf("Incorrect to-value at index 1 (got '%s')", actual)
	}
}

func TestCsvEmptyColumnValues(t *testing.T) {
	t.Run("Empty from value", func(t *testing.T) {
		csv := `,bar`

		_, err := parseCsvFile(&csv)

		if err == nil {
			t.Fatalf("Error should be set if the from value is empty")
		}
	})
	t.Run("Empty to value", func(t *testing.T) {
		csv := `foo,`

		_, err := parseCsvFile(&csv)

		if err == nil {
			t.Fatalf("Error should be set if the to value is empty")
		}
	})
}

func TestCsvIgnoreEmptyLines(t *testing.T) {
	from0, to0 := "cat", "dog"
	from1, to1 := "horse", "zebra"
	csv := fmt.Sprintf(`
		%s,%s

		%s,%s
	`, from0, to0, from1, to1)

	wm, err := parseCsvFile(&csv)
	if err != nil {
		t.Fatalf("Error should be nil for this test (got '%s')", err)
	}

	if wm.Size() != 2 {
		t.Fatalf("The WordMap size should be 2 (got %d)", wm.Size())
	}

	actual := wm.GetFrom(0)
	if actual != from0 {
		t.Errorf("Incorrect from-value at index 0 (got '%s')", actual)
	}

	actual = wm.GetTo(0)
	if actual != to0 {
		t.Errorf("Incorrect to-value at index 0 (got '%s')", actual)
	}

	actual = wm.GetFrom(1)
	if actual != from1 {
		t.Errorf("Incorrect from-value at index 1 (got '%s')", actual)
	}

	actual = wm.GetTo(1)
	if actual != to1 {
		t.Errorf("Incorrect to-value at index 1 (got '%s')", actual)
	}
}

func TestCsvIgnoresWhitespaceInRow(t *testing.T) {
	from0, to0 := "cat", "dog"
	from1, to1 := "horse", "zebra"
	csv := fmt.Sprintf(`
		%s, %s

		%s  , %s
	`, from0, to0, from1, to1)

	wm, err := parseCsvFile(&csv)
	if err != nil {
		t.Fatalf("Error should be nil for this test (got '%s')", err)
	}

	if wm.Size() != 2 {
		t.Fatalf("The WordMap size should be 2 (got %d)", wm.Size())
	}

	actual := wm.GetFrom(0)
	if actual != from0 {
		t.Errorf("Incorrect from-value at index 0 (got '%s')", actual)
	}

	actual = wm.GetTo(0)
	if actual != to0 {
		t.Errorf("Incorrect to-value at index 0 (got '%s')", actual)
	}

	actual = wm.GetFrom(1)
	if actual != from1 {
		t.Errorf("Incorrect from-value at index 1 (got '%s')", actual)
	}

	actual = wm.GetTo(1)
	if actual != to1 {
		t.Errorf("Incorrect to-value at index 1 (got '%s')", actual)
	}
}

func TestCsvToFewColumns(t *testing.T) {
	csv := `zebra`
	_, err := parseCsvFile(&csv)

	if err == nil {
		t.Fatal("Error should be set for incorrect CSV file")
	}

	if !strings.Contains(err.Error(), "Unexpected row") {
		t.Errorf("Incorrect error message for (got '%s')", err)
	}
}

func TestCsvToManyColumns(t *testing.T) {
	csv := `cat,dog,horse`
	_, err := parseCsvFile(&csv)

	if err == nil {
		t.Fatal("Error should be set for incorrect CSV file")
	}

	if !strings.Contains(err.Error(), "Unexpected row") {
		t.Errorf("Incorrect error message for (got '%s')", err)
	}
}
