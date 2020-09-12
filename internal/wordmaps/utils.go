package wordmaps

func addMany(target map[string]string, froms []string, to string) {
	for _, from := range froms {
		target[from] = to
	}
}
