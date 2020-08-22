package wordmaps

import (
	"fmt"
	"testing"
)

func TestWordMapAddFileUnknownType(t *testing.T) {
	wm := make(StringMap, 1)

	content := "Hello world"
	format := ".bar"

	err := wm.AddFile(&content, format)
	if err == nil {
		t.Error("Expected error to be set but it was not")
	}
}

func TestWordMapAddFileKnownType(t *testing.T) {
	wm := make(StringMap, 1)

	from, to := "foo", "bar"
	content := fmt.Sprintf("%s,%s", from, to)
	format := ".csv"

	err := wm.AddFile(&content, format)
	if err != nil {
		t.Fatalf("Error should not be set for this test (got '%s')", err)
	}

	expected := make([][]string, 1)
	expected[0] = []string{from, to}
	checkWordMap(t, wm, expected)
}

func TestWordMapAddOne(t *testing.T) {
	wm := make(StringMap, 1)
	from, to := "cat", "dog"

	wm.addOne(from, to)

	expected := make([][]string, 1)
	expected[0] = []string{from, to}
	checkWordMap(t, wm, expected)
}

func TestWordMapAddMany(t *testing.T) {
	wm := make(StringMap, 1)
	from1, from2, to := "doge", "puppy", "dog"

	wm.addMany([]string{from1, from2}, to)

	expected := make([][]string, 2)
	expected[0] = []string{from1, to}
	expected[1] = []string{from2, to}
	checkWordMap(t, wm, expected)
}

func TestWordMapEmptyValues(t *testing.T) {
	wm := make(StringMap, 1)

	t.Run("Empty from value", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("AddOne did not need recovery, but should have")
			}
		}()

		wm.addOne("", "bar")
		t.Error("AddOne should have panicked but did not")
	})
	t.Run("Empty to value", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("AddOne did not need recovery, but should have")
			}
		}()

		wm.addOne("foo", "")
		t.Error("AddOne should have panicked but did not")
	})
}

func TestWordMapAddFrom(t *testing.T) {
	wmA := make(StringMap, 1)
	wmB := make(StringMap, 1)

	wmA.addOne("cat", "dog")
	if len(wmA) != 1 || len(wmB) != 0 {
		t.Fatal("The initial sizes of the WordMaps was incorrect for this test")
	}

	wmA.addFrom(wmB)
	if len(wmA) != 1 {
		t.Error("Adding an empty WordMap should not change that WordMap's size")
	}

	wmB.addFrom(wmA)
	if len(wmB) != 1 {
		t.Error("Adding a non-empty WordMap should increase that WordMap's size")
	}
}

func TestWordMapContains(t *testing.T) {
	wm := make(StringMap, 1)
	from, to := "cat", "dog"

	wm.addOne(from, to)
	if _, ok := wm[from]; !ok {
		t.Error("The WordMap should contain a word added by AddOne")
	}

	if _, ok := wm[to]; ok {
		t.Error("The WordMap should NOT contain the 'to' word that was added")
	}
}

func TestWordMapGet(t *testing.T) {
	wm := make(StringMap, 1)
	from, to := "cat", "dog"

	wm.addOne(from, to)

	expected := make([][]string, 1)
	expected[0] = []string{from, to}
	checkWordMap(t, wm, expected)
}

func TestWordMapInvert(t *testing.T) {
	wm := make(StringMap, 3)
	from1, to1 := "cat", "dog"
	from2, to2 := "dog", "cat"
	from3, to3 := "Hello", "world"

	wm.Invert()
	if len(wm) != 0 {
		t.Error("Inverting an empty WordMap should not do anything")
	}

	wm.addOne(from1, to1)
	wm.addOne(from2, to2)
	wm.addOne(from3, to3)
	if len(wm) != 3 {
		t.Fatalf("The size should be 1 after AddOne (got %d)", len(wm))
	}

	wm = wm.Invert()

	expected := make([][]string, 3)
	expected[0] = []string{to1, from1}
	expected[1] = []string{to2, from2}
	expected[2] = []string{to3, from3}
	checkWordMap(t, wm, expected)
}
