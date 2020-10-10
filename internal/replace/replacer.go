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

import (
	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/logger"
)

// Get the replacement string including prefix/suffix given the match `m`.
func getReplacement(m *match, s string) string {
	keepPrefix, keepSuffix := detectAffix(s)

	replacement := s
	if keepPrefix {
		replacement = m.prefix + replacement[1:]
	}

	if keepSuffix {
		replacement = replacement[:len(replacement)-1] + m.suffix
	}

	return replacement
}

// Replace all instances of `from` by `to` in `s`.
func replaceOne(s, from, to string) string {
	var sb stringsx.Builder

	lastIndex := 0
	for match := range matches(s, from) {
		replacement := getReplacement(match, to)
		replacement, offset := maintainFormatting(match.full, replacement)

		sb.WriteString(s[lastIndex:maxInt(match.start, lastIndex)])
		sb.WriteString(replacement)
		lastIndex = match.end + offset
	}

	if lastIndex < len(s) {
		sb.WriteString(s[lastIndex:])
	}

	return sb.String()
}

// Replace all instances of `from` by `to` in `s`, or return the original string
// if the mapping is invalid.
func safeReplaceOne(s, from, to string) string {
	if !stringsx.IsValidUTF8(from) {
		logger.Warningf("Invalid character in mapping '%s'", from)
		return s
	}

	cleanFrom := stringsx.TrimSpace(removeAffixNotation(from))
	cleanTo := stringsx.TrimSpace(removeAffixNotation(to))
	if stringsx.IsEmpty(cleanFrom) || stringsx.IsEmpty(cleanTo) {
		logger.Warningf("Invalid mapping value '%s,%s'", from, to)
		return s
	}

	return replaceOne(s, from, to)
}

// All replaces substrings of `s` according to the mapping defined by `m`.
func All(s string, m map[string]string) string {
	for from, to := range m {
		s = safeReplaceOne(s, from, to)
	}

	return s
}
