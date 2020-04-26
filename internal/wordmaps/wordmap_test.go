package wordmaps

import "testing"

import "github.com/ericcornelissen/wordrow/internal/fs"


func TestWordMapEmpty(t *testing.T)  {
  var wm WordMap

  if wm.Size() != 0 {
    t.Errorf("The size of a new WordMap must be 0 (was: %d)", wm.Size())
  }
}

func TestWordMapAddFileUnknownType(t *testing.T)  {
  var wm WordMap
  file := fs.File{
    Content: "",
    Ext: "bar",
    Path: "foo.bar",
  }

  err := wm.AddFile(file)
  if err == nil {
    t.Error("Expected error to be set but was not")
  }
}

func TestWordMapAddFileKnownType(t *testing.T)  {
  var wm WordMap
  file := fs.File{
    Content: "foo,bar",
    Ext: "csv",
    Path: "foo.csv",
  }

  err := wm.AddFile(file)
  if err != nil {
    t.Fatalf("Error should not be set for this test")
  }

  if wm.Size() != 1 {
    t.Errorf("Expected wm size to be 1 (was %d)", wm.Size())
  }

  if wm.GetFrom(0) != "foo" {
    t.Errorf("Unexpected from value (was '%s')", wm.GetFrom(0))
  }

  if wm.GetTo(0) != "bar" {
    t.Errorf("Unexpected from value (was '%s')", wm.GetTo(0))
  }
}

func TestWordMapAddOne(t *testing.T)  {
  var wm WordMap

  wm.AddOne("cat", "dog")
  if wm.Size() != 1 {
    t.Errorf("The size after WordMap.AddOne be 1 (was: %d)", wm.Size())
  }

  actual, expected := wm.GetFrom(0), "cat"
  if wm.Size() != 1 {
    t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
  }

  actual, expected = wm.GetTo(0), "dog"
  if actual != expected {
    t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
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

func TestWordMapAddFrom(t *testing.T)  {
  var wmA, wmB WordMap

  wmA.AddOne("cat", "dog")
  if wmA.Size() != 1 || wmB.Size() != 0 {
    t.Error("The initial size of the WordMaps was incorrect for this test")
  }

  wmA.AddFrom(wmB)
  if wmA.Size() != 1 {
    t.Errorf("Adding an empty WordMap should not change a WordMaps size")
  }

  wmB.AddFrom(wmA)
  if wmB.Size() != 1 {
    t.Error("The size of WordMap B must be 1 after adding WordMap A")
  }
}

func TestWordMapContains(t *testing.T) {
  var wm WordMap

  if wm.Contains("a") {
    t.Error("A new WordMap should not contain anything")
  }

  wm.AddOne("cat", "dog")
  if !wm.Contains("cat") {
    t.Error("The WordMap should contain a word added by AddOne")
  }

  if wm.Contains("dog") {
    t.Error("The WordMap should NOT contain the 'to' word that was added")
  }
}

func TestWordMapGet(t *testing.T) {
  var wm WordMap

  wm.AddOne("cat", "dog")
  if wm.Size() != 1 {
    t.Fatalf("The size of the WordMap must be 1 (was: %d)", wm.Size())
  }

  outOfRangeIndex := wm.Size() + 1

  t.Run("GetFrom", func(t *testing.T) {
    if wm.GetFrom(0) != "cat" {
      t.Errorf("Incorrect from-value at index 0 (actual %s)", wm.GetFrom(0))
    }
  })
  t.Run("GetTo", func(t *testing.T) {
    if wm.GetTo(0) != "dog" {
      t.Errorf("Incorrect to-value at index 0 (actual %s)", wm.GetTo(0))
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

  wm.Invert()
  if wm.Size() != 0 {
    t.Error("Invertping an empty wm should not do antyhing")
  }

  wm.AddOne("cat", "dog")
  if wm.Size() != 1 {
    t.Fatalf("The size of a WordMap should be 1 after AddOne fo this test (was %d)", wm.Size())
  }

  wm.Invert()
  if wm.Size() != 1 {
    t.Fatalf("The size of a WordMap should be the same after swapping (was %d)", wm.Size())
  }

  actual, expected := wm.GetFrom(0), "dog"
  if wm.Size() != 1 {
    t.Errorf("Incorrect from-value at index 0 after swap (actual %s)", actual)
  }

  actual, expected = wm.GetTo(0), "cat"
  if actual != expected {
    t.Errorf("Incorrect to-value at index 0 after swap (actual %s)", actual)
  }
}

func TestWordMapIter(t *testing.T) {
  var wm WordMap
  wm.AddOne("cat", "dog")
  wm.AddOne("horse", "zebra")

  if wm.Size() != 2 {
    t.Fatalf("The size of the WordMap must be 2 (was: %d)", wm.Size())
  }

  expectedFrom := []string{"cat", "horse"}
  expectedTo := []string{"dog", "zebra"}

  i := 0
  for mapping := range wm.Iter() {
    if mapping.from != expectedFrom[i] {
      t.Errorf("Incorrect from-value at index %d (got %s)", i, mapping.from)
    }
    if mapping.to != expectedTo[i] {
      t.Errorf("Incorrect to-value at index %d (got %s)", i, mapping.to)
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
