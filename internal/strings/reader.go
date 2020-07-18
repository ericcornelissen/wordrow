package strings

import "strings"

// NewReader returns a new reader that outputs a string.
func NewReader(s string) *strings.Reader {
	return strings.NewReader(s)
}
