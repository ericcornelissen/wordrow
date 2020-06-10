// Package strings is an extensions of Go's standard strings package with some
// additional utility functions.
package strings

import "strings"

// Any returns true if at least one item in the list fulfills the condition and
// false otherwise.
func Any(v []string, condition func(string) bool) bool {
	for _, s := range v {
		if condition(s) {
			return true
		}
	}

	return false
}

// HasPrefix runs strings.HasPrefix.
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix runs strings.HasSuffix.
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// IsEmpty returns true if the string is empty and false otherwise.
func IsEmpty(s string) bool {
	return s == ""
}

// Map maps every string in a list using the specified function. Note that this
// function operates in place.
func Map(v []string, fn func(string) string) {
	for i, s := range v {
		v[i] = fn(s)
	}
}

// Split runs strings.Split.
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// TrimSpace runs strings.TrimSpace.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}
