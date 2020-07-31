/*
Package replace provides a function to smart replace strings in plaintext. To
this end it provides one function that accepts a string and a mapping and
returns the string with all words of the mapping replaced.

	var s string
	var m WordMap
	All(s, m)

The replacement will do some clever things to maintain the formatting of the
original text. Namely:

 • Maintain capitalization of words.
 • Maintain newline characters.
*/
package replace

import (
	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/replace/mapping"
)

// Replace all instances of `from` by `to` in `s`.
func replaceOne(s string, m mapping.Mapping) string {
	var sb stringsx.Builder

	lastIndex := 0
	for match := range m.Match(s) {
		replacement, offset := maintainFormatting(match.Full, match.Replacement)

		sb.WriteString(s[lastIndex:match.Start])
		sb.WriteString(replacement)
		lastIndex = match.End + offset
	}

	if lastIndex < len(s) {
		sb.WriteString(s[lastIndex:])
	}

	return sb.String()
}

// All replaces substrings of `s` according to the mapping defined by `m`.
func All(s string, m map[string]string) string {
	for from, to := range m {
		m := mapping.New(from, to)
		s = replaceOne(s, m)
	}

	return s
}
