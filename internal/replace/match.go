package replace

import (
	"fmt"
	"regexp"

	"github.com/ericcornelissen/stringsx"
)

// The match type represents a matching substring in a larger string of a
// Mapping, possibly including a prefix and/or suffix.
type match struct {
	// The full match, i.e. the word including prefix and suffix.
	full string

	// The matched word as it appears in the original string.
	word string

	// The prefix of the matched word.
	prefix string

	// The suffix of the matched word.
	suffix string

	// The starting index of the (full) match in the original string.
	start int

	// The ending index of the (full) match in the original string.
	end int
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

// Remove from the string `s` the *wordrow* syntax for prefixes and/or suffixes.
func removeAffixNotation(s string) string {
	prefix, suffix := detectAffix(s)
	if prefix && !stringsx.IsEmpty(s) {
		s = s[1:]
	}
	if suffix && !stringsx.IsEmpty(s) {
		s = s[:len(s)-1]
	}

	return s
}

// Get string `s` as a safe regular expression (escaping special characters) as
// well as removing any *wordrow* specific syntax.
func toSafeString(s string) (safeString string) {
	safeString = removeAffixNotation(s)
	safeString = stringsx.ReplaceAll(safeString, `\\`, `\`)
	safeString = stringsx.ReplaceAll(safeString, `\-`, `-`)
	safeString = regexp.QuoteMeta(safeString)
	return whitespaceExpr.ReplaceAllString(safeString, `\s+`)
}

// Check if a given match `m` is valid for the `query` string. I.e. if the match
// includes a prefix and/or suffix, is this allowed by the query string.
func isValidFor(m *match, query string) bool {
	withPrefix, withSuffix := detectAffix(query)
	if !withPrefix && m.prefix != "" {
		return false
	}

	if !withSuffix && m.suffix != "" {
		return false
	}

	return true
}

// Convert a slice of 8 indices (in the range of `s`) and turn it into a match
// struct.
func indicesToMatch(s string, indices []int) *match {
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
func matches(s, query string) chan *match {
	ch := make(chan *match)
	go func() {
		defer close(ch)

		safeQuery := toSafeString(query)
		rawExpr := fmt.Sprintf(`(?i)([A-z0-9]*)(%s)([A-z0-9]*)`, safeQuery)
		expr := regexp.MustCompile(rawExpr)
		for _, indices := range expr.FindAllStringSubmatchIndex(s, -1) {
			if m := indicesToMatch(s, indices); isValidFor(m, query) {
				ch <- m
			}
		}
	}()

	return ch
}
