package common

// AddToMapping adds one or more mappings to one value to a map.
func AddToMapping(target map[string]string, froms []string, to string) {
	for _, from := range froms {
		target[from] = to
	}
}
