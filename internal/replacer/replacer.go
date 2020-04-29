package replacer

import "regexp"
import "strings"
import "unicode"

import "github.com/ericcornelissen/wordrow/internal/wordmaps"


// A Regular Expression that matches groups of whitespace characters.
var whitespaceExpr = regexp.MustCompile(`(\s+)`)

// A Regular Expression that matches newlines.
var newlineExpr = regexp.MustCompile(`\r|\n|\r\n`)


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

// If the `from` phrase contains more whitespace (spaces, tabs, newlines) and
// the whitespace in the same spot as where the `to` phrase ends is a newline,
// it will return the `to` string with a trailing newline.
func maintainTrailingNewlines(from, to string) (string, int) {
  offset := 0

  fromWhitespace := whitespaceExpr.FindAllStringSubmatchIndex(from, -1)
  toWhitespace := whitespaceExpr.FindAllStringSubmatchIndex(to, -1)

  if len(fromWhitespace) > len(toWhitespace) {
    lastMatchIndex := len(toWhitespace)
    lastFromMatch := fromWhitespace[lastMatchIndex]

    fromStart, fromEnd := lastFromMatch[0], lastFromMatch[1]
    trailingFromWhitespace := from[fromStart:fromEnd]

    if newlineExpr.MatchString(trailingFromWhitespace) {
      to += trailingFromWhitespace
      offset = 1
    }
  }

  return to, offset
}

// If the `from` phrase contains whitespace (spaces, tabs, newlines), it will
// return the `to` phrase with the same kinds of whitespace. Otherwise, the `to`
// string is returned unchanged.
func maintainWhitespace(from, to string) (string, int) {
  fromWhitespace := whitespaceExpr.FindAllStringSubmatchIndex(from, -1)
  toWhitespace := whitespaceExpr.FindAllStringSubmatchIndex(to, -1)

  shortest := len(fromWhitespace)
  if len(toWhitespace) < len(fromWhitespace) {
    shortest = len(toWhitespace)
  }

  for i := 0; i < shortest; i++ {
    fromMatch := fromWhitespace[i]
    fromStart, fromEnd := fromMatch[0], fromMatch[1]

    toMatch := toWhitespace[i]
    toStart, toEnd := toMatch[0], toMatch[1]

    to = to[:toStart] + from[fromStart:fromEnd] + to[toEnd:]
  }

  return maintainTrailingNewlines(from, to)
}

// Format the `to` string based on the format of the `from` string.
//
// This function does the following:
//  - Maintain all caps.
//  - Maintain first letter capitalization.
//  - Maintain newlines, tabs, etc.
func maintainFormatting(from, to string) (string, int) {
  to = maintainAllCaps(from, to)
  to = maintainCapitalization(from, to)
  to, offset := maintainWhitespace(from, to)
  return to, offset
}

// Replace all instances of `from` by `to` in `s`.
func replaceOne(s string, mapping wordmaps.Mapping) string {
  var sb strings.Builder

  lastIndex := 0
  for match := range mapping.Match(s) {
    replacement, offset := maintainFormatting(match.Full, match.Replacement)

    sb.WriteString(s[lastIndex:match.Start])
    sb.WriteString(replacement)
    lastIndex = match.End + offset
  }

  if lastIndex < len(s) {
    sb.WriteString(s[lastIndex:])
  }

  return sb.String()
}


// Replace substrings of `s` according to the mapping in `wordmap`.
func ReplaceAll(s string, wp wordmaps.WordMap) string {
  for mapping := range wp.Iter() {
    s = replaceOne(s, mapping)
  }

  return s
}
