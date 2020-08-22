// Package wordmaps provides two structures for mappings and replacement. The
// first structure, `WordMap`, is a map-like that provides certain guarantees on
// its contents.
package wordmaps

import (
	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/errors"
)

// The StringMap type provides a mapping from one set of strings to another set
// of strings.
type StringMap map[string]string

// AddFile TODO
func (sm *StringMap) AddFile(content *string, format string) error {
	other, err := parseFile(content, format)
	if err != nil {
		return errors.Newf("Error when parsing file: %s", err)
	}

	sm.addFrom(other)
	return nil
}

// Invert TODO
func (sm StringMap) Invert() StringMap {
	inverted := make(StringMap, len(sm))
	for key, value := range sm {
		inverted[value] = key
	}

	return inverted
}

func (sm *StringMap) addFrom(other StringMap) {
	x := (*sm)
	for key, value := range other {
		x[key] = value
	}
}

func (sm *StringMap) addOne(from, to string) {
	fromValue := stringsx.TrimSpace(from)
	toValue := stringsx.TrimSpace(to)
	if fromValue == "" || toValue == "" {
		panic(1)
	}

	(*sm)[from] = to
}

func (sm *StringMap) addMany(froms []string, to string) {
	for _, from := range froms {
		sm.addOne(from, to)
	}
}
