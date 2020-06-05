package wordmaps

import (
	"fmt"
	"strings"
	"testing"
)

func TestMarkDownTableOnly(t *testing.T) {
	from, to := "cat", "dog"
	markdown := fmt.Sprintf(`
		| from | to  |
		| ---- | --- |
		| %s   | %s  |
	`, from, to)

	wm, err := parseMarkDownFile(&markdown)
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

func TestMarkDownTextAndTable(t *testing.T) {
	from0, to0 := "cat", "dog"
	from1, to1 := "horse", "zebra"
	markdown := fmt.Sprintf(`
		# Translation table

		Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque quam
		mauris, sollicitudin et mauris quis, luctus bibendum risus. Vestibulum
		vitae ligula et ex semper ullamcorper at eu massa.

		| from | to  |
		| ---- | --- |
		| %s   | %s  |
		| %s   | %s  |

		Suspendisse ante ante, interdum id felis vel, posuere.
	`, from0, to0, from1, to1)

	wm, err := parseMarkDownFile(&markdown)
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

func TestMarkDownTwoTables(t *testing.T) {
	from0, to0 := "horse", "zebra"
	from1, to1 := "cat", "dog"
	markdown := fmt.Sprintf(`
		| from | to  |
		| ---- | --- |
		| %s   | %s  |

		| from | to  |
		| ---- | --- |
		| %s   | %s  |
	`, from0, to0, from1, to1)

	wm, err := parseMarkDownFile(&markdown)
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

func TestMarkDownEmptyColumnValues(t *testing.T) {
	t.Run("Empty from value", func(t *testing.T) {
		markdown := `
			| from | to  |
			| ---- | --- |
			|      | bar |
		`

		_, err := parseMarkDownFile(&markdown)

		if err == nil {
			t.Fatalf("Error should be set if the from value is empty")
		}
	})
	t.Run("Empty to value", func(t *testing.T) {
		markdown := `
			| from | to |
			| ---- | -- |
			| foo  |    |
		`

		_, err := parseMarkDownFile(&markdown)

		if err == nil {
			t.Fatalf("Error should be set if the to value is empty")
		}
	})
}

func TestMarkDownIncorrectHeader(t *testing.T) {
	markdown := `
		| foo |
		| --- | --- |
		| cat | dog |
	`

	_, err := parseMarkDownFile(&markdown)

	if err == nil {
		t.Fatal("Error should be set for incorrect table header")
	}

	if !strings.Contains(err.Error(), "Incorrect table header") {
		t.Errorf("Incorrect error message for (got '%s')", err)
	}
}

func TestMarkDownMissingDivider(t *testing.T) {
	markdown := `
		| foo | bar |
		| cat | dog |
	`

	_, err := parseMarkDownFile(&markdown)

	if err == nil {
		t.Fatal("Error should be set for missing table header divider")
	}

	if !strings.Contains(err.Error(), "Missing table header divider") {
		t.Errorf("Incorrect error message for (got '%s')", err)
	}
}

func TestMarkDownIncorrectDivider(t *testing.T) {
	markdown := `
		| foo | bar |
		| --- |
		| cat | dog |
	`

	_, err := parseMarkDownFile(&markdown)

	if err == nil {
		t.Fatal("Error should be set for incorrect header divider")
	}

	if !strings.Contains(err.Error(), "Missing table header divider") {
		t.Errorf("Incorrect error message for (got '%s')", err)
	}
}

func TestMarkDownMissingTableBody(t *testing.T) {
	markdown := `
		| foo | bar |
		| --- | --- |
	`

	_, err := parseMarkDownFile(&markdown)

	if err == nil {
		t.Fatal("Error should be set for missing table body")
	}

	if !strings.Contains(err.Error(), "Missing table body") {
		t.Errorf("Incorrect error message for (got '%s')", err)
	}
}

func TestMarkDownIncorrectTableRow(t *testing.T) {
	markdown := `
		| foo   | bar   |
		| ----- | ----- |
		| dog   | cat   |
		| hello | world | ! |
		| horse | zebra |
	`

	_, err := parseMarkDownFile(&markdown)

	if err == nil {
		t.Fatal("Error should be set for row with incorrect number of columns")
	}

	if !strings.Contains(err.Error(), "Unexpected table row format") {
		t.Errorf("Incorrect error message for (got '%s')", err)
	}
}

func TestMarkdownIncompleteTableEndOfFile(t *testing.T) {
	markdown := `
		# Foobar

		| foo | bar |
	`

	_, err := parseMarkDownFile(&markdown)

	if err == nil {
		t.Fatal("Error should be set for incomplete table at the end of the file")
	}
}
