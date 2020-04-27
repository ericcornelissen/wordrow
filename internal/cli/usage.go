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

// Get an option and option alias as a bullet for the usage message.
func getOptionBullet(option, optionAlias string) string {
  if optionAlias == "" {
    return fmt.Sprintf("  %s :", option)
  } else {
    return fmt.Sprintf("  %s, %s :", option, optionAlias)
  }
}

// Print the usage of a single option.
func printOption(option, optionAlias, msg string) {
  bullet := getOptionBullet(option, optionAlias)
  indent := asWhitespace(bullet)

  msg = strings.ReplaceAll(msg, "\n", " ")
  msg = strings.TrimSpace(msg)

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
  logger.Println("Flags:")
  printOption(helpFlag, "", `Output this help message.`)
  printOption(versionFlag, "", `Output the version number of wordrow.`)
  printOption(dryRunFlag, "", `
    Run wordrow without writing changes back to the input files.
  `)
  printOption(invertFlagAlias, invertFlag, `Invert all specified mappings.`)
  printOption(silentFlagAlias, silentFlag, `Don't output informative logging.`)
  printOption(verboseFlagAlias, verboseFlag, `Output debug logging.`)

  logger.Println("\nOptions:")
  printOption(configOptionAlias, configOption, `Specify a configuration file.`)
  printOption(mapfileOptionAlias, mapfileOption, `
    Specify a dictionary file. To use multiple dictionary files you can use this
    option multiple times.
  `)
  printOption(mappingOptionAlias, mappingOption, `
    Specify a single mapping. Use a comma to separate the words of the mapping.
    If spaces are required use quotation marks. This option can be used multiple
    times.
  `)
}

// Print the usage message of the program.
func printUsage() {
  logger.Printf("Usage: wordrow [%s] [%s]\n", helpFlag, versionFlag)
  logger.Printf("               [%s | %s <file>]\n", configOptionAlias, configOption)
  logger.Printf("               [%s | %s <file>]\n", mapfileOptionAlias, mapfileOption)
  logger.Printf("               [%s | %s <file>]\n", mappingOptionAlias, mappingOption)
  logger.Printf("               [%s] [%s | %s]\n", dryRunFlag, silentFlagAlias, silentFlag)
  logger.Println("               <files>")
  logger.Println()
  printOptions()
}
