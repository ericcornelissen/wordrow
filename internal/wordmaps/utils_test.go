package wordmaps

import "testing"

func TestContainsEmpty(t *testing.T) {
	var l []string

	result := contains(l, "foobar")

	if result != false {
		t.Error("Expected empty slice not to contain 'foobar'")
	}
}

func TestContainsDoesNotContain(t *testing.T) {
	l := []string{"foo", "bar"}

	result := contains(l, "foobar")

	if result != false {
		t.Error("Expected slice not to contain 'foobar'")
	}
}

func TestContainsDoesContain(t *testing.T) {
	l := []string{"foo", "bar"}

	result := contains(l, "foo")

	if result != true {
		t.Error("Expected slice to contain 'foo'")
	}
}
