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

// Get `s` as a safe regular expression, as well as an indication of whether or
// not `s` allows prefixes and suffixes.
func getRawExpr(s string) (pattern string, prefix, suffix bool) {
	escapedHyphen := `\\-`

	withPrefix, withSuffix := false, false
	pattern = regexp.QuoteMeta(stringsx.ReplaceAll(s, `\\`, `\`))

	if stringsx.HasPrefix(pattern, escapedHyphen) {
		pattern = pattern[2:]
	} else if stringsx.HasPrefix(pattern, `-`) {
		withPrefix = true
		pattern = pattern[1:]
	}

	if stringsx.HasSuffix(pattern, escapedHyphen) {
		pattern = pattern[:len(pattern)-3] + "-"
	} else if stringsx.HasSuffix(pattern, `-`) {
		withSuffix = true
		pattern = pattern[:len(pattern)-1]
	}

	return whitespaceExpr.ReplaceAllString(pattern, `\s+`), withPrefix, withSuffix
}

// Find all matches of `substr` in a target string `s`.
//
// This function will panic if any non-UTF8 characters are used.
func allMatches(s, expr string) chan match {
	ch := make(chan match)
	go func() {
		defer close(ch)

		rawExpr := fmt.Sprintf(`(?i)([A-z0-9]*)(%s)([A-z0-9]*)`, expr)
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

// Find all matches of `substr` in a target string `s` excluding those instances
// where `from` has a prefix/suffix in `s` when this is not allowed.
//
// This function will panic if any non-UTF8 characters are used.
func findMatches(s, substr string) chan match {
	ch := make(chan match)
	go func() {
		defer close(ch)

		rawExpr, withPrefix, withSuffix := getRawExpr(substr)
		for match := range allMatches(s, rawExpr) {
			if !withPrefix && match.prefix != "" {
				continue
			}

			if !withSuffix && match.suffix != "" {
				continue
			}

			ch <- match
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

		for match := range findMatches(s, from) {
			replacement := to

			_, keepPrefix, keepSuffix := getRawExpr(to)
			if keepPrefix {
				replacement = match.prefix + replacement[1:]
			}
			if keepSuffix {
				replacement = replacement[:len(replacement)-1] + match.suffix
			}

			match.replacement = replacement
			ch <- match
		}
	}()

	return ch
}
