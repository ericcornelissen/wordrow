package wordmap

import "fmt"
import "regexp"


// Check if a string ends with the suffix symbol.
func endsWithSuffixSymbol(s string) bool {
  return s[len(s) - 1:] == "-"
}

// Remove the prefix and suffix symbols from a string. If the symbols are not
// there, the original string is returned.
func removePrefixAndSuffixSymbols(s string) string {
  value := s

  if startsWithPrefixSymbol(s) {
    value = value[1:]
  }

  if endsWithSuffixSymbol(s) {
    value = value[:len(value) - 1]
  }

  return value
}

// Check if a string starts with the prefix symbol.
func startsWithPrefixSymbol(s string) bool {
  return s[0:1] == "-"
}


// The Match type represents a matching substring in a larger string of a
// Mapping, possibly including a prefix and/or suffix.
type Match struct {
  // The full match, i.e. the Word including prefix and suffix.
  Full string

  // The matched word as it appears in the original string.
  Word string

  // The replacement of the Word based on the mapping that created the Match.
  Replacement string

  // The prefix of the matched Word.
  Prefix string

  // The suffix of the matched Word.
  Suffix string

  // The starting index of the (full) match in the original string.
  Start int

  // The ending index of the (full) match in the original string.
  End int
}

// Find matches of some substring in a larger string, potentially with a prefix
// and/or suffix.
func getAllMatches(s, substr string) (chan Match) {
  ch := make(chan Match)

  go func() {
    defer close(ch)

    rawExpr := fmt.Sprintf(`(?i)([A-z0-9]*)(%s)([A-z0-9]*)`, substr)
    expr := regexp.MustCompile(rawExpr)

    for _, indices := range expr.FindAllStringSubmatchIndex(s, -1) {
      matchStart, matchEnd := indices[0], indices[1]
      prefixStart, prefixEnd := indices[2], indices[3]
      wordStart, wordEnd := indices[4], indices[5]
      suffixStart, suffixEnd := indices[6], indices[7]

      ch <- Match{
        Full: s[matchStart:matchEnd],
        Word: s[wordStart:wordEnd],
        Prefix: s[prefixStart:prefixEnd],
        Suffix: s[suffixStart:suffixEnd],
        Start: matchStart,
        End: matchEnd,
      }
    }
  }()

  return ch
}


// The Mapping type provides a guaranteed mapping from one string to another. As
// well as functionality to find matches in a string.
type Mapping struct {
  // The Mapping's "from" value. I.e. the value it wants to replace.
  from string

  // The Mapping's "to" value. I.e. he value it wants to repace "from" with.
  to string
}

// Get the replacement value, given a prefix and suffix. The return value will
// be the "to" value including, if necessary, the prefix and suffix.
//
// The prefix and suffix can always be an empty string.
func (mapping *Mapping) getReplacement(prefix, suffix string) string {
  replacement := mapping.GetTo()

  if mapping.keepPrefix() {
    replacement = prefix + replacement
  }

  if mapping.keepSuffix() {
    replacement = replacement + suffix
  }

  return replacement
}

// Check if a Match is valid for this Mapping. I.e. check if the Match has a
// prefix and/or suffix and if those are allowed by the Mapping's "from" value.
func (mapping *Mapping) isValid(match Match) bool {
  if !mapping.mayIncludePrefix() && match.Prefix != "" {
    return false
  }

  if !mapping.mayIncludeSuffix() && match.Suffix != "" {
    return false
  }

  return true
}

// Check if the mapping wants to keep the suffix in the replacement value.
func (mapping *Mapping) keepPrefix() bool {
  return startsWithPrefixSymbol(mapping.to)
}

// Check if the mapping wants to keep the suffix in the replacement value.
func (mapping *Mapping) keepSuffix() bool {
  return endsWithSuffixSymbol(mapping.to)
}

// Check if the mapping includes matches if they have a prefix.
func (mapping *Mapping) mayIncludePrefix() bool {
  return startsWithPrefixSymbol(mapping.from)
}

// Check if the mapping includes matches if they have a suffix.
func (mapping *Mapping) mayIncludeSuffix() bool {
  return endsWithSuffixSymbol(mapping.from)
}

// Get the Mapping's "from" value.
func (mapping *Mapping) GetFrom() string {
  return removePrefixAndSuffixSymbols(mapping.from)
}

// Get the Mapping's "to" value.
func (mapping *Mapping) GetTo() string {
  return removePrefixAndSuffixSymbols(mapping.to)
}

// Find matches of the "from" value of the Mapping in a string.
func (mapping *Mapping) Match(s string) (chan Match) {
  ch := make(chan Match)

  go func() {
    defer close(ch)

    matches := getAllMatches(s, mapping.GetFrom())
    for match := range matches {
      if mapping.isValid(match) {
        match.Replacement = mapping.getReplacement(match.Prefix, match.Suffix)
        ch <- match
      }
    }
  }()

  return ch
}

// Get the WordMap as a human readable string.
func (mapping *Mapping) String() string {
  return fmt.Sprintf("[%s -> %s]",
    mapping.from,
    mapping.to,
  )
}
