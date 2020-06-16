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

// Fields runs strings.Fields.
func Fields(s string) []string {
	return strings.Fields(s)
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

// Repeat runs strings.Repeat.
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// ReplaceAll runs strings.ReplaceAll.
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// ToLower runs strings.ToLower.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper runs strings.ToUpper.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// TrimSpace runs strings.TrimSpace.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// TrimSuffix runs strings.TrimSuffix.
func TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}
