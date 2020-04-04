package replacer

import "regexp"
import "strings"
import "unicode"

import "github.com/ericcornelissen/wordrow/internal/dicts"


// The regular expression for a single letter.
var reLetter = regexp.MustCompile("[A-Za-z]")


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
  firstChar := []rune(from)[0]
  if unicode.IsUpper(firstChar) {
    return strings.ToUpper(to[:1]) + to[1:]
  } else {
    return to
  }
}

// Format the `to` string based on the format of the `from` string.
//
// This function does the following:
//  - Maintain all caps.
//  - Maintain first letter capitalization.
func smartReplace(from, to string) string {
  to = maintainAllCaps(from, to)
  to = maintainCapitalization(from, to)
  return to
}


// Check if the substring from `start` to `end` in `s` is a new word.
func isNewWordMatch(s string, start, end int) bool {
  if start > 0 && reLetter.MatchString(s[start - 1:start]) {
    return false
  } else {
    return true
  }
}

// Replace all instances of `from` by `to` in `s`.
func replaceOne(s string, from, to string) string {
  var sb strings.Builder

  re := regexp.MustCompile("(?i)" + from)
  indices := re.FindAllStringIndex(s, -1)

  prevIndex := 0
  for i := 0; i < len(indices); i++ {
    start, end := indices[i][0], indices[i][1]
    if !isNewWordMatch(s, start, end) {
      continue
    }

    matchedString := s[start:end]
    replacement := smartReplace(matchedString, to)

    sb.WriteString(s[prevIndex:start])
    sb.WriteString(replacement)
    prevIndex = end
  }

  sb.WriteString(s[prevIndex:])
  return sb.String()
}


// Replace substrings of `s` according to the mapping in `wordmap`.
func ReplaceAll(s string, wordmap dicts.WordMap) string {
  for _, mapping := range wordmap.Iter() {
    s = replaceOne(s, mapping.From, mapping.To)
  }

  return s
}
