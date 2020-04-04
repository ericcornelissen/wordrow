package cli

import "fmt"
import "strings"

import "github.com/ericcornelissen/wordrow/internal/logger"


// The maximum line length for the usage message.
const maxLineLen = 80


// Utility function to get a string of spaces the length of a given string.
func asWhitespace(s string) string {
  return strings.Repeat(" ", len(s))
}


// Print the usage of a single option.
func printOption(option, optionAlias, msg string) {
  var bullet string
  if optionAlias == "" {
    bullet = fmt.Sprintf("  %s :", option)
  } else {
    bullet = fmt.Sprintf("  %s, %s :", option, optionAlias)
  }
  indent := asWhitespace(bullet)

  var sentences []string
  sentence := bullet
  for _, word := range strings.Fields(msg) {
    if len(sentence + " " + word) > maxLineLen {
      sentences = append(sentences, sentence)
      sentence = indent
    }
    sentence = sentence + " " + word
  }
  sentences = append(sentences, sentence)

  for _, sentence := range sentences {
    logger.Println(sentence)
  }
}

// Print the usage of the options for the program.
func printOptions() {
  logger.Println("Options:")
  printOption(helpOption, "", "Output this help message.")
  printOption(versionOption, "", "Output the version number of wordrow.")
  printOption(dryRunOption, "", "Run wordrow without writing changes back to the input files.")
  printOption(silentOptionAlias, silentOption, "Don't output informative logging.")
  printOption(configOptionAlias, configOption, "Specify a configuration file.")
  printOption(mapfileOptionAlias, mapfileOption, "Specify a dictionary file. To use multiple dictionary files you can use this option multiple times.")
}

// Print the usage message of the program.
func printUsage() {
  logger.Printf("Usage: wordrow [%s] [%s]\n", helpOption, versionOption)
  logger.Printf("               [%s | %s <file>]\n", configOptionAlias, configOption)
  logger.Printf("               [%s | %s <file>]\n", mapfileOptionAlias, mapfileOption)
  logger.Printf("               [%s] [%s | %s]\n", dryRunOption, silentOptionAlias, silentOption)
  logger.Println("               <files>")
  logger.Println()
  printOptions()
}
