package logger

import (
	"fmt"
	"strings"
)

// The maximum log level that should be logged.
var maxLogLevel = INFO

// SetLogLevel configures the log level for the entire program.
func SetLogLevel(newLogLevel LogLevel) {
	maxLogLevel = newLogLevel
}

// Convert a string to sentence case, i.e. starting with a capital letter.
func toSentenceCase(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

// Internal function that should be use for printing. It will only print if the
// provided `logLevel` is okay given the `maxLogLevel`.
func log(logLevel LogLevel, msg string) {
	if logLevel >= maxLogLevel {
		fmt.Printf("[%s] %s", logLevel, toSentenceCase(msg))
	}
}

// Debug prints messages as a debug message.
func Debug(msgs ...interface{}) {
	msg := fmt.Sprintln(msgs...)
	log(DEBUG, msg)
}

// Debugf formats and prints a message as a debug message.
func Debugf(msg string, args ...interface{}) {
	formattedMsg := fmt.Sprintf(msg, args...)
	Debug(formattedMsg + "\n")
}

// Error prints messages as an error message.
func Error(msgs ...interface{}) {
	msg := fmt.Sprintln(msgs...)
	log(ERROR, msg)
}

// Errorf formats and prints a message as an error message.
func Errorf(msg string, args ...interface{}) {
	formattedMsg := fmt.Sprintf(msg, args...)
	Error(formattedMsg + "\n")
}

// Fatal prints messages as a fatal message.
func Fatal(msgs ...interface{}) {
	msg := fmt.Sprintln(msgs...)
	log(FATAL, msg)
}

// Fatalf formats and prints a message as a fatal message.
func Fatalf(msg string, args ...interface{}) {
	formattedMsg := fmt.Sprintf(msg, args...)
	Fatal(formattedMsg + "\n")
}

// Info prints messages as an info message.
func Info(msgs ...interface{}) {
	msg := fmt.Sprintln(msgs...)
	log(INFO, msg)
}

// Infof formats and prints a message as an info message.
func Infof(msg string, args ...interface{}) {
	formattedMsg := fmt.Sprintf(msg, args...)
	Info(formattedMsg + "\n")
}

// Println prints messages.
func Println(msgs ...interface{}) {
	fmt.Println(msgs...)
}

// Printf formats and prints a message.
func Printf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

// Warning prints messages as a warning message.
func Warning(msgs ...interface{}) {
	msg := fmt.Sprintln(msgs...)
	log(WARNING, msg)
}

// Warningf formats and prints a message as a warning message.
func Warningf(msg string, args ...interface{}) {
	formattedMsg := fmt.Sprintf(msg, args...)
	Warning(formattedMsg + "\n")
}
