package errors

import "fmt"

func ExampleNewf() {
	err := Newf("foo%s", "bar")
	fmt.Print(err)
	// Output: foobar
}

func ExampleNewIncorrectFormat() {
	err := NewIncorrectFormat("foobar")
	fmt.Print(err)
	// Output: Incorrect format (in 'foobar')
}

func ExampleNewMissingValue() {
	err := NewMissingValue("foobar")
	fmt.Print(err)
	// Output: Missing value (in 'foobar')
}
