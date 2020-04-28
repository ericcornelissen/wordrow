package utils

import "testing"


func TestShortest(t *testing.T) {
  t.Run("First slice is the shortest", func(t *testing.T) {
    a := []string{"foo", "bar"}
    b := []string{"foo"}
    result := Shortest(a, b)

    if result != len(b) {
      t.Errorf("Unexpected shortest length (got %d)", result)
    }
  })
  t.Run("Second slice is the shortest", func(t *testing.T) {
    a := []string{"Dark", "Souls"}
    b := []string{"Praise", "the", "sun"}
    result := Shortest(a, b)

    if result != len(a) {
      t.Errorf("Unexpected shortest length (got %d)", result)
    }
  })
  t.Run("Slices are same length", func(t *testing.T) {
    a := []string{"The", "Spanish", "Inquisition"}
    b := []string{"Praise", "the", "sun"}

    if len(a) != len(b) {
      t.Fatal("Test slices must be the same length")
    }

    result := Shortest(a, b)

    if result != len(a) {
      t.Errorf("Unexpected shortest length (got %d)", result)
    }
  })
  t.Run("Not a slice", func(t *testing.T) {
    slice := []string{"The", "Spanish", "Inquisition"}

    Shortest(0, slice)
    Shortest(slice, 1)
    Shortest(0, 1)
  })
}
