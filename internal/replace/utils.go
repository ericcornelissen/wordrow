package replace

// Get the highest integer value out of n > 1 integer values.
func maxInt(r int, options ...int) int {
	for _, option := range options {
		if option > r {
			r = option
		}
	}

	return r
}

// Get the lowest integer value out of n > 1 integer values.
func minInt(r int, options ...int) int {
	for _, option := range options {
		if option < r {
			r = option
		}
	}

	return r
}
