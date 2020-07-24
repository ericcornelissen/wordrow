package mapping

import (
	"fmt"
	"regexp"

	"github.com/ericcornelissen/wordrow/internal/logger"
	"github.com/ericcornelissen/wordrow/internal/strings"
)

// The Match type represents a matching substring in a larger string of a
// Mapping, possibly including a prefix and/or suffix.
type Match struct {
	// The full match, i.e. the Word including prefix and suffix.
	Full string

	// The matched word as it appears in the original string.
	Word string

	// The replacement of the Word based on the WordMap that created the Match.
	Replacement string

	// The prefix of the matched Word.
	Prefix string

	// The suffix of the matched Word.
	Suffix string

	// The starting index of the (Full) match in the original string.
	Start int

	// The ending index of the (Full) match in the original string.
	End int
}

// Given a query string to find Matches for, clean it so as to avoid any
// problems when using it to match against a target string.
func cleanStringToMatch(s string) string {
	s = strings.ReplaceAll(s, `\\`, `\`)
	s = strings.ReplaceAll(s, `\-`, `-`)
	s = regexp.QuoteMeta(s)
	return whitespaceExpr.ReplaceAllString(s, `\s+`)
}

// Find matches of some substring in a larger string, potentially with a prefix
// and/or suffix.
func getAllMatches(s, substr string) chan Match {
	ch := make(chan Match)

	go func() {
		defer close(ch)

		if !strings.IsValidUTF8(substr) {
			logger.Warningf("Invalid mapping value '%s'", substr)
			return
		}

		strToMatch := cleanStringToMatch(substr)
		rawExpr := fmt.Sprintf(`(?i)([A-z0-9]*)(%s)([A-z0-9]*)`, strToMatch)
		expr := regexp.MustCompile(rawExpr)

		for _, indices := range expr.FindAllStringSubmatchIndex(s, -1) {
			matchStart, matchEnd := indices[0], indices[1]
			prefixStart, prefixEnd := indices[2], indices[3]
			wordStart, wordEnd := indices[4], indices[5]
			suffixStart, suffixEnd := indices[6], indices[7]

			ch <- Match{
				Full:   s[matchStart:matchEnd],
				Word:   s[wordStart:wordEnd],
				Prefix: s[prefixStart:prefixEnd],
				Suffix: s[suffixStart:suffixEnd],
				Start:  matchStart,
				End:    matchEnd,
			}
		}
	}()

	return ch
}
