package replacer

import "regexp"
import "strings"
import "unicode"

import "github.com/ericcornelissen/wordrow/internal/utils"


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
//
// If the `from` string consists of multiple words, the capitalization will be
// maintained for all words in the string.
func maintainCapitalization(fromPhrase, toPhrase string) string {
  var sb strings.Builder

  re := regexp.MustCompile(`([A-z]+)([^A-z]*)`)
  fromWords := re.FindAllStringSubmatch(fromPhrase, -1)
  toWords := re.FindAllStringSubmatch(toPhrase, -1)

  for i := 0; i < utils.Shortest(fromWords, toWords); i++ {
    fromWord, toWord := fromWords[i][1], toWords[i][1]
    toDivider := toWords[i][2]

    if isUpperChar(fromWord[0]) {
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
