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

import "bytes"

// Get the replacement string including prefix/suffix given the match `m`.
func getReplacement(m *match, s string) string {
	keepPrefix, keepSuffix := detectAffix(s)

	replacement := s
	if keepPrefix {
		replacement = string(m.prefix) + replacement[1:]
	}

	if keepSuffix {
		replacement = replacement[:len(replacement)-1] + string(m.suffix)
	}

	return replacement
}

// Replace all instances of `from` by `to` in `s`.
func replaceOne(s []byte, from, to string) []byte {
	var bb bytes.Buffer

	lastIndex := 0
	for match := range matches(s, from) {
		replacement := getReplacement(match, to)
		replacement, offset := maintainFormatting(string(match.full), replacement)

		bb.Write(s[lastIndex:match.start])
		bb.WriteString(replacement)
		lastIndex = match.end + offset
	}

	if lastIndex < len(s) {
		bb.Write(s[lastIndex:])
	}

	return bb.Bytes()
}

// All replaces substrings of `s` according to the mapping defined by `m`.
func All(s []byte, m map[string]string) []byte {
	for from, to := range m {
		s = replaceOne(s, from, to)
	}

	return s
}
