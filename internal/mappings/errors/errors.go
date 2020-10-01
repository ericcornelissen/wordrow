// Package errors is a simple utilities package that provides custom error
// creation function for mappings parsers.
package errors

import "github.com/ericcornelissen/wordrow/internal/errors"

const (
	// IncorrectFormat is the error text for an incorrect format in a mapping.
	incorrectFormat = "Incorrect format (in '%s')"

	// MissingValue is the error text for a missing value in a mapping.
	missingValue = "Missing value (in '%s')"
)

// Newf calls errors.Newf.
func Newf(s string, args ...interface{}) error {
	return errors.Newf(s, args...)
}

// NewIncorrectFormat returns a new error for an incorrect format.
func NewIncorrectFormat(s string) error {
	return errors.Newf(incorrectFormat, s)
}

// NewMissingValue returns a new error for a missing value.
func NewMissingValue(s string) error {
	return errors.Newf(missingValue, s)
}
