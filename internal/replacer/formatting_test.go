package replacer

import "testing"


func TestToSentenceCase(t *testing.T) {
  result := toSentenceCase("hello world!")

  if result != "Hello world!" {
    t.Errorf("Unexpected result (got '%s')", result)
  }
}
