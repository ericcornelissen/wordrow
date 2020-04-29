package errors

import "errors"
import "fmt"


// New creates a new `error` with a certain error text.
func New(text string) error {
  return errors.New(text)
}

// Newf creates a new `error` with a formatted error text.
func Newf(text string, args ...interface{}) error {
  formattedText := fmt.Sprintf(text, args...)
  return New(formattedText)
}
