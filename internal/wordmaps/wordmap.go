// Package wordmaps provides two structures for mappings and replacement. The
// first structure, `WordMap`, is a map-like that provides certain guarantees on
// its contents.
package wordmaps

import (
	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/errors"
	"github.com/ericcornelissen/wordrow/internal/logger"
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
	new := make(StringMap, len(sm))
	for key, value := range sm {
		new[value] = key
	}

	return new
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

// The WordMap type provides a guaranteed mapping from one set of strings to
// another set of strings.
type WordMap struct {
	from []string
	to   []string
}

// Check if a given index is within the range of the WordMap.
func (wm *WordMap) inRange(i int) bool {
	return i < 0 || i >= wm.Size()
}

// AddFile parses a file and adds its mapping to the WordMap.
//
// The function sets the error if an error occurs when parsing the file.
func (wm *WordMap) AddFile(content *string, format string) error {
	sm, err := parseFile(content, format)
	if err != nil {
		return errors.Newf("Error when parsing file: %s", err)
	}

	for key, value := range sm {
		wm.AddOne(key, value)
	}

	return nil
}

// AddFrom adds all mappings from another WordMap to the WordMap.
func (wm *WordMap) AddFrom(other WordMap) {
	wm.from = append(wm.from, other.from...)
	wm.to = append(wm.to, other.to...)
}

// AddOne adds a single mapping from one word to another to the WordMap.
//
// This function panics if an empty string is provided.
func (wm *WordMap) AddOne(from, to string) {
	fromValue := stringsx.TrimSpace(from)
	toValue := stringsx.TrimSpace(to)
	if fromValue == "" || toValue == "" {
		panic(1)
	}

	wm.from = append(wm.from, fromValue)
	wm.to = append(wm.to, toValue)
}

// AddMany adds multiple mappings from multiple words to a single word to the
// WordMap.
//
// This function panics if an empty string is added.
func (wm *WordMap) AddMany(froms []string, to string) {
	for _, from := range froms {
		wm.AddOne(from, to)
	}
}

// Contains checks whether the WordMap contains a mapping from a certain string.
// Note that this only returns `true` if the queried string is in the "from"
// part of the WordMap.
func (wm *WordMap) Contains(query string) bool {
	for _, from := range wm.from {
		if query == from {
			return true
		}
	}
	return false
}

// GetFrom returns the 'from' value mapped at a specific index.
//
// This function panics if the index is outside the range of the WordMap.
func (wm *WordMap) GetFrom(i int) string {
	if wm.inRange(i) {
		logger.Fatalf("%d is outside the range of the WordMap", i)
		panic(1)
	}

	return wm.from[i]
}

// GetTo returns the 'to' value mapped at a specific index.
//
// This function panics if the index is outside the range of the WordMap.
func (wm *WordMap) GetTo(i int) string {
	if wm.inRange(i) {
		logger.Fatalf("%d is outside the range of the WordMap", i)
		panic(1)
	}

	return wm.to[i]
}

// Invert changes the direction of the WordMap. I.e. it inverts the `to` and
// `from` values in the WordMap.
func (wm *WordMap) Invert() {
	tmp := wm.from
	wm.from = wm.to
	wm.to = tmp
}

// Iter returns the contents of the WordMap as an iterable. Note that the order
// of the iterable is not fixed.
func (wm *WordMap) Iter() map[string]string {
	m := make(map[string]string, wm.Size())

	for i := 0; i < wm.Size(); i++ {
		from, to := wm.from[i], wm.to[i]
		m[from] = to
	}

	return m
}

// Size returns the size of the WordMap. I.e. the number of words mapped from
// some value a to some value b.
func (wm *WordMap) Size() int {
	return len(wm.from)
}
