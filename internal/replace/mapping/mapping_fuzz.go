// +build gofuzz

package mapping

import "unicode/utf8"

func Fuzz(data []byte) int {
	if len(data) < 10 {
		return -1
	}

	mappingData := data[0 : len(data)/2]
	stringData := string(data[len(data)/2:])

	m := New(
		string(mappingData[0:len(mappingData)/2]),
		string(mappingData[len(mappingData)/2:]),
	)

	if !utf8.ValidString(m.from) {
		return -1 // Ignore substrings that contain non-UTF8 characters for now
	}

	i := 0
	for _ = range m.Match(stringData) {
		i++
	}

	if i == 0 {
		return -1
	}

	return 1
}
