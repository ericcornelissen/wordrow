package cli

import "fmt"
import "strings"

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
func formatOption(topic, message string, lineLen int) string {
	var sb strings.Builder

	message = clean(message)
	indentation := asWhitespace(topic)

	sb.WriteString(topic)
	for _, word := range strings.Fields(message) {
		if sb.Len()+len(word) > lineLen {
			sb.WriteRune('\n')
			sb.WriteString(indentation)

			lineLen += lineLen
		}

		sb.WriteRune(' ')
		sb.WriteString(word)
	}

	return sb.String()
}

// Print the usage of a single option.
func printOption(option, optionAlias, message string) {
	topic := getOptionBullet(option, optionAlias)
	optionDoc := formatOption(topic, message, maxLineLen)
	fmt.Println(optionDoc)
}

func printSectionTitle(title string) {
	fmt.Printf("\n%s:\n", title)
}

// Print the usage of the options for the program.
func printOptions() {
	printSectionTitle("Flags")
	printOption(helpFlag, "", `Output this help message.`)
	printOption(versionFlag, "", `Output the version number of wordrow.`)
	printOption(dryRunFlag, "", `
		Run wordrow without writing changes back to the input files.
	`)
	printOption(invertFlagAlias, invertFlag, `Invert all specified mappings.`)
	printOption(silentFlagAlias, silentFlag, `Don't output informative logging.`)
	printOption(verboseFlagAlias, verboseFlag, `Output debug logging.`)

	printSectionTitle("Options")
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

// Print the usage of the CLI of the program.
func printInterface() {
	base := "Usage: wordrow"
	indentation := asWhitespace(base)

	fmt.Printf("%s [%s] [%s]\n",
		base,
		helpFlag,
		versionFlag,
	)
	fmt.Printf("%s [%s] [%s | %s]\n",
		indentation,
		dryRunFlag,
		silentFlagAlias,
		silentFlag,
	)
	fmt.Printf("%s [%s | %s <file>]\n",
		indentation,
		configOptionAlias,
		configOption,
	)
	fmt.Printf("%s [%s | %s <file>]\n",
		indentation,
		mapfileOptionAlias,
		mapfileOption,
	)
	fmt.Printf("%s [%s | %s <file>]\n",
		indentation,
		mappingOptionAlias,
		mappingOption,
	)
	fmt.Printf("%s <files>\n", indentation)
}

// Print the usage message of the program.
func printUsage() {
	printInterface()
	printOptions()
}
