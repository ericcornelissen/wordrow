// Package errors is a simple utilities package that streamlines error creation
// in Go. Primarily, it combines `errors.New` and `fmt.Errorf` into a single
// package.
package errors

import (
	"errors"
	"fmt"
)

// New creates a new `error` with a certain error text.
func New(text string) error {
	return errors.New(text)
}

// Newf creates a new `error` with a formatted error text.
func Newf(text string, args ...interface{}) error {
	formattedText := fmt.Sprintf(text, args...)
	return New(formattedText)
}
