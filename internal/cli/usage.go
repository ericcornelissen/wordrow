package cli

import (
	"fmt"
	"strings"
)

// The maximum line length for the usage message.
const maxLineLen = 80

// Utility function to clean a source-formatted string into a single line
// string.
func clean(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.TrimSpace(s)

	return s
}

// Utility function to get a string of spaces the length of a given string.
func asWhitespace(s string) string {
	return strings.Repeat(" ", len(s))
}

// Get an option and option alias as a bullet for the usage message.
func getOptionBullet(option, optionAlias string) string {
	if optionAlias == "" {
		return fmt.Sprintf("  %s:", option)
	}

	return fmt.Sprintf("  %s, %s:", option, optionAlias)
}

// Format the usage of a single option.
func formatOption(option Option, message string) string {
	var sb strings.Builder
	var lineCount = 1

	message = clean(message)

	topic := getOptionBullet(option.name, option.alias)
	indentation := asWhitespace(topic)

	sb.WriteString(topic)
	for _, word := range strings.Fields(message) {
		if (sb.Len() + len(word)) > (lineCount * maxLineLen) {
			sb.WriteRune('\n')
			sb.WriteString(indentation)

			lineCount++
		}

		sb.WriteRune(' ')
		sb.WriteString(word)
	}

	return sb.String()
}

// Print the usage of a single option.
func printOption(option Option, message string) {
	optionDoc := formatOption(option, message)
	fmt.Println(optionDoc)
}

func printSectionTitle(title string) {
	fmt.Printf("\n%s:\n", title)
}

// Print the usage of the options for the program.
func printOptions() {
	printSectionTitle("Flags")
	printOption(helpFlag, `Output this help message.`)
	printOption(versionFlag, `Output the version number of wordrow.`)
	printOption(dryRunFlag, `Don't make any changes to the input files.`)
	printOption(invertFlag, `Invert all specified mappings.`)
	printOption(silentFlag, `Don't output informative logging.`)
	printOption(verboseFlag, `Output debug logging.`)

	printSectionTitle("Options")
	printOption(configOption, `Specify a configuration file.`)
	printOption(mapfileOption, `
		Specify a dictionary file. To use multiple dictionary files you can use this
		option multiple times.
	`)
	printOption(mappingOption, `
		Specify a single mapping. Use a comma to separate the words of the mapping.
		If spaces are required use quotation marks. This option can be used multiple
		times.
	`)
}

// Print the usage of the CLI of the program.
func printInterface() {
	base := "Usage: wordrow"
	indentation := asWhitespace(base)

	fmt.Printf("%s [%s] [%s]\n",
		base,
		helpFlag.name,
		versionFlag.name,
	)
	fmt.Printf("%s [%s] [%s | %s]\n",
		indentation,
		dryRunFlag.name,
		silentFlag.alias,
		silentFlag.name,
	)
	fmt.Printf("%s [%s | %s <file>]\n",
		indentation,
		configOption.alias,
		configOption.name,
	)
	fmt.Printf("%s [%s | %s <file>]\n",
		indentation,
		mapfileOption.alias,
		mapfileOption.name,
	)
	fmt.Printf("%s [%s | %s <file>]\n",
		indentation,
		mappingOption.alias,
		mappingOption.name,
	)
	fmt.Printf("%s <files>\n", indentation)
}

// Print the usage message of the program.
func printUsage() {
	printInterface()
	printOptions()
}
