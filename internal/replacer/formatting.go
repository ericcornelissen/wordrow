package replacer

import "regexp"
import "strings"
import "unicode"


// Regular Expression to match the words, and substrings between the words, of a
// phrase.
var phraseToWordsExpr = regexp.MustCompile(`([A-z]+)([^A-z]*)`)


// Check if a string starts an uppercase letter.
func startsWithCapital(s string) bool {
  firstChar := s[0]
  firstCharRune := rune(firstChar)
  return unicode.IsUpper(firstCharRune)
}

// Convert a string to sentence case. I.e. make the first letter in the string
// uppercase.
func toSentenceCase(s string) string {
  return strings.ToUpper(s[:1]) + s[1:]
}


// If the `from` string is all caps, it will return `to` as all caps as well.
// Otherwise, the `to` string is returned unchanged.
func maintainAllCaps(from, to string) string {
  if strings.ToUpper(from) == from {
    return strings.ToUpper(to)
  }

  return to
}

// If the `from` string starts with a capital letter, it will return the `to`
// string starting with a capital letter as well. Otherwise, the `to` string is
// returned unchanged.
//
// If the `from` string consists of multiple words, the capitalization will be
// maintained for every word in the string.
func maintainCapitalization(fromPhrase, toPhrase string) string {
  var sb strings.Builder

  fromWords := phraseToWordsExpr.FindAllStringSubmatch(fromPhrase, -1)
  toWords := phraseToWordsExpr.FindAllStringSubmatch(toPhrase, -1)

  shortestLen := len(fromWords)
  if len(toWords) < len(fromWords) {
    shortestLen = len(toWords)
  }

  for i := 0; i < shortestLen; i++ {
    fromWord, toWord := fromWords[i][1], toWords[i][1]
    toDivider := toWords[i][2]

    if startsWithCapital(fromWord) {
      toWord = toSentenceCase(toWord)
    }

    sb.WriteString(toWord)
    sb.WriteString(toDivider)
  }

  return sb.String()
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
