package wordmap

import "strings"

import "github.com/ericcornelissen/wordrow/internal/errors"
import "github.com/ericcornelissen/wordrow/internal/fs"
import "github.com/ericcornelissen/wordrow/internal/logger"


// The WordMap type provides a guaranteed mapping from one set of strings to
// another set of strings.
type WordMap struct {
  from []string
  to []string
}


// Check if a given index is within the range of the WordMap.
func (m *WordMap) inRange(i int) bool {
  return i < 0 || i >= m.Size()
}


// Parse a File and add its mapping to the WordMap.
//
// The function sets the error if an error occurs when parsing the File.
func (m *WordMap) AddFile(file fs.File) error {
  parserFn, err := getParserForFile(file.Path)
  if err != nil {
    return errors.Newf("Unknown file type of %s", file.Path)
  }

  err = parseFile(&file.Content, parserFn, m)
  if err != nil {
    return errors.Newf("Error when parsing %s: %s", file.Path, err)
  }

  return nil
}

// Add all mappings from another WordMap to the WordMap.
func (m *WordMap) AddFrom(other WordMap) {
  m.from = append(m.from, other.from...)
  m.to = append(m.to, other.to...)
}

// Add a single mapping from one word to another to the WordMap.
//
// This function panics if an empty string is provided as first or second
// argument.
func (m *WordMap) AddOne(from, to string) {
  fromValue := strings.TrimSpace(strings.ToLower(from))
  toValue := strings.TrimSpace(strings.ToLower(to))
  if fromValue == "" || toValue == "" {
    panic(1)
  }

  m.from = append(m.from, fromValue)
  m.to = append(m.to, toValue)
}

// Check whether the WordMap contains a mapping from a certain string. Note that
// this only returns `true` if the queried string is in the "from" part of the
// mapping.
func (m *WordMap) Contains(x string) bool {
  for _, y := range m.from {
    if x == y {
      return true
    }
  }
  return false
}

// Get the 'from' value mapped at a specific index.
//
// This function panics if the index is outside the range of the WordMap.
func (m *WordMap) GetFrom(i int) string {
  if m.inRange(i) {
    logger.Fatalf("%d is outside the range of the WordMap", i)
    panic(1)
  }

  return m.from[i]
}

// Get the 'to' value mapped at a specific index.
//
// This function panics if the index is outside the range of the WordMap.
func (m *WordMap) GetTo(i int) string {
  if m.inRange(i) {
    logger.Fatalf("%d is outside the range of the WordMap", i)
    panic(1)
  }

  return m.to[i]
}

// Change the direction of the WordMap. I.e. invert the `to` and `from` values.
func (m *WordMap) Invert() {
  tmp := m.from
  m.from = m.to
  m.to = tmp
}

// Get the contents of the WordMap as an iterable slice.
func (m *WordMap) Iter() (chan Mapping) {
  ch := make(chan Mapping)

  go func() {
    defer close(ch)

    for i := 0; i < len(m.from); i++ {
      from, to := m.from[i], m.to[i]
      ch <- Mapping{from, to}
    }
  }()

  return ch
}

// Get the size of the WordMap. I.e. the number of words mapped from some value
// a to some value b.
func (m *WordMap) Size() int {
  return len(m.from)
}

// Get the WordMap as a human readable string.
func (m *WordMap) String() string {
  var sb strings.Builder
  for i := 0; i < len(m.from); i++ {
    from, to := m.from[i], m.to[i]
    sb.WriteString("[")
    sb.WriteString(from)
    sb.WriteString("->")
    sb.WriteString(to)
    sb.WriteString("]")
  }

  return "{" + sb.String() + "}"
}
