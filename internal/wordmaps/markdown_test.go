package wordmaps

import (
	"strings"
	"testing"
)

func TestMarkDownTableOnly(t *testing.T) {
	markdown := `
		| foo | bar |
		| --- | --- |
		| cat | dog |
	`

	wm, err := parseMarkDownFile(&markdown)

	if err != nil {
		t.Fatalf("Error should be nil for this test (got '%s')", err)
	}

	if wm.Size() != 1 {
		t.Fatalf("The WordMap size should be 1 (got %d)", wm.Size())
	}

	actual, expected := wm.GetFrom(0), "cat"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 0 (got '%s')", actual)
	}

	actual, expected = wm.GetTo(0), "dog"
	if actual != expected {
		t.Errorf("Incorrect to-value at index 0 (got '%s')", actual)
	}
}

func TestMarkDownTextAndTable(t *testing.T) {
	markdown := `
		# Translation table

		Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque quam
		mauris, sollicitudin et mauris quis, luctus bibendum risus. Vestibulum
		vitae ligula et ex semper ullamcorper at eu massa.

		| foo   | bar   |
		| ----- | ----- |
		| cat   | dog   |
		| horse | zebra |

		Suspendisse ante ante, interdum id felis vel, posuere.
	`

	wm, err := parseMarkDownFile(&markdown)

	if err != nil {
		t.Fatalf("Error should be nil for this test (got '%s')", err)
	}

	if wm.Size() != 2 {
		t.Fatalf("The WordMap size should be 2 (got %d)", wm.Size())
	}

	actual, expected := wm.GetFrom(0), "cat"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 0 (got '%s')", actual)
	}

	actual, expected = wm.GetTo(0), "dog"
	if actual != expected {
		t.Errorf("Incorrect to-value at index 0 (got '%s')", actual)
	}

	actual, expected = wm.GetFrom(1), "horse"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 1 (got '%s')", actual)
	}

	actual, expected = wm.GetTo(1), "zebra"
	if actual != expected {
		t.Errorf("Incorrect to-value at index 1 (got '%s')", actual)
	}
}

func TestMarkDownTwoTables(t *testing.T) {
	markdown := `
		| foo   | bar   |
		| ----- | ----- |
		| zebra | horse |

		| foo | bar |
		| --- | --- |
		| dog | cat |
	`

	wm, err := parseMarkDownFile(&markdown)

	if err != nil {
		t.Fatalf("Error should be nil for this test (got '%s')", err)
	}

	if wm.Size() != 2 {
		t.Fatalf("The WordMap size should be 2 (got %d)", wm.Size())
	}

	actual, expected := wm.GetFrom(0), "zebra"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 0 (got '%s')", actual)
	}

	actual, expected = wm.GetTo(0), "horse"
	if actual != expected {
		t.Errorf("Incorrect to-value at index 0 (got '%s')", actual)
	}

	actual, expected = wm.GetFrom(1), "dog"
	if actual != expected {
		t.Errorf("Incorrect from-value at index 1 (got '%s')", actual)
	}

	actual, expected = wm.GetTo(1), "cat"
	if actual != expected {
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
