// +build gofuzz

package wordmaps

func FuzzCsv(data []byte) int {
	s := string(data)
	parseCsvFile(&s)
	return 0
}
