package wordmaps

import (
	"strings"

	"github.com/ericcornelissen/wordrow/internal/errors"
	"github.com/ericcornelissen/wordrow/internal/logger"
)

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
	err := parseFile(content, format, wm)
	if err != nil {
		return errors.Newf("Error when parsing file: %s", err)
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
	fromValue := strings.TrimSpace(strings.ToLower(from))
	toValue := strings.TrimSpace(strings.ToLower(to))
	if fromValue == "" || toValue == "" {
		panic(1)
	}

	wm.from = append(wm.from, fromValue)
	wm.to = append(wm.to, toValue)
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

// Iter returns the contents of the WordMap as an iterable.
func (wm *WordMap) Iter() chan Mapping {
	ch := make(chan Mapping, wm.Size())

	go func() {
		defer close(ch)

		for i := 0; i < wm.Size(); i++ {
			from, to := wm.from[i], wm.to[i]
			ch <- Mapping{from, to}
		}
	}()

	return ch
}

// Size returns the size of the WordMap. I.e. the number of words mapped from
// some value a to some value b.
func (wm *WordMap) Size() int {
	return len(wm.from)
}
