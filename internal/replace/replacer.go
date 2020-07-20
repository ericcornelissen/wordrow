/*
Package replace provides a function to smart replace strings in plaintext. To
this end it provides one function that accepts a string and a mapping and
returns the string with all words of the mapping replaced.

	var s string
	var m WordMap
	ReplaceAll(s, m)

The replacement will do some clever things to maintain the formatting of the
original text. Namely:

 • Maintain capitalization of words.
 • Maintain newline characters.
*/
package replace

import (
	"github.com/ericcornelissen/wordrow/internal/strings"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

// Replace all instances of `from` by `to` in `s`.
func replaceOne(s string, mapping wordmaps.Mapping) string {
	var sb strings.Builder

	lastIndex := 0
	for match := range mapping.Match(s) {
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

// All replaces substrings of `s` according to the mapping in `wm`.
func All(s string, wp wordmaps.WordMap) string {
	for mapping := range wp.Iter() {
		s = replaceOne(s, mapping)
	}

	return s
}
