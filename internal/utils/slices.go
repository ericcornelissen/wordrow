package utils

import "reflect"


// Get the minimum length of two slices.
//
// If either input is not a slice the return value will be -1.
func Shortest(x, y interface{}) int {
  a, b := reflect.ValueOf(x), reflect.ValueOf(y)
  if a.Kind() == reflect.Slice && b.Kind() == reflect.Slice {
    if a.Len() < b.Len() {
      return a.Len()
    } else {
      return b.Len()
    }
  }

  return -1
}
