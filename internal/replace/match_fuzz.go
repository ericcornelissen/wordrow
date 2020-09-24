// +build gofuzz

package replace

func Fuzz(data []byte) int {
	if len(data) < 2 {
		return -1
	}

	queryValue := string(data[:len(data)/2])
	searchString := string(data[len(data)/2:])

	i := 0
	for range matches(searchString, queryValue) {
		i++
	}

	if i == 0 {
		return -1
	}

	return 1
}
