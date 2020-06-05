package strings

import (
	"testing"
)

func TestTrimSpaceAll(t *testing.T) {
	subject := []string{
		"a",
		" b",
		"c ",
		" d ",
	}

	TrimSpaceAll(subject)

	if subject[0] != "a" {
		t.Errorf("Incorrect first value (got '%s')", subject[0])
	}

	if subject[1] != "b" {
		t.Errorf("Incorrect second value (got '%s')", subject[1])
	}

	if subject[2] != "c" {
		t.Errorf("Incorrect third value (got '%s')", subject[2])
	}

	if subject[3] != "d" {
		t.Errorf("Incorrect fourth value (got '%s')", subject[3])
	}
}

func TestAnyString(t *testing.T) {
	condition := func(s string) bool {
		return s == ""
	}

	t.Run("empty list", func(t *testing.T) {
		result := Any([]string{}, condition)

		if result == true {
			t.Error("Unexpected result `true` for empty list")
		}
	})
	t.Run("condition never holds", func(t *testing.T) {
		result := Any([]string{"a", "b", "c"}, condition)

		if result == true {
			t.Error("Unexpected result `true` for list")
		}
	})
	t.Run("condition holds once", func(t *testing.T) {
		result := Any([]string{"a", "", "c"}, condition)

		if result == false {
			t.Error("Unexpected result `false` for list")
		}
	})
	t.Run("condition holds often", func(t *testing.T) {
		result := Any([]string{"", "", "c"}, condition)

		if result == false {
			t.Error("Unexpected result `false` for list")
		}
	})
}

func TestIsEmptyString(t *testing.T) {
	result := IsEmpty("")
	if result == false {
		t.Error("Unexpected result `false` for empty string")
	}

	result = IsEmpty("Hello world!")
	if result == true {
		t.Error("Unexpected result `true` for non-empty string")
	}
}
