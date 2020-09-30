package csv

import "bytes"

// The ByteMap type provides a guaranteed mapping from one set of strings to
// another set of strings.
type ByteMap struct {
	from [][]byte
	to   [][]byte
}

// AddOne adds a single mapping from one word to another to the ByteMap.
//
// This function panics if an empty string is provided.
func (bm *ByteMap) AddOne(from, to []byte) {
	fromValue := bytes.TrimSpace(from)
	toValue := bytes.TrimSpace(to)
	if len(fromValue) == 0 || len(toValue) == 0 {
		panic(1)
	}

	bm.from = append(bm.from, fromValue)
	bm.to = append(bm.to, toValue)
}

// AddMany adds multiple mappings from multiple words to a single word to the
// ByteMap.
//
// This function panics if an empty string is added.
func (bm *ByteMap) AddMany(froms [][]byte, to []byte) {
	for _, from := range froms {
		bm.AddOne(from, to)
	}
}
