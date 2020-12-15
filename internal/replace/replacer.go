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
	"bytes"
	"unicode/utf8"

	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/logger"
)

// Mapping is a utility struct representing a single mapping `from` one string
// `to` another.
type mapping struct {
	from string
	to   string
}

// Mapping is a utility struct representing a single mapping `from` one string
// `to` another.
type byteMapping struct {
	from []byte
	to   []byte
}

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
func replaceOne(s []byte, m *mapping) []byte {
	var bb bytes.Buffer

	lastIndex := 0
	for match := range matches(s, m.from) {
		replacement := getReplacement(match, m.to)
		replacement, offset := maintainFormatting(string(match.full), replacement)

		bb.Write(s[lastIndex:maxInt(match.start, lastIndex)])
		bb.WriteString(replacement)
		lastIndex = match.end + offset
	}

	if lastIndex < len(s) {
		bb.Write(s[lastIndex:])
	}

	return bb.Bytes()
}

// Replace all instances of `from` by `to` defined b `m` in `s`, or return the
// original string if the mapping is invalid.
func safeReplaceOne(s []byte, m *mapping) []byte {
	if !stringsx.IsValidUTF8(m.from) {
		logger.Warningf("Invalid character in mapping '%s'", m.from)
		return s
	}

	cleanFrom := stringsx.TrimSpace(removeAffixNotation(m.from))
	cleanTo := stringsx.TrimSpace(removeAffixNotation(m.to))
	if stringsx.IsEmpty(cleanFrom) || stringsx.IsEmpty(cleanTo) {
		logger.Warningf("Invalid mapping value '%s,%s'", m.from, m.to)
		return s
	}

	return replaceOne(s, m)
}

// All replaces substrings of `s` according to the mapping defined by `m`.
func All(s []byte, m map[string]string) []byte {
	for from, to := range m {
		s = safeReplaceOne(s, &mapping{from: from, to: to})
	}

	return s
}

// V -- Bytes implementation -- V //

// Get the replacement string including prefix/suffix given the match `m`.
func getReplacementBytes(m *match, s []byte) (replacement []byte) {
	keepPrefix, keepSuffix := detectAffixBytes(s)

	if keepPrefix {
		replacement = append(replacement, m.prefix...)
	}
	replacement = append(replacement, s...)
	if keepSuffix {
		replacement = append(replacement, m.suffix...)
	}

	return replacement
}

// Replace all instances of `from` by `to` in `s`.
func replaceOneBytes(s []byte, m *byteMapping) []byte {
	var bb bytes.Buffer

	lastIndex := 0
	for match := range matchesBytes(s, m.from) {
		replacement := getReplacementBytes(match, m.to)
		replacement, offset := maintainFormattingBytes(match.full, replacement)

		bb.Write(s[lastIndex:maxInt(match.start, lastIndex)])
		bb.Write(replacement)
		lastIndex = match.end + offset
	}

	if lastIndex < len(s) {
		bb.Write(s[lastIndex:])
	}

	return bb.Bytes()
}

// Replace all instances of `from` by `to` defined b `m` in `s`, or return the
// original string if the mapping is invalid.
func safeReplaceOneBytes(s []byte, m *byteMapping) []byte {
	if !utf8.Valid(m.from) {
		logger.Warningf("Invalid character in mapping '%s'", m.from)
		return s
	}

	cleanFrom := bytes.TrimSpace(removeAffixNotationByte(m.from))
	cleanTo := bytes.TrimSpace(removeAffixNotationByte(m.to))
	if len(cleanFrom) == 0 || len(cleanTo) == 0 {
		logger.Warningf("Invalid mapping value '%s,%s'", m.from, m.to)
		return s
	}

	return replaceOneBytes(s, m)
}

// All replaces substrings of `s` according to the mapping defined by `m`.
func AllBytes(s []byte, m [][]byte) []byte {
	for i := 0; i < len(m); i += 2 {
		s = safeReplaceOneBytes(s, &byteMapping{from: m[i], to: m[i+1]})
	}

	return s
}

// All replaces substrings of `s` according to the mapping defined by `m`.
func AllBytes2(s []byte, x ByteMap) []byte {
	for i, from := range x.from {
		to := x.to[i]
		s = safeReplaceOneBytes(s, &byteMapping{from: from, to: to})
	}

	return s
}
