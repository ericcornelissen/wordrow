package replace

import (
	"fmt"
	"regexp"

	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/logger"
)

// The match type represents a matching substring in a larger string of a
// Mapping, possibly including a prefix and/or suffix.
type match struct {
	// The full match, i.e. the word including prefix and suffix.
	full []byte

	// The matched word as it appears in the original string.
	word []byte

	// The prefix of the matched word.
	prefix []byte

	// The suffix of the matched word.
	suffix []byte

	// The starting index of the (full) match in the original string.
	start int

	// The ending index of the (full) match in the original string.
	end int
}

// Get an empty channel of matches.
func emptyChannel() chan *match {
	ch := make(chan *match)
	close(ch)
	return ch
}

// Detect whether the string `s` contains the *wordrow* syntax for prefixes
// and/or suffixes.
func detectAffix(s string) (prefix, suffix bool) {
	if stringsx.HasPrefix(s, `-`) {
		prefix = true
	}

	if stringsx.HasSuffix(s, `-`) && !stringsx.HasSuffix(s, `\-`) {
		suffix = true
	}

	return prefix, suffix
}

// Get string `s` as a safe regular expression (escaping special characters) as
// well as removing any *wordrow* specific syntax.
func toSafeString(s string) (safeString string) {
	prefix, suffix := detectAffix(s)
	if prefix {
		s = s[1:]
	}
	if suffix {
		s = s[:len(s)-1]
	}

	safeString = stringsx.ReplaceAll(s, `\\`, `\`)
	safeString = stringsx.ReplaceAll(safeString, `\-`, `-`)
	safeString = regexp.QuoteMeta(safeString)
	return whitespaceExpr.ReplaceAllString(safeString, `\s+`)
}

// Check if a given match `m` is valid for the `query` string. I.e. if the match
// includes a prefix and/or suffix, is this allowed by the query string.
func isValidFor(m *match, query string) bool {
	withPrefix, withSuffix := detectAffix(query)
	if !withPrefix && len(m.prefix) != 0 {
		return false
	}

	if !withSuffix && len(m.suffix) != 0 {
		return false
	}

	return true
}

// Convert a slice of 8 indices (in the range of `s`) and turn it into a match
// struct.
func indicesToMatch(s []byte, indices []int) *match {
	matchStart, matchEnd := indices[0], indices[1]
	prefixStart, prefixEnd := indices[2], indices[3]
	wordStart, wordEnd := indices[4], indices[5]
	suffixStart, suffixEnd := indices[6], indices[7]

	return &match{
		full:   s[matchStart:matchEnd],
		word:   s[wordStart:wordEnd],
		prefix: s[prefixStart:prefixEnd],
		suffix: s[suffixStart:suffixEnd],
		start:  matchStart,
		end:    matchEnd,
	}
}

// Find all matches of a `query` string in a target string `s`.
//
// Note that non-UTF8 characters are not allowed, if any non-UTF characters are
// detected the function will panic.
func findMatches(s []byte, query string) chan *match {
	ch := make(chan *match)
	go func() {
		defer close(ch)

		safeQuery := toSafeString(query)
		rawExpr := fmt.Sprintf(`(?i)([A-z0-9]*)(%s)([A-z0-9]*)`, safeQuery)
		expr := regexp.MustCompile(rawExpr)
		for _, indices := range expr.FindAllSubmatchIndex(s, -1) {
			m := indicesToMatch(s, indices)
			if isValidFor(m, query) {
				ch <- m
			}
		}
	}()

	return ch
}

// Find all matches of a `query` string in a target string `s`.
//
// Note that non-UTF8 characters are not allowed, if any non-UTF characters are
// detected no matches will be returned.
func matches(s []byte, query string) chan *match {
	if !stringsx.IsValidUTF8(query) || stringsx.IsEmpty(query) {
		logger.Warningf("Invalid mapping value '%s'", query)
		return emptyChannel()
	}

	return findMatches(s, query)
}
