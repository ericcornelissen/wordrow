package wordmaps

import (
	"fmt"
	"testing"
)

func TestWordMapEmpty(t *testing.T) {
	var wm WordMap

	if wm.Size() != 0 {
		t.Errorf("The size of a new WordMap must be 0 (got %d)", wm.Size())
	}
}

func TestWordMapAddFileUnknownType(t *testing.T) {
	var wm WordMap

	content := "Hello world"
	format := ".bar"

	err := wm.AddFile(&content, format)
	if err == nil {
		t.Error("Expected error to be set but it was not")
	}
}

func TestWordMapAddFileKnownType(t *testing.T) {
	var wm WordMap

	from, to := "foo", "bar"
	content := fmt.Sprintf("%s,%s", from, to)
	format := ".csv"

	err := wm.AddFile(&content, format)
	if err != nil {
		t.Fatalf("Error should not be set for this test (got '%s')", err)
	}

	if wm.Size() != 1 {
		t.Fatalf("The WordMap size should to be 1 (got %d)", wm.Size())
	}

	actual := wm.GetFrom(0)
	if actual != from {
		t.Errorf("Unexpected from-value at index 0 (got '%s')", actual)
	}

	actual = wm.GetTo(0)
	if actual != to {
		t.Errorf("Unexpected to-value at index 0 (got '%s')", actual)
	}
}

func TestWordMapAddOne(t *testing.T) {
	var wm WordMap
	from, to := "cat", "dog"

	wm.AddOne(from, to)
	if wm.Size() != 1 {
		t.Fatalf("The size after WordMap.AddOne should be 1 (got %d)", wm.Size())
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

func TestWordMapEmptyValues(t *testing.T) {
	var wm WordMap

	t.Run("Empty from value", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("AddOne did not need recovery, but should have")
			}
		}()

		wm.AddOne("", "bar")
		t.Error("AddOne should have panicked but did not")
	})
	t.Run("Empty to value", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("AddOne did not need recovery, but should have")
			}
		}()

		wm.AddOne("foo", "")
		t.Error("AddOne should have panicked but did not")
	})
}

func TestWordMapAddFrom(t *testing.T) {
	var wmA, wmB WordMap

	wmA.AddOne("cat", "dog")
	if wmA.Size() != 1 || wmB.Size() != 0 {
		t.Fatal("The initial sizes of the WordMaps was incorrect for this test")
	}

	wmA.AddFrom(wmB)
	if wmA.Size() != 1 {
		t.Error("Adding an empty WordMap should not change that WordMap's size")
	}

	wmB.AddFrom(wmA)
	if wmB.Size() != 1 {
		t.Error("Adding a non-empty WordMap should increase that WordMap's size")
	}
}

func TestWordMapContains(t *testing.T) {
	var wm WordMap
	from, to := "cat", "dog"

	if wm.Contains(from) || wm.Contains(to) {
		t.Fatal("A new WordMap should not contain anything")
	}

	wm.AddOne(from, to)
	if !wm.Contains(from) {
		t.Error("The WordMap should contain a word added by AddOne")
	}

	if wm.Contains(to) {
		t.Error("The WordMap should NOT contain the 'to' word that was added")
	}
}

func TestWordMapGet(t *testing.T) {
	var wm WordMap
	from, to := "cat", "dog"

	wm.AddOne(from, to)
	if wm.Size() != 1 {
		t.Fatalf("The size of the WordMap must be 1 (got %d)", wm.Size())
	}

	outOfRangeIndex := wm.Size() + 1

	t.Run("GetFrom", func(t *testing.T) {
		actual := wm.GetFrom(0)
		if actual != from {
			t.Errorf("Incorrect from-value at index 0 (got '%s')", actual)
		}
	})
	t.Run("GetTo", func(t *testing.T) {
		actual := wm.GetTo(0)
		if actual != to {
			t.Errorf("Incorrect to-value at index 0 (got '%s')", actual)
		}
	})
	t.Run("GetFrom out of range", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("GetFrom did not need recovery, but should have")
			}
		}()

		wm.GetFrom(outOfRangeIndex)
		t.Error("GetFrom should have panicked but did not")
	})
	t.Run("GetTo out of range", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("GetTo did not need recovery, but should have")
			}
		}()

		wm.GetTo(outOfRangeIndex)
		t.Error("GetTo should have panicked but did not")
	})
}

func TestWordMapInvert(t *testing.T) {
	var wm WordMap
	from, to := "cat", "dog"

	wm.Invert()
	if wm.Size() != 0 {
		t.Error("Inverting an empty WordMap should not do anything")
	}

	wm.AddOne(from, to)
	if wm.Size() != 1 {
		t.Fatalf("The size should be 1 after AddOne (got %d)", wm.Size())
	}

	wm.Invert()
	if wm.Size() != 1 {
		t.Fatalf("The size should be the same after inverting (got %d)", wm.Size())
	}

	actual := wm.GetFrom(0)
	if actual != to {
		t.Errorf("Incorrect from-value at index 0 after inverting (got '%s')", actual)
	}

	actual = wm.GetTo(0)
	if actual != from {
		t.Errorf("Incorrect to-value at index 0 after inverting (got '%s')", actual)
	}
}

func TestWordMapIter(t *testing.T) {
	var wm WordMap
	from0, to0 := "cat", "dog"
	from1, to1 := "horse", "zebra"

	wm.AddOne(from0, to0)
	wm.AddOne(from1, to1)

	if wm.Size() != 2 {
		t.Fatalf("The size of the WordMap must be 2 (got %d)", wm.Size())
	}

	expectedFrom := []string{from0, from1}
	expectedTo := []string{to0, to1}

	i := 0
	for mapping := range wm.Iter() {
		if mapping.from != expectedFrom[i] {
			t.Errorf("Incorrect from-value at index %d (got '%s')", i, mapping.from)
		}
		if mapping.to != expectedTo[i] {
			t.Errorf("Incorrect to-value at index %d (got '%s')", i, mapping.to)
		}

		i++
	}
}

func TestWordMapString(t *testing.T) {
	var wm WordMap

	if wm.String() == "" {
		t.Error("A new WordMap should return a non-empty string for String()")
	}

	wm.AddOne("cat", "dog")
	if wm.String() == "" {
		t.Error("A non-empty WordMap should return a non-empty string for String()")
	}
}
