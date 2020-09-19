// +build gofuzz

package mappings

func FuzzCsv(data []byte) int {
	s := string(data)
	parseCsvFile(&s)
	return 0
}
