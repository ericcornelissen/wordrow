// +build gofuzz

package mappings

func FuzzMarkDown(data []byte) int {
	s := string(data)
	parseMarkDownFile(&s)
	return 0
}