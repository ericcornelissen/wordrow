/*
Package replace provides a function to smart replace strings in plaintext. To
this end it provides one function that accepts a string and a mapping and
returns the string with all words of the mapping replaced.

	var s string
	var m map[string]string
	All(s, m)

The replacement will do some clever things to maintain the formatting of the
original text. Namely:

 • Maintain capitalization of words.
 • Maintain newline characters.
*/
package replace

import "github.com/ericcornelissen/stringsx"

// Replace all instances of `from` by `to` in `s`.
func replaceOne(s, from, to string) string {
	var sb stringsx.Builder

	lastIndex := 0
	for match := range matches(s, from, to) {
		replacement, offset := maintainFormatting(match.full, match.replacement)

		sb.WriteString(s[lastIndex:match.start])
		sb.WriteString(replacement)
		lastIndex = match.end + offset
	}

	if lastIndex < len(s) {
		sb.WriteString(s[lastIndex:])
	}

	return sb.String()
}

// All replaces substrings of `s` according to the mapping defined by `m`.
func All(s string, m map[string]string) string {
	for from, to := range m {
		s = replaceOne(s, from, to)
	}

	return s
}
