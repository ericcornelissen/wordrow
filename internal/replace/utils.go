package replace

// Get the lowest integer value out of n > 1 integer values.
func lowestInt(r int, options ...int) int {
	for _, option := range options {
		if option < r {
			r = option
		}
	}

	return r
}
