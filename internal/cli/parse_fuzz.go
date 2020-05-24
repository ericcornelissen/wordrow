package cli

import "strings"

// Fuzz will fuzz the CLI arguments parser.
func Fuzz(data []byte) int {
	s := string(data)
	args := strings.Split(s, ",")
	if len(args) < 2 {
		return -1
	}

	ParseArgs(args)
	return 0
}
