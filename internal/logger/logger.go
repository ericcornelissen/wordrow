package logger

import "fmt"
import "strings"


// Convert a string to sentence case, i.e. starting with a capital letter.
func toSentenceCase(s string) string {
  return strings.ToUpper(s[:1]) + s[1:]
}


// Print a set of messages as a debug message.
func Debug(msgs ... interface{}) {
  msg := fmt.Sprintln(msgs...)
  fmt.Printf("[D] %s", toSentenceCase(msg))
}

// Print and format a message as a debug message.
func Debugf(msg string, args ...interface{}) {
  formattedMsg := fmt.Sprintf(msg, args...)
  Debug(formattedMsg + "\n")
}

// Print a set of messages as an error message.
func Error(msgs ...interface{}) {
  msg := fmt.Sprintln(msgs...)
  fmt.Printf("[E] %s", toSentenceCase(msg))
}

// Print and format a message as an error message.
func Errorf(msg string, args ...interface{}) {
  formattedMsg := fmt.Sprintf(msg, args...)
  Error(formattedMsg + "\n")
}

// Print a set of messages as a fatal message.
func Fatal(msgs ...interface{}) {
  msg := fmt.Sprintln(msgs...)
  fmt.Printf("[F] %s", toSentenceCase(msg))
}

// Print and format a message as a fatal message.
func Fatalf(msg string, args ...interface{}) {
  formattedMsg := fmt.Sprintf(msg, args...)
  Fatal(formattedMsg + "\n")
}

// Print a set of messages as an info message.
func Info(msgs ...interface{}) {
  msg := fmt.Sprintln(msgs...)
  fmt.Printf("[I] %s", toSentenceCase(msg))
}

// Print and format a message as an info message.
func Infof(msg string, args ...interface{}) {
  formattedMsg := fmt.Sprintf(msg, args...)
  Info(formattedMsg + "\n")
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
  fmt.Printf("[W] %s", toSentenceCase(msg))
}

// Print and format a message as a warning message.
func Warningf(msg string, args ...interface{}) {
  formattedMsg := fmt.Sprintf(msg, args...)
  Warning(formattedMsg + "\n")
}
