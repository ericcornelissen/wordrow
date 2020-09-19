// +build gofuzz

package replace

func Fuzz(data []byte) int {
	if len(data) < 10 {
		return -1
	}

	mappingData := data[0 : len(data)/2]
	fromValue := string(mappingData[0 : len(mappingData)/2])
	toValue := string(mappingData[len(mappingData)/2:])
	searchString := string(data[len(data)/2:])

	i := 0
	for range matches(searchString, fromValue, toValue) {
		i++
	}

	if i == 0 {
		return -1
	}

	return 1
}
