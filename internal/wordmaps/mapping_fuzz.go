// +build gofuzz

package wordmaps

import "strings"

func FuzzMapping(_data []byte) int {
	data := string(_data)
	tmp := strings.Split(data, "\n")

	substr := tmp[0]
	s := strings.Join(tmp[1:], "\n")

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
