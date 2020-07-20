// +build gofuzz

package cli

import "strings"

func Fuzz(data []byte) int {
	s := string(data)
	args := strings.Split(s, ",")
	if len(args) < 2 {
		return -1
	}

	ParseArgs(args)
	return 0
}
