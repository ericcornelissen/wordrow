// +build gofuzz

package markdown

func Fuzz(data []byte) int {
	s := string(data)
	Parse(&s)
	return 0
}
