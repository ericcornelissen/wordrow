package common

import (
	"bytes"

	"github.com/ericcornelissen/wordrow/internal/mappings/errors"
)

// AddValuesToMap adds the values defined to the provided map such that each
// value other than the last is mapped to te last value.
func AddValuesToMap(mapping map[string]string, values [][]byte) {
	last := len(values) - 1
	to := string(values[last])
	for _, from := range values[0:last] {
		mapping[string(from)] = to
	}
}

// TrimValues output all input values trimmed, or an error if any of the trimmed
// values is empty.
func TrimValues(inp [][]byte) ([][]byte, error) {
	out := make([][]byte, len(inp))

	for i, value := range inp {
		out[i] = bytes.TrimSpace(value)
		if len(out[i]) == 0 {
			return nil, errors.Newf("Empty value '%s'", inp[i])
		}
	}

	return out, nil
}
