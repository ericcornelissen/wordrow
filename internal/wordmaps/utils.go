package wordmaps

// Check if a slice of strings contains a certain string.
func contains(l []string, query string) bool {
	for _, v := range l {
		if v == query {
			return true
		}
	}

	return false
}
