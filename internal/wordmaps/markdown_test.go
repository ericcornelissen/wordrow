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

	expected := make([][]string, 1)
	expected[0] = []string{from, to}
	checkWordMap(t, wm, expected)
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

	expected := make([][]string, 2)
	expected[0] = []string{from0, to0}
	expected[1] = []string{from1, to1}
	checkWordMap(t, wm, expected)
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

	expected := make([][]string, 2)
	expected[0] = []string{from0, to0}
	expected[1] = []string{from1, to1}
	checkWordMap(t, wm, expected)
}

func TestMarkDownManyColumns(t *testing.T) {
	from01, from02, to0 := "cat", "doggy", "cat"
	from11, from12, to1 := "horse", "donkey", "zebra"
	markdown := fmt.Sprintf(`
		| from 1 | from 2 | to  |
		| ------ | ------ | --- |
		| %s     | %s     | %s  |
		| %s     | %s     | %s  |
	`, from01, from02, to0, from11, from12, to1)

	wm, err := parseMarkDownFile(&markdown)
	if err != nil {
		t.Fatalf("Error should be nil for this test (got '%s')", err)
	}

	expected := make([][]string, 4)
	expected[0] = []string{from01, to0}
	expected[1] = []string{from02, to0}
	expected[2] = []string{from11, to1}
	expected[3] = []string{from12, to1}
	checkWordMap(t, wm, expected)
}

func TestMarkDownEmptyColumnValues(t *testing.T) {
	t.Run("Empty from value", func(t *testing.T) {
		markdown := `
			| from | to  |
			| ---- | --- |
			| from | to  |
			|      | bar |
			| from | to  |
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
			| from | to |
			| foo  |    |
			| from | to |
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
