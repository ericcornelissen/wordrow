package logger


// The LogLevel type is used to determine what should and what shouldn't be
// logged.
type LogLevel int

// The LogLevel enum values.
const (
  // Log everything.
  DEBUG LogLevel = iota

  // Log informative messages and higher.
  INFO

  // Log warning messages and higher.
  WARNING

  // Log error messages and higher.
  ERROR

  // Only log fatal crashes.
  FATAL
)

// Get the LogLevel as a human readable string.
func (level LogLevel) String() string {
  names := []string{
    "Debug",
    "Info",
    "Warning",
    "Error",
    "Fatal",
  }

  return names[level]
}
