// +build gofuzz

package cli

import "github.com/ericcornelissen/stringsx"

func Fuzz(data []byte) int {
	s := string(data)
	args := stringsx.Split(s, ",")
	if len(args) < 2 {
		return -1
	}

	ParseArgs(args)
	return 0
}
