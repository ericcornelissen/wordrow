package logger

import "fmt"
import "strings"


// Convert a string to sentence case, i.e. starting with a capital letter.
func toSentenceCase(s string) string {
  return strings.ToUpper(s[:1]) + s[1:]
}

// Print a message as an error message.
func printError(msg string) {
  fmt.Printf("[E] %s", toSentenceCase(msg))
}

// Print a message as a fatal message.
func printFatal(msg string) {
  fmt.Printf("[F] %s", toSentenceCase(msg))
}

// Print a message as an info message.
func printInfo(msg string) {
  fmt.Printf("[I] %s", toSentenceCase(msg))
}

// Print a message as a warning message.
func printWarning(msg string) {
  fmt.Printf("[W] %s", toSentenceCase(msg))
}


// Print a set of messages as an error message.
func Error(msgs ...interface{}) {
  msg := fmt.Sprintln(msgs...)
  printError(msg)
}

// Print and format a message as an error message.
func Errorf(msg string, args ...interface{}) {
  formattedMsg := fmt.Sprintf(msg, args...)
  printError(formattedMsg + "\n")
}

// Print a set of messages as a fatal message.
func Fatal(msgs ...interface{}) {
  msg := fmt.Sprintln(msgs...)
  printFatal(msg)
}

// Print and format a message as a fatal message.
func Fatalf(msg string, args ...interface{}) {
  formattedMsg := fmt.Sprintf(msg, args...)
  printFatal(formattedMsg + "\n")
}

// Print a set of messages as an info message.
func Info(msgs ...interface{}) {
  msg := fmt.Sprintln(msgs...)
  printInfo(msg)
}

// Print and format a message as an info message.
func Infof(msg string, args ...interface{}) {
  formattedMsg := fmt.Sprintf(msg, args...)
  printInfo(formattedMsg + "\n")
}

// Print a set of messages.
func Println(msgs ...interface{}) {
  fmt.Println(msgs...)
}

// Print and format a message.
func Printf(msg string, args ...interface{}) {
  fmt.Printf(msg, args...)
}

// Print a set of messages as a warning message.
func Warning(msgs ...interface{}) {
  msg := fmt.Sprintln(msgs...)
  printWarning(msg)
}

// Print and format a message as a warning message.
func Warningf(msg string, args ...interface{}) {
  formattedMsg := fmt.Sprintf(msg, args...)
  printWarning(formattedMsg + "\n")
}
