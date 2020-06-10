package strings

import "strings"

// Builder is strings.Builder.
type Builder struct {
	sb strings.Builder
}

// Len runs strings.Builder.Len.
func (builder *Builder) Len() int {
	return builder.sb.Len()
}

// WriteRune runs strings.Builder.WriteRune.
func (builder *Builder) String() string {
	return builder.sb.String()
}

// WriteRune runs strings.Builder.WriteRune.
func (builder *Builder) WriteRune(r rune) {
	builder.sb.WriteRune(r)
}

// WriteString runs strings.Builder.WriteString.
func (builder *Builder) WriteString(s string) {
	builder.sb.WriteString(s)
}
