package wordmap

import "testing"


func TestWordMapEmpty(t *testing.T)  {
  var wordmap WordMap

  if wordmap.Size() != 0 {
    t.Errorf("The size of a new WordMap must be 0 (was: %d)", wordmap.Size())
  }
}

func TestWordMapAddOne(t *testing.T)  {
  var wordmap WordMap

  wordmap.AddOne("cat", "dog")
  if wordmap.Size() != 1 {
    t.Errorf("The size after WordMap.AddOne be 1 (was: %d)", wordmap.Size())
  }

  actual, expected := wordmap.GetFrom(0), "cat"
  if wordmap.Size() != 1 {
    t.Errorf("Incorrect from-value at index 0 (actual %s)", actual)
  }

  actual, expected = wordmap.GetTo(0), "dog"
  if actual != expected {
    t.Errorf("Incorrect to-value at index 0 (actual %s)", actual)
  }
}

func TestWordMapEmptyValues(t *testing.T) {
  var wordmap WordMap

  t.Run("Empty from value", func(t *testing.T) {
    defer func() {
      if r := recover(); r == nil {
        t.Error("AddOne did not need recovery, but should have")
      }
    }()

    wordmap.AddOne("", "bar")
    t.Error("AddOne should have panicked but did not")
  })
  t.Run("Empty to value", func(t *testing.T) {
    defer func() {
      if r := recover(); r == nil {
        t.Error("AddOne did not need recovery, but should have")
      }
    }()

    wordmap.AddOne("foo", "")
    t.Error("AddOne should have panicked but did not")
  })
}

func TestWordMapAddFrom(t *testing.T)  {
  var wordmapA, wordmapB WordMap

  wordmapA.AddOne("cat", "dog")
  if wordmapA.Size() != 1 || wordmapB.Size() != 0 {
    t.Error("The initial size of the WordMaps was incorrect for this test")
  }

  wordmapA.AddFrom(wordmapB)
  if wordmapA.Size() != 1 {
    t.Errorf("Adding an empty WordMap should not change a WordMaps size")
  }

  wordmapB.AddFrom(wordmapA)
  if wordmapB.Size() != 1 {
    t.Error("The size of WordMap B must be 1 after adding WordMap A")
  }
}

func TestWordMapContains(t *testing.T) {
  var wordmap WordMap

  if wordmap.Contains("a") {
    t.Error("A new WordMap should not contain anything")
  }

  wordmap.AddOne("cat", "dog")
  if !wordmap.Contains("cat") {
    t.Error("The WordMap should contain a word added by AddOne")
  }

  if wordmap.Contains("dog") {
    t.Error("The WordMap should NOT contain the 'to' word that was added")
  }
}

func TestWordMapGet(t *testing.T) {
  var wordmap WordMap

  wordmap.AddOne("cat", "dog")
  if wordmap.Size() != 1 {
    t.Fatalf("The size of the WordMap must be 1 (was: %d)", wordmap.Size())
  }

  outOfRangeIndex := wordmap.Size() + 1

  t.Run("GetFrom", func(t *testing.T) {
    if wordmap.GetFrom(0) != "cat" {
      t.Errorf("Incorrect from-value at index 0 (actual %s)", wordmap.GetFrom(0))
    }
  })
  t.Run("GetTo", func(t *testing.T) {
    if wordmap.GetTo(0) != "dog" {
      t.Errorf("Incorrect to-value at index 0 (actual %s)", wordmap.GetTo(0))
    }
  })
  t.Run("GetFrom out of range", func(t *testing.T) {
    defer func() {
      if r := recover(); r == nil {
        t.Error("GetFrom did not need recovery, but should have")
      }
    }()

    wordmap.GetFrom(outOfRangeIndex)
    t.Error("GetFrom should have panicked but did not")
  })
  t.Run("GetTo out of range", func(t *testing.T) {
    defer func() {
      if r := recover(); r == nil {
        t.Error("GetTo did not need recovery, but should have")
      }
    }()

    wordmap.GetTo(outOfRangeIndex)
    t.Error("GetTo should have panicked but did not")
  })
}

func TestWordMapInvert(t *testing.T) {
  var wordmap WordMap

  wordmap.Invert()
  if wordmap.Size() != 0 {
    t.Error("Invertping an empty wordmap should not do antyhing")
  }

  wordmap.AddOne("cat", "dog")
  if wordmap.Size() != 1 {
    t.Fatalf("The size of a WordMap should be 1 after AddOne fo this test (was %d)", wordmap.Size())
  }

  wordmap.Invert()
  if wordmap.Size() != 1 {
    t.Fatalf("The size of a WordMap should be the same after swapping (was %d)", wordmap.Size())
  }

  actual, expected := wordmap.GetFrom(0), "dog"
  if wordmap.Size() != 1 {
    t.Errorf("Incorrect from-value at index 0 after swap (actual %s)", actual)
  }

  actual, expected = wordmap.GetTo(0), "cat"
  if actual != expected {
    t.Errorf("Incorrect to-value at index 0 after swap (actual %s)", actual)
  }
}

func TestWordMapIter(t *testing.T) {
  var wordmap WordMap
  wordmap.AddOne("cat", "dog")
  wordmap.AddOne("horse", "zebra")

  if wordmap.Size() != 2 {
    t.Fatalf("The size of the WordMap must be 2 (was: %d)", wordmap.Size())
  }

  expectedFrom := []string{"cat", "horse"}
  expectedTo := []string{"dog", "zebra"}
  for i, mapping := range wordmap.Iter() {
    if mapping.from != expectedFrom[i] {
      t.Errorf("Incorrect from-value at index %d (%s != %s)", i, mapping.from, expectedFrom[i])
    }
    if mapping.to != expectedTo[i] {
      t.Errorf("Incorrect to-value at index %d (%s != %s)", i, mapping.to, expectedTo[i])
    }
  }
}

func TestWordMapString(t *testing.T) {
  var wordmap WordMap

  if wordmap.String() == "" {
    t.Error("A new WordMap should return a non-empty string for String()")
  }

  wordmap.AddOne("cat", "dog")
  if wordmap.String() == "" {
    t.Error("A non-empty WordMap should return a non-empty string for String()")
  }
}
