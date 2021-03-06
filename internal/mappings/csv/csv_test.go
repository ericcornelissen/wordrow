package csv

import (
	"fmt"
	"testing"

	"github.com/ericcornelissen/stringsx"
	. "github.com/ericcornelissen/wordrow/internal/mappings/testing"
)

func TestCsvOneRow(t *testing.T) {
	from, to := "cat", "dog"
	csv := fmt.Sprintf("%s,%s", from, to)

	reader := NewTestReader(&csv)
	mapping, err := Parse(reader)
	if err != nil {
		t.Fatalf("Error should be nil for this test (got '%s')", err)
	}

	expected := make([][]string, 1)
	expected[0] = []string{from, to}
	CheckMapping(t, mapping, expected)
}

func TestCsvMultipleRows(t *testing.T) {
	from0, to0 := "cat", "dog"
	from1, to1 := "horse", "zebra"
	csv := fmt.Sprintf(`
		%s,%s
		%s,%s
	`, from0, to0, from1, to1)

	reader := NewTestReader(&csv)
	mapping, err := Parse(reader)
	if err != nil {
		t.Fatalf("Error should be nil for this test (got '%s')", err)
	}

	expected := make([][]string, 2)
	expected[0] = []string{from0, to0}
	expected[1] = []string{from1, to1}
	CheckMapping(t, mapping, expected)
}

func TestCsvManyColumns(t *testing.T) {
	from1, from2, to := "cat", "dog", "horse"
	csv := fmt.Sprintf("%s,%s,%s", from1, from2, to)
	reader := NewTestReader(&csv)
	mapping, err := Parse(reader)

	if err != nil {
		t.Fatalf("Error should be nil for this test (Error: %s)", err)
	}

	expected := make([][]string, 2)
	expected[0] = []string{from1, to}
	expected[1] = []string{from2, to}
	CheckMapping(t, mapping, expected)
}

func TestCsvEmptyColumnValues(t *testing.T) {
	t.Run("Empty from value", func(t *testing.T) {
		csv := `,bar`

		reader := NewTestReader(&csv)
		_, err := Parse(reader)

		if err == nil {
			t.Fatalf("Error should be set if the from value is empty")
		}
	})
	t.Run("Empty to value", func(t *testing.T) {
		csv := `foo,`

		reader := NewTestReader(&csv)
		_, err := Parse(reader)

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

	reader := NewTestReader(&csv)
	mapping, err := Parse(reader)
	if err != nil {
		t.Fatalf("Error should be nil for this test (got '%s')", err)
	}

	expected := make([][]string, 2)
	expected[0] = []string{from0, to0}
	expected[1] = []string{from1, to1}
	CheckMapping(t, mapping, expected)
}

func TestCsvIgnoresWhitespaceInRow(t *testing.T) {
	from0, to0 := "cat", "dog"
	from1, to1 := "horse", "zebra"
	csv := fmt.Sprintf(`
		%s, %s

		%s  , %s
	`, from0, to0, from1, to1)

	reader := NewTestReader(&csv)
	mapping, err := Parse(reader)
	if err != nil {
		t.Fatalf("Error should be nil for this test (got '%s')", err)
	}

	expected := make([][]string, 2)
	expected[0] = []string{from0, to0}
	expected[1] = []string{from1, to1}
	CheckMapping(t, mapping, expected)
}

func TestCsvToFewColumns(t *testing.T) {
	csv := `zebra`
	reader := NewTestReader(&csv)
	_, err := Parse(reader)

	if err == nil {
		t.Fatal("Error should be set for incorrect CSV file")
	}

	if !stringsx.Contains(err.Error(), "Incorrect format") {
		t.Errorf("Incorrect error message for (got '%s')", err)
	}
}
