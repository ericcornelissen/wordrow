package mapping

import (
	"fmt"
	"regexp"
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

// Find matches of some substring in a larger string, potentially with a prefix
// and/or suffix.
func getAllMatches(s, substr string) chan Match {
	ch := make(chan Match)

	go func() {
		defer close(ch)

		strToMatch := whitespaceExpr.ReplaceAllString(substr, `\s+`)

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
