package replacer

import "fmt"
import "regexp"

import "github.com/ericcornelissen/wordrow/internal/wordmap"


// TODO
type Match struct {
  // The full match (i.e. including prefix and suffix).
  full string

  // The matched word (i.e. with original capitalization).
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

// Check if the Match has a prefix.
func (m *Match) HasPrefix() bool {
  return len(m.prefix) != 0
}

// Check if the Match has a Suffix.
func (m *Match) HasSuffix() bool {
  return len(m.suffix) != 0
}

// Check if th Match is allowed by a Mapping.
func (m *Match) IsAllowedBy(mapping wordmap.Mapping) bool {
  if m.HasPrefix() && mapping.From.PrefixAllowed == false {
    return false
  }

  if m.HasSuffix() && mapping.From.SuffixAllowed == false {
    return false
  }

  return true
}


// Find all matches, including those with a prefix or suffix, of a substring
// in a string of interest.
func findMatchesWithPrefixAndSuffix(s string, substr string) (chan Match) {
  ch := make(chan Match)

  rawExpr := fmt.Sprintf(`(?i)([A-z0-9]*)(%s)([A-z0-9]*)`, substr)
  expr := regexp.MustCompile(rawExpr)
  allIndices := expr.FindAllStringSubmatchIndex(s, -1)

  go func() {
    for _, indices := range allIndices {
      fullStart, fullEnd := indices[0], indices[1]
      prefixStart, prefixEnd := indices[2], indices[3]
      wordStart, wordEnd := indices[4], indices[5]
      suffixStart, suffixEnd := indices[6], indices[7]

      ch <- Match{
        full: s[fullStart:fullEnd],
        word: s[wordStart:wordEnd],

        prefix: s[prefixStart:prefixEnd],
        suffix: s[suffixStart:suffixEnd],

        start: fullStart,
        end: fullEnd,
      }
    }

    close(ch)
  }()

  return ch
}
