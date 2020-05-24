package wordmaps

// FuzzCsv will fuzz the CSV parser.
func FuzzCsv(data []byte) int {
	s := string(data)
	parseCsvFile(&s)
	return 0
}
