package wordmap

import "strings"
import "testing"


func TestParseError(t *testing.T)  {
  msg := "Unexpected format"
  src := "tHiS dOeS nOt MaKe SeNsE"
  err := &parseError{msg, src}

  if err.Error() == "" {
    t.Fatal("The error should not give an empty string for non-empty message")
  }

  if !strings.Contains(err.Error(), msg) {
    t.Error("The message must be in the error string")
  }

  if !strings.Contains(err.Error(), src) {
    t.Error("The source string must be in the error string")
  }
}
