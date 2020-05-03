package cli

import "strings"
import "testing"

func TestAsWhitespace(t *testing.T) {
	original := "Hello world!"
	result := asWhitespace(original)

	if len(result) != len(original) {
		t.Errorf("Length of string as whitespace incorrect (%d != %d)", len(result), len(original))
	}

	if strings.TrimSpace(result) != "" {
		t.Errorf("The string as whitespace was not only whitespace (%s)", result)
	}
}

func TestPrintUsage(t *testing.T) {
	printUsage()
}
