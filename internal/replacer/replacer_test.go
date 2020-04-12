package replacer

import "fmt"
import "testing"

import "github.com/ericcornelissen/wordrow/internal/wordmap"


func reportIncorrectReplacement(t *testing.T, expected, actual string) {
  t.Errorf("Replacement did not work as intended\n expected : '%s'\n got      : '%s'", expected, actual)
}


func TestReplaceEmptyString(t *testing.T) {
  var wordmap wordmap.WordMap

  source := ""
  result := ReplaceAll(source, wordmap)

  if result != source {
    t.Errorf("Result was not en empty string but: '%s'", result)
  }
}

func TestReplaceEmptyWordmap(t *testing.T) {
  var wordmap wordmap.WordMap

  source := "Hello world!"
  result := ReplaceAll(source, wordmap)

  if result != source {
    reportIncorrectReplacement(t, result, source)
  }
}

func TestReplaceOneWordInWordMap(t *testing.T) {
  from, to := "foo", "bar"

  var wordmap wordmap.WordMap
  wordmap.AddOne(from, to)

  t.Run("source is 'from' in the WordMap", func(t *testing.T) {
    source := from
    result := ReplaceAll(source, wordmap)

    if result != to {
      reportIncorrectReplacement(t, result, to)
    }
  })
  t.Run("source is 'to' in the WordMap", func(t *testing.T) {
    source := to
    result := ReplaceAll(source, wordmap)

    if result != source {
      reportIncorrectReplacement(t, result, to)
    }
  })
  t.Run("One line", func(t *testing.T) {
    template := "This is a %s."
    source := fmt.Sprintf(template, from)
    result := ReplaceAll(source, wordmap)

    expected := fmt.Sprintf(template, to)
    if result != expected {
      reportIncorrectReplacement(t, result, expected)
    }
  })
  t.Run("Multiple lines", func(t *testing.T) {
    template := `
      This is a %s. And this is
      an %s as well. Aren't %ss
      amazing!
    `
    source := fmt.Sprintf(template, from, from, from)
    result := ReplaceAll(source, wordmap)

    expected := fmt.Sprintf(template, to, to, to)
    if result != expected {
      reportIncorrectReplacement(t, result, expected)
    }
  })
}

func TestReplaceMultipleWordsInWordMap(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("foo", "bar")
  wordmap.AddOne("color", "colour")

  t.Run("All words", func(t *testing.T) {
    source := "A foo is a creature in this world. It can change its color."
    result := ReplaceAll(source, wordmap)

    expected := "A bar is a creature in this world. It can change its colour."
    if result != expected {
      reportIncorrectReplacement(t, result, expected)
    }
  })
  t.Run("Only one word", func(t *testing.T) {
    source := "A foo is a creature in this world."
    result := ReplaceAll(source, wordmap)

    expected := "A bar is a creature in this world."
    if result != expected {
      reportIncorrectReplacement(t, result, expected)
    }
  })
}

func TestReplaceIgnoreMappingCapitalization(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("Foo", "Bar")

  source := "There once was a foo in the world."
  result := ReplaceAll(source, wordmap)

  expected := "There once was a bar in the world."
  if result != expected {
    reportIncorrectReplacement(t, result, expected)
  }
}

func TestReplaceMaintainCapitalization(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("foo", "bar")

  source := "There once was a foo in the world. Foo did things."
  result := ReplaceAll(source, wordmap)

  expected := "There once was a bar in the world. Bar did things."
  if result != expected {
    reportIncorrectReplacement(t, result, expected)
  }
}

func TestReplaceWordAllCaps(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("foo", "bar")

  source := "This is the FOO."
  result := ReplaceAll(source, wordmap)

  expected := "This is the BAR."
  if result != expected {
    reportIncorrectReplacement(t, result, expected)
  }
}

func TestReplaceWordWithSuffix(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("color", "colour")

  source := "The vase is colored beautifully"
  result := ReplaceAll(source, wordmap)

  expected := "The vase is coloured beautifully"
  if result != expected {
    reportIncorrectReplacement(t, result, expected)
  }
}

func TestReplaceWordWithPrefix(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("mail", "email")

  source := "I send them a mail. And later another email."
  result := ReplaceAll(source, wordmap)

  expected := "I send them a email. And later another email."
  if result != expected {
    reportIncorrectReplacement(t, result, expected)
  }
}

func TestReplaceByShorterString(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("fooo", "foo")

  t.Run("one instance of word", func(t *testing.T) {
    source := "This is a fooo."
    result := ReplaceAll(source, wordmap)

    expected := "This is a foo."
    if result != expected {
      reportIncorrectReplacement(t, result, expected)
    }
  })
  t.Run("multiple instances of word", func(t *testing.T) {
    source := "This is a FOOO and this is a fooo as well."
    result := ReplaceAll(source, wordmap)

    expected := "This is a FOO and this is a foo as well."
    if result != expected {
      reportIncorrectReplacement(t, result, expected)
    }
  })
}

func TestReplaceByLongerString(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("fo", "foo")

  t.Run("one instance of word", func(t *testing.T) {
    source := "This is a fo."
    result := ReplaceAll(source, wordmap)

    expected := "This is a foo."
    if result != expected {
      reportIncorrectReplacement(t, result, expected)
    }
  })
  t.Run("multiple instances of word", func(t *testing.T) {
    source := "This is a FO and this is a fo as well."
    result := ReplaceAll(source, wordmap)

    expected := "This is a FOO and this is a foo as well."
    if result != expected {
      reportIncorrectReplacement(t, result, expected)
    }
  })
}

func BenchmarkReplaceOne(b *testing.B) {
  for n := 0; n < b.N; n++ {
    replaceOne("the word foo appears foo times in this foo", "foo", "bar")
  }
}
