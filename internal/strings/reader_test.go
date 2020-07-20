package strings

import "testing"

func TestNewReader(t *testing.T) {
	s := "It's dangerous to go alone!"
	r := NewReader(s)

	b := make([]byte, len(s))
	_, err := r.Read(b)
	if err != nil {
		t.Fatalf("Unexpected error (%s)", err)
	}

	result := string(b)
	if result != s {
		t.Errorf("Unexpected value read (got '%s')", result)
	}
}
