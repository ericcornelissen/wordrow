package strings

import (
	"testing"
)

func TestAny(t *testing.T) {
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

func TestIsEmpty(t *testing.T) {
	result := IsEmpty("")
	if result == false {
		t.Error("Unexpected result `false` for empty string")
	}

	result = IsEmpty("Hello world!")
	if result == true {
		t.Error("Unexpected result `true` for non-empty string")
	}
}

func TestSplit(t *testing.T) {
	result := Split("foo,bar", ",")

	if len(result) != 2 {
		t.Fatalf("Unexpected length from Split (got %d)", len(result))
	}

	if result[0] != "foo" {
		t.Errorf("Unexpected first value in result (got '%s')", result[0])
	}

	if result[1] != "bar" {
		t.Errorf("Unexpected second value in result (got '%s')", result[1])
	}
}

func TestTrimSpace(t *testing.T) {
	result := TrimSpace(" foobar ")

	if result != "foobar" {
		t.Fatalf("Unexpected trimmed value (got '%s')", result)
	}
}

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
