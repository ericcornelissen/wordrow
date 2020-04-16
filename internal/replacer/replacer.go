package replacer

import "regexp"
import "strings"
import "unicode"

import "github.com/ericcornelissen/wordrow/internal/wordmap"


// The regular expression for a single letter.
var reLetter = regexp.MustCompile("[A-Za-z]")


// Check if a character (as byte) is an uppercase letter.
func isUpperChar(s byte) bool {
  firstChar := rune(s)
  return unicode.IsUpper(firstChar)
}

// Convert a string to sentence case. I.e. make the first letter in the string
// upper case.
func toSentenceCase(s string) string {
  return strings.ToUpper(s[:1]) + s[1:]
}


// If the `from` string is all caps, it will return `to` as all caps as well.
// Otherwise, the `to` string is returned unchanged.
func maintainAllCaps(from, to string) string {
  if strings.ToUpper(from) == from {
    return strings.ToUpper(to)
  } else {
    return to
  }
}

// If the `from` string starts with a capital letter, it will return the `to`
// starting with a capital letter as well. Otherwise, the `to` string is
// returned unchanged.
func maintainCapitalization(from, to string) string {
  if isUpperChar(from[0]) {
    return toSentenceCase(to)
  } else {
    return to
  }
}

// Format the `to` string based on the format of the `from` string.
//
// This function does the following:
//  - Maintain all caps.
//  - Maintain first letter capitalization.
func maintainFormatting(from, to string) string {
  to = maintainAllCaps(from, to)
  to = maintainCapitalization(from, to)
  return to
}

// Replace all instances of `from` by `to` in `s`.
func replaceOne(s string, mapping wordmap.Mapping) string {
  var sb strings.Builder

  lastIndex := 0
  for match := range mapping.Match(s) {
    replacement := maintainFormatting(match.Full, match.Replacement)

    sb.WriteString(s[lastIndex:match.Start])
    sb.WriteString(replacement)
    lastIndex = match.End
  }

  sb.WriteString(s[lastIndex:])
  return sb.String()
}


// Replace substrings of `s` according to the mapping in `wordmap`.
func ReplaceAll(s string, wordmap wordmap.WordMap) string {
  for mapping := range wordmap.Iter() {
    s = replaceOne(s, mapping)
  }

  return s
}
