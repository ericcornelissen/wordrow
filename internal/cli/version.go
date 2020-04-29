package cli

import "github.com/ericcornelissen/wordrow/internal/logger"


// The version of the program as a string.
const VERSION_STR = "v0.2.0-beta"


// Print the version of the program.
func printVersion() {
  logger.Printf("wordrow %s (c) Eric Cornelissen\n", VERSION_STR)
}
