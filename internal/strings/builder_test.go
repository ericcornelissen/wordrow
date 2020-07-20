package strings

import "testing"

func TestBuilderLen(t *testing.T) {
	var builder Builder

	result := builder.Len()
	if result != 0 {
		t.Errorf("Unexpected initial size (got %d)", result)
	}

	s := "Hello world"
	builder.sb.WriteString(s)

	result = builder.Len()
	if result != len(s) {
		t.Errorf("Unexpected size (got %d)", result)
	}
}

func TestBuilderString(t *testing.T) {
	var builder Builder

	result := builder.String()
	if result != "" {
		t.Errorf("Unexpected initial value (got '%s')", result)
	}

	s := "Hello world"
	builder.sb.WriteString(s)

	result = builder.String()
	if result != s {
		t.Errorf("Unexpected value (got '%s')", result)
	}
}

func TestBuilderWriteRune(t *testing.T) {
	var builder Builder

	builder.WriteRune('e')

	result := builder.sb.String()
	if result != "e" {
		t.Errorf("Unexpected value (got '%s')", result)
	}
}

func TestBuilderWriteString(t *testing.T) {
	var builder Builder

	builder.WriteString("foobar")

	result := builder.sb.String()
	if result != "foobar" {
		t.Errorf("Unexpected value (got '%s')", result)
	}
}
