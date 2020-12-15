package replace

import (
	"regexp"
	"unicode"

	"github.com/ericcornelissen/stringsx"
)

// A Regular Expression that matches newlines.
var newlineExpr = regexp.MustCompile(`\r|\n|\r\n`)

// Regular Expression to match the words, and substrings between the words, of a
// phrase.
var phraseToWordsExpr = regexp.MustCompile(`([A-z]+)([^A-z]*)`)

// A Regular Expression that matches groups of whitespace characters.
var whitespaceExpr = regexp.MustCompile(`(\s+)`)

// Check if a string starts an uppercase letter.
func startsWithCapital(s string) bool {
	firstChar := s[0]
	firstCharRune := rune(firstChar)
	return unicode.IsUpper(firstCharRune)
}

// Convert a string to sentence case. I.e. make the first letter in the string
// uppercase.
func toSentenceCase(s string) string {
	return stringsx.ToUpper(s[:1]) + s[1:]
}

// If the `from` string is all caps, it will return `to` as all caps as well.
// Otherwise, the `to` string is returned unchanged.
func maintainAllCaps(from, to string) string {
	if stringsx.ToUpper(from) == from {
		return stringsx.ToUpper(to)
	}

	return to
}

// Maintain the capitalization as it is found in `fromWords` in `toWords` for
// each word in `fromWords`. If len(toWords) > len(fromWords) the trailing words
// are omitted.
func maintainCapitalizationWordByWord(fromWords, toWords [][]string) string {
	var sb stringsx.Builder
	shortestLen := minInt(len(fromWords), len(toWords))
	for i := 0; i < shortestLen; i++ {
		fromWord, toWord, toDivider := fromWords[i][1], toWords[i][1], toWords[i][2]
		if startsWithCapital(fromWord) {
			toWord = toSentenceCase(toWord)
		}

		sb.WriteString(toWord + toDivider)
	}

	return sb.String()
}

// Get the trailing words in `toWords` compared to `fromWords`.
func getTrailingWords(fromWords, toWords [][]string) string {
	var sb stringsx.Builder

	shortestLen := minInt(len(fromWords), len(toWords))
	for i := shortestLen; i < len(toWords); i++ {
		sb.WriteString(toWords[i][0])
	}

	return sb.String()
}

// If the `from` string starts with a capital letter, it will return the `to`
// string starting with a capital letter as well. Otherwise, the `to` string is
// returned unchanged.
//
// If the `from` string consists of multiple words, the capitalization will be
// maintained for every word in the string.
func maintainCapitalization(fromPhrase, toPhrase string) string {
	fromWords := phraseToWordsExpr.FindAllStringSubmatch(fromPhrase, -1)
	toWords := phraseToWordsExpr.FindAllStringSubmatch(toPhrase, -1)

	replacement := maintainCapitalizationWordByWord(fromWords, toWords)
	trailingWords := getTrailingWords(fromWords, toWords)

	return replacement + trailingWords
}

// If the `from` phrase contains whitespace (spaces, tabs, newlines), it will
// return the `to` phrase with the same kinds of whitespace. Otherwise, the `to`
// string is returned unchanged.
func maintainWhitespace(from, to string) (newTo string, offset int) {
	fromWhitespace := whitespaceExpr.FindAllStringSubmatchIndex(from, -1)
	toWhitespace := whitespaceExpr.FindAllStringSubmatchIndex(to, -1)

	shortestLen := minInt(len(fromWhitespace), len(toWhitespace))
	for i := 0; i < shortestLen; i++ {
		fromMatch, toMatch := fromWhitespace[i], toWhitespace[i]
		fromStart, fromEnd := fromMatch[0], fromMatch[1]
		toStart, toEnd := toMatch[0], toMatch[1]

		// Replace the whitespace in the `to` phrase with the whitespace of the
		// `from` phrase.
		to = to[:toStart] + from[fromStart:fromEnd] + to[toEnd:]
	}

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

// changesFormattingOnly checks whether the from and to values are the same
// except for their formatting, e.g. different capitalization and whitespace.
func changesFormattingOnly(from, to string) bool {
	normalizedFrom := whitespaceExpr.ReplaceAllString(from, " ")
	normalizedTo := whitespaceExpr.ReplaceAllString(to, " ")
	return stringsx.ToLower(normalizedFrom) == stringsx.ToLower(normalizedTo)
}

// Format the `to` string based on the format of the `from` string.
//
// This function does the following:
//  - Maintain all caps.
//  - Maintain first letter capitalization.
//  - Maintain newlines, tabs, etc.
func maintainFormatting(from, to string) (newTo string, offset int) {
	if !changesFormattingOnly(from, to) {
		to = stringsx.ToLower(to)
		to = maintainAllCaps(from, to)
		to = maintainCapitalization(from, to)
	}

	to, offset = maintainWhitespace(from, to)
	return to, offset
}

// Format the `to` string based on the format of the `from` string.
//
// This function does the following:
//  - Maintain all caps.
//  - Maintain first letter capitalization.
//  - Maintain newlines, tabs, etc.
func maintainFormattingBytes(_from, _to []byte) (newTo []byte, offset int) {
	from, to := string(_from), string(_to)

	if !changesFormattingOnly(from, to) {
		to = stringsx.ToLower(to)
		to = maintainAllCaps(from, to)
		to = maintainCapitalization(from, to)
	}

	to, offset = maintainWhitespace(from, to)
	return []byte(to), offset
}
