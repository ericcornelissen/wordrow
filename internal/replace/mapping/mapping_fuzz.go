// +build gofuzz

package mapping

import (
	"strings"
	"unicode/utf8"
)

func Fuzz(_data []byte) int {
	data := string(_data)
	tmp := strings.Split(data, "\n")

	substr := tmp[0]
	s := strings.Join(tmp[1:], "\n")

	if !utf8.ValidString(substr) {
		return -1 // Ignore substrings that contain non-UTF8 characters for now
	}

	if substr == "" || s == "" {
		return -1
	}

	i := 0
	for _ = range getAllMatches(s, substr) {
		i++
	}

	if i == 0 {
		return -1
	}

	return 1
}
