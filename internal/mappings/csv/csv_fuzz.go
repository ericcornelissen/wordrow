// +build gofuzz

package csv

func Fuzz(data []byte) int {
	s := string(data)
	Parse(&s)
	return 0
}
