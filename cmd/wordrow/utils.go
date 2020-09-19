package main

import "os"

// Handler represents a function to handle a (string) value and return an error.
type handler func(value string) error

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

// Invert the map `m`. I.e. swap each (key, value)-pair in the map.
func invert(m map[string]string) map[string]string {
	inverted := make(map[string]string, len(m))
	for key, value := range m {
		inverted[value] = key
	}

	return inverted
}

// Merge the maps `target` and `other` into map `target`. Keys present in both
// `target` and `other` will end up with the value of `other`.
func merge(target, other map[string]string) {
	for key, value := range other {
		target[key] = value
	}
}

// Check if the program received input from STDIN.
//
// based on: https://stackoverflow.com/a/38612652
func hasStdin() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	return (fi.Mode() & os.ModeNamedPipe) != 0
}
