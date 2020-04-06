package errors

import "errors"
import "fmt"

func New(text string) error {
  return errors.New(text)
}

func Newf(text string, args ...interface{}) error {
  formattedText := fmt.Sprintf(text, args...)
  return New(formattedText)
}
