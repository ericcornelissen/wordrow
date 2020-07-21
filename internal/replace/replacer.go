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
func replaceOne(s, from, to string) string {
	var sb strings.Builder
	mapping := Mapping{from, to}

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

// All replaces substrings of `s` according to the mapping in `wordmap`.
func All(s string, wordmap wordmaps.WordMap) string {
	for from, to := range wordmap.Iter() {
		s = replaceOne(s, from, to)
	}

	return s
}
