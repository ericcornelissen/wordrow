// +build gofuzz

package csv

import (
	"bufio"
	"bytes"
)

func Fuzz(data []byte) int {
	rawReader := bytes.NewReader(data)
	bufReader := bufio.NewReader(rawReader)
	Parse(bufReader)
	return 0
}
