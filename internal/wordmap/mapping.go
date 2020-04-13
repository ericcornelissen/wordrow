package wordmap

import "fmt"
import "regexp"


// Convert a string to a string indicating that a prefix and/or suffix is
// included.
func toPrefixSuffixString(s string, prefix, suffix bool) string {
  if prefix == true {
    s = "-" + s
  }

  if suffix == true {
    s = s + "-"
  }

  return s
}

// Check if a string starts with the prefix symbol. The first return value is
// a boolean indicating this; the second return value is the input string minus
// the first character iff it is the prefix symbol.
func startsWithPrefix(s string) (bool, string) {
  prefix, value := false, s
  if s[0:1] == "-" {
    prefix = true
    value = s[1:]
  }

  return prefix, value
}

// Check if a string ends with the suffix symbol. The first return value is
// a boolean indicating this; the second return value is the input string minus
// the last character iff it is the suffix symbol.
func endsWithSuffix(s string) (bool, string) {
  suffix, value := false, s
  if s[len(s) - 1:] == "-" {
    suffix = true
    value = s[:len(s) - 1]
  }

  return suffix, value
}


// The toValue type is used for the "from" key in Mapping.
type fromValue struct {
  // The value of interest. This should not include the prefix and/or suffix
  // symbols.
  Value string

  // Should the Value be considered if it appears with a prefix.
  IncludePrefix bool

  // Should the Value be considered if it appears with a suffix.
  IncludeSuffix bool
}

// Create a new fromValue from any string.
func newFromValue(rawValue string) fromValue {
  includePrefix, value := startsWithPrefix(rawValue)
  includeSuffix, value := endsWithSuffix(value)
  return fromValue{value, includePrefix, includeSuffix}
}

// Get the fromValue as human-readable string.
func (from *fromValue) String() string {
  return toPrefixSuffixString(from.Value, from.IncludePrefix, from.IncludeSuffix)
}


// The toValue type is used for the "to" key in Mapping.
type toValue struct {
  // The value of interest. This should not include the prefix and/or suffix
  // symbols.
  Value string

  // Should the Value keep the prefix of the fromValue.
  KeepPrefix bool

  // Should the Value keep the suffix of the fromValue.
  KeepSuffix bool
}

// Create a new ToValue from any string.
func newToValue(rawValue string) toValue {
  keepPrefix, value := startsWithPrefix(rawValue)
  keepSuffix, value := endsWithSuffix(value)
  return toValue{value, keepPrefix, keepSuffix}
}

// Get the ToValue as human-readable string.
func (to *toValue) String() string {
  return toPrefixSuffixString(to.Value, to.KeepPrefix, to.KeepSuffix)
}


// The Match type represents a matching substring in a larger string of a
// Mapping, possibly including a prefix and/or suffix.
type Match struct {
  // The full match (i.e. including prefix and suffix).
  Full string

  // The matched word as it appears in the original string.
  Word string

  // The prefix of the matched word.
  Prefix string

  // The suffix of the matched word.
  Suffix string

  // The starting index of the (full) match in the original string.
  Start int

  // The ending index of the (full) match in the original string.
  End int
}


// The Mapping type provides a guaranteed mapping from one string to another.
type Mapping struct {
  // TODO
  from fromValue

  // TODO
  to toValue
}

// Create a new Mapping Value
func newMapping(from, to string) Mapping {
  return Mapping{
    newFromValue(from),
    newToValue(to),
  }
}

// TODO
func (mapping *Mapping) isAllowedWith(prefix, suffix string) bool {
  if !mapping.from.IncludePrefix && prefix != "" {
    return false
  }

  if !mapping.from.IncludeSuffix && suffix != "" {
    return false
  }

  return true
}

// Get the replacement given a prefix and suffix. The return value will be the
// "to" value including, if necessary, the prefix and suffix.
func (mapping *Mapping) GetReplacement(prefix, suffix string) string {
  replacement := mapping.to.Value

  if mapping.to.KeepPrefix {
    replacement = prefix + replacement
  }

  if mapping.to.KeepSuffix {
    replacement = replacement + suffix
  }

  return replacement
}

// Find matches of the "from" value of the Mapping in a string.
func (mapping *Mapping) Match(s string) (chan Match) {
  ch := make(chan Match)

  go func() {
    rawExpr := fmt.Sprintf(`(?i)([A-z0-9]*)(%s)([A-z0-9]*)`, mapping.from.Value)
    expr := regexp.MustCompile(rawExpr)
    for _, indices := range expr.FindAllStringSubmatchIndex(s, -1) {
      fullStart, fullEnd := indices[0], indices[1]
      prefixStart, prefixEnd := indices[2], indices[3]
      wordStart, wordEnd := indices[4], indices[5]
      suffixStart, suffixEnd := indices[6], indices[7]

      match := Match{
        Full: s[fullStart:fullEnd],
        Word: s[wordStart:wordEnd],

        Prefix: s[prefixStart:prefixEnd],
        Suffix: s[suffixStart:suffixEnd],

        Start: fullStart,
        End: fullEnd,
      }

      if mapping.isAllowedWith(match.Prefix, match.Suffix) {
          ch <- match
      }
    }

    close(ch)
  }()

  return ch
}

// Get the WordMap as a human readable string.
func (mapping *Mapping) String() string {
  return fmt.Sprintf("[%s -> %s]",
    mapping.from.String(),
    mapping.to.String(),
  )
}
