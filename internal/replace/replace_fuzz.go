// +build gofuzz

package replace

import "github.com/ericcornelissen/stringsx"

func FuzzReplaceAll(data []byte) int {
	lines := stringsx.Split(string(data), "\n")
	if len(lines) < 2 {
		return -1
	}

	mappings := stringsx.Split(lines[0], ",")
	if len(mappings) < 2 {
		return -1
	}

	mapping := make(map[string]string, 1)
	for i := 0; i < len(mappings); i += 2 {
		from := mappings[i]

		to := "charmander"
		if i+1 < len(mappings) {
			to = mappings[i+1]
		}

		mapping[from] = to
	}

	s := stringsx.Join(lines[1:], "\n")
	result := All(s, mapping)
	if result != s {
		return 1
	}

	return 0
}
