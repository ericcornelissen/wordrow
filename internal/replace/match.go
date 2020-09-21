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
	full string

	// The matched word as it appears in the original string.
	word string

	// The replacement of the word based on the Mapping that created the Match.
	replacement string

	// The prefix of the matched word.
	prefix string

	// The suffix of the matched word.
	suffix string

	// The starting index of the (full) match in the original string.
	start int

	// The ending index of the (full) match in the original string.
	end int
}

func detectAffix(s string) (prefix, suffix bool) {
	if stringsx.HasPrefix(s, `-`) {
		prefix = true
	}

	if stringsx.HasSuffix(s, `-`) && !stringsx.HasSuffix(s, `\-`) {
		suffix = true
	}

	return prefix, suffix
}

// Get `s` as a safe regular expression, as well as an indication of whether or
// not `s` allows prefixes and suffixes.
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

// Check if a given match `m` is valid for the query string `s`. I.e. if the
// match includes a prefix and/or suffix, is this supported by the query string.`
func isValidFor(m *match, s string) bool {
	withPrefix, withSuffix := detectAffix(s)

	if !withPrefix && m.prefix != "" {
		return false
	}

	if !withSuffix && m.suffix != "" {
		return false
	}

	return true
}

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

// Find all matches of `substr` in a target string `s`.
//
// This function will panic if any non-UTF8 characters are used.
func findAllMatches(s, substr string) chan match {
	ch := make(chan match)
	go func() {
		defer close(ch)

		rawExpr := fmt.Sprintf(`(?i)([A-z0-9]*)(%s)([A-z0-9]*)`, substr)
		expr := regexp.MustCompile(rawExpr)
		for _, indices := range expr.FindAllStringSubmatchIndex(s, -1) {
			matchstart, matchend := indices[0], indices[1]
			prefixstart, prefixend := indices[2], indices[3]
			wordstart, wordend := indices[4], indices[5]
			suffixstart, suffixend := indices[6], indices[7]

			ch <- match{
				full:   s[matchstart:matchend],
				word:   s[wordstart:wordend],
				prefix: s[prefixstart:prefixend],
				suffix: s[suffixstart:suffixend],
				start:  matchstart,
				end:    matchend,
			}
		}
	}()

	return ch
}

// Find all matches of the `from` string in a target string `s` with the correct
// replacement based on the `to` string.
//
// Note that non-UTF8 characters are not allowed, if any non-UTF* characters are
// detected no matches will be returned.
func matches(s, from, to string) chan match {
	ch := make(chan match)
	go func() {
		defer close(ch)

		if !stringsx.IsValidUTF8(from) || stringsx.IsEmpty(from) {
			logger.Warningf("Invalid mapping value '%s'", from)
			return
		}

		safeSubstr := toSafeString(from)
		for match := range findAllMatches(s, safeSubstr) {
			m := match
			if !isValidFor(&m, from) {
				continue
			}

			match.replacement = getReplacement(&m, to)
			ch <- match
		}
	}()

	return ch
}
