package wordmaps

// FuzzMarkDown will fuzz the MarkDown parser.
func FuzzMarkDown(data []byte) int {
	s := string(data)
	parseMarkDownFile(&s)
	return 0
}
