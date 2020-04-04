package dicts

import "strings"
import "testing"


func TestCsvOneRow(t *testing.T) {
  csv := `cat,dog`
  wordmap, err := parseCsvFile(&csv)

  if err != nil {
    t.Fatalf("Error should be nil for this test (Error: %s)", err)
  }

  if wordmap.Size() != 1 {
    t.Fatalf("The WordMap size should be 1 (was %d)", wordmap.Size())
  }

  actual, expected := wordmap.GetFrom(0), "cat"
  if actual != expected {
    t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
  }

  actual, expected = wordmap.GetTo(0), "dog"
  if actual != expected {
    t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
  }
}

func TestCsvMultipleRows(t *testing.T) {
  csv := `
    cat,dog
    horse,zebra
  `
  wordmap, err := parseCsvFile(&csv)

  if err != nil {
    t.Fatalf("Error should be nil for this test (Error: %s)", err)
  }

  if wordmap.Size() != 2 {
    t.Fatalf("The WordMap size should be 2 (was %d)", wordmap.Size())
  }

  actual, expected := wordmap.GetFrom(0), "cat"
  if actual != expected {
    t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
  }

  actual, expected = wordmap.GetTo(0), "dog"
  if actual != expected {
    t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
  }

  actual, expected = wordmap.GetFrom(1), "horse"
  if actual != expected {
    t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
  }

  actual, expected = wordmap.GetTo(1), "zebra"
  if actual != expected {
    t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
  }
}

func TestCsvEmptyToValue(t *testing.T) {
  csv := `cat,`
  wordmap, err := parseCsvFile(&csv)

  if err != nil {
    t.Fatalf("Error should be nil for this test (Error: %s)", err)
  }

  if wordmap.Size() != 1 {
    t.Fatalf("The WordMap size should be 1 (was %d)", wordmap.Size())
  }

  actual, expected := wordmap.GetFrom(0), "cat"
  if actual != expected {
    t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
  }

  actual, expected = wordmap.GetTo(0), ""
  if actual != expected {
    t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
  }
}

func TestCsvIgnoreEmptyLines(t *testing.T) {
  csv := `
    cat,dog

    horse,zebra
  `
  wordmap, err := parseCsvFile(&csv)

  if err != nil {
    t.Fatalf("Error should be nil for this test (Error: %s)", err)
  }

  if wordmap.Size() != 2 {
    t.Fatalf("The WordMap size should be 2 (was %d)", wordmap.Size())
  }

  actual, expected := wordmap.GetFrom(0), "cat"
  if actual != expected {
    t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
  }

  actual, expected = wordmap.GetTo(0), "dog"
  if actual != expected {
    t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
  }

  actual, expected = wordmap.GetFrom(1), "horse"
  if actual != expected {
    t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
  }

  actual, expected = wordmap.GetTo(1), "zebra"
  if actual != expected {
    t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
  }
}

func TestCsvIgnoresWhiteSpaceInRow(t *testing.T) {
  csv := `
    cat, dog
    horse  , zebra
  `
  wordmap, err := parseCsvFile(&csv)

  if err != nil {
    t.Fatalf("Error should be nil for this test (Error: %s)", err)
  }

  if wordmap.Size() != 2 {
    t.Fatalf("The WordMap size should be 2 (was %d)", wordmap.Size())
  }

  actual, expected := wordmap.GetFrom(0), "cat"
  if actual != expected {
    t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
  }

  actual, expected = wordmap.GetTo(0), "dog"
  if actual != expected {
    t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
  }

  actual, expected = wordmap.GetFrom(1), "horse"
  if actual != expected {
    t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
  }

  actual, expected = wordmap.GetTo(1), "zebra"
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

func TestCsvToManyColumns(t *testing.T) {
  csv := `cat,dog,horse`
  _, err := parseCsvFile(&csv)

  if err == nil {
    t.Fatal("Error should be set for incorrect CSV file")
  }

  if !strings.Contains(err.Error(), "Unexpected row") {
    t.Errorf("Incorrect error message for (actual '%s')", err)
  }
}
