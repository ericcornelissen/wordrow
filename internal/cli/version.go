package cli

import "github.com/ericcornelissen/wordrow/internal/logger"

// The version of the program as a string.
const versionString = "v0.3.0-beta"

// Print the version of the program.
func printVersion() {
	logger.Printf("wordrow %s (c) Eric Cornelissen\n", versionString)
}
