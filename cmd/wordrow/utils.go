package main

import "os"

// Handler represents a function to handle a (string) value and return an error.
type handler func(value string) error

// Drains `n` items from channel `ch` and returns all non-null errors.
func drain(ch chan error, n int) (errs []error) {
	for i := 0; i < n; i++ {
		if err := <-ch; err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

// ForEach executes a handler for each of the provided values. Any error that
// occurs is accumulated and only returned once all values are handled.
func forEach(values []string, fn handler) (errs []error) {
	for _, value := range values {
		err := fn(value)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

// Check if the program received input from STDIN.
//
// based on: https://stackoverflow.com/a/38612652
func hasStdin() bool {
	stdin, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	return (stdin.Mode() & os.ModeNamedPipe) != 0
}

// Invert the map `m`. I.e. swap each (key, value)-pair in the map.
func invert(m map[string]string) map[string]string {
	inverted := make(map[string]string, len(m))
	for key, value := range m {
		inverted[value] = key
	}

	return inverted
}
