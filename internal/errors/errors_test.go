package errors

import "fmt"
import "testing"


func ExampleNew() {
  err := New("foobar")
  fmt.Print(err)
  // Output: foobar
}

func ExampleNewf() {
  err := Newf("foo%s", "bar")
  fmt.Print(err)
  // Output: foobar
}


func TestNew(t *testing.T) {
  err := New("foobar")

  if err == nil {
    t.Error("Error should not be nil")
  }

  if err.Error() != "foobar" {
    t.Error("not good")
  }
}

func TestNewf(t *testing.T) {
  err := Newf("foo%s", "bar")

  if err == nil {
    t.Error("Error should not be nil")
  }

  if err.Error() != "foobar" {
    t.Error("not good")
  }
}
