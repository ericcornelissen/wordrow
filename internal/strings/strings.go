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

// IsEmpty returns true if the string is empty and false otherwise.
func IsEmpty(s string) bool {
	return s == ""
}

// Split runs strings.Split.
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// TrimSpace runs strings.TrimSpace.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// TrimSpaceAll trims leading and trailing spaces in all strings. Note that this
// function operates in place.
func TrimSpaceAll(v []string) {
	for i, s := range v {
		v[i] = strings.TrimSpace(s)
	}
}
