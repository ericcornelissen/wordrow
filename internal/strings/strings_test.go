package strings

import (
	"fmt"
	"testing"
)

func ExampleAny() {
	s := []string{"Hello", "world", "!"}

	any := Any(s, IsEmpty)
	fmt.Print(any)
	// Output: false
}

func ExampleMap() {
	s := []string{" foo  ", "  bar "}

	Map(s, TrimSpace)
	fmt.Print(s)
	// Output: [foo bar]
}

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

func TestContains(t *testing.T) {
	t.Run("String contains substring", func(t *testing.T) {
		substr := "foo"
		s := fmt.Sprintf("%sbar", substr)

		result := Contains(s, substr)
		if result != true {
			t.Error("Expected result to be true")
		}
	})
	//Add "Contains" function to package strings
	t.Run("String does not contain substring", func(t *testing.T) {
		substr := "foo"
		s := "bar"

		result := Contains(s, substr)
		if result != false {
			t.Error("Expected result to be false")
		}
	})
	t.Run("Empty input string", func(t *testing.T) {
		substr := "bar"
		s := ""

		result := Contains(s, substr)
		if result != false {
			t.Error("Expected result to be false")
		}
	})
	t.Run("Empty query string", func(t *testing.T) {
		substr := ""
		s := "If you see a rat the size of a car, you're playing the wrong game"

		result := Contains(s, substr)
		if result != true {
			t.Error("Expected result to be true")
		}
	})
}

func TestFields(t *testing.T) {
	result := Fields("foo bar")

	if len(result) != 2 {
		t.Fatalf("Unexpected length from Fields (got %d)", len(result))
	}

	if result[0] != "foo" {
		t.Errorf("Unexpected first value in result (got '%s')", result[0])
	}

	if result[1] != "bar" {
		t.Errorf("Unexpected second value in result (got '%s')", result[1])
	}
}

func TestHasPrefix(t *testing.T) {
	result := HasPrefix("foobar", "foo")
	if result == false {
		t.Error("Unexpected result 'false' for string")
	}

	result = HasPrefix("foobar", "bar")
	if result == true {
		t.Error("Unexpected result 'true' for string")
	}
}

func TestHasSuffix(t *testing.T) {
	result := HasSuffix("foobar", "foo")
	if result == true {
		t.Error("Unexpected result 'true' for string")
	}

	result = HasSuffix("foobar", "bar")
	if result == false {
		t.Error("Unexpected result 'false' for string")
	}
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

func TestMap(t *testing.T) {
	subject := []string{
		"a",
		" b",
		"c ",
		" d ",
	}

	Map(subject, TrimSpace)

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

func TestRepeat(t *testing.T) {
	result := Repeat("foo", 0)
	if result != "" {
		t.Errorf("Unexpected result (got '%s')", result)
	}

	result = Repeat("foo", 1)
	if result != "foo" {
		t.Errorf("Unexpected result (got '%s')", result)
	}

	result = Repeat("foo", 2)
	if result != "foofoo" {
		t.Errorf("Unexpected result (got '%s')", result)
	}
}

func TestReplaceAll(t *testing.T) {
	result := ReplaceAll("foobar bar", "bar", "baz")

	if result != "foobaz baz" {
		t.Errorf("Unexpected result (got '%s')", result)
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

func TestToLower(t *testing.T) {
	result := ToLower("FOOBAR")

	if result != "foobar" {
		t.Fatalf("Unexpected lower value (got '%s')", result)
	}
}

func TestToUpper(t *testing.T) {
	result := ToUpper("foobar")

	if result != "FOOBAR" {
		t.Fatalf("Unexpected upper value (got '%s')", result)
	}
}

func TestTrimSpace(t *testing.T) {
	result := TrimSpace(" foobar ")

	if result != "foobar" {
		t.Fatalf("Unexpected trimmed value (got '%s')", result)
	}
}

func TestTrimSuffix(t *testing.T) {
	result := TrimSuffix("foobar", "bar")
	if result != "foo" {
		t.Fatalf("Unexpected trimmed value (got '%s')", result)
	}

	result = TrimSuffix("foobar bar", "bar")
	if result != "foobar " {
		t.Fatalf("Unexpected trimmed value (got '%s')", result)
	}
}
