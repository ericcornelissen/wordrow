package wordmaps

import "fmt"
import "strings"

// An error representing an issue when parsing a file. The error should be
// provided, first, an error message describing why parsing failed, and second,
// the string causing the error.
type parseError struct {
	msg string
	src string
}

// Get the error message for the error.
func (e *parseError) Error() string {
	return fmt.Sprintf("%s (in '%s')", e.msg, strings.TrimSpace(e.src))
}
