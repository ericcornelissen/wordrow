package replacer

import "fmt"
import "testing"

import "github.com/ericcornelissen/wordrow/internal/wordmap"


func reportIncorrectReplacement(t *testing.T, expected, actual string) {
  t.Errorf(`Replacement did not work as intended
    expected : '%s'
    got      : '%s'
  `, expected, actual)
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
    reportIncorrectReplacement(t, source, result)
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
      reportIncorrectReplacement(t, to, result)
    }
  })
  t.Run("source is 'to' in the WordMap", func(t *testing.T) {
    source := to
    result := ReplaceAll(source, wordmap)

    if result != source {
      reportIncorrectReplacement(t, to, result)
    }
  })
  t.Run("One line", func(t *testing.T) {
    template := "This is a %s."
    source := fmt.Sprintf(template, from)
    result := ReplaceAll(source, wordmap)

    expected := fmt.Sprintf(template, to)
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
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
      reportIncorrectReplacement(t, expected, result)
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
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("Only one word", func(t *testing.T) {
    source := "A foo is a creature in this world."
    result := ReplaceAll(source, wordmap)

    expected := "A bar is a creature in this world."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
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
    reportIncorrectReplacement(t, expected, result)
  }
}

func TestReplaceMaintainCapitalization(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("foo", "bar")
  wordmap.AddOne("hello world", "hey planet")
  wordmap.AddOne("so called", "so-called")

  t.Run("single word mapping", func(t *testing.T) {
    source := "There once was a foo in the world. Foo did things."
    result := ReplaceAll(source, wordmap)

    expected := "There once was a bar in the world. Bar did things."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("two word mapping", func(t *testing.T) {
    source := "Hello World!"
    result := ReplaceAll(source, wordmap)

    expected := "Hey Planet!"
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("two word to hyphenated word mapping", func(t *testing.T) {
    source := "A So called 'hypnotoad'"
    result := ReplaceAll(source, wordmap)

    expected := "A So-called 'hypnotoad'"
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }

    source = "A So Called 'hypnotoad'"
    result = ReplaceAll(source, wordmap)

    expected = "A So-Called 'hypnotoad'"
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}

func TestReplaceWordAllCaps(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("foo", "bar")

  source := "This is the FOO."
  result := ReplaceAll(source, wordmap)

  expected := "This is the BAR."
  if result != expected {
    reportIncorrectReplacement(t, expected, result)
  }
}

func TestReplaceWordWithSuffix(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("color", "colour")

  source := "The vase is colored beautifully"
  result := ReplaceAll(source, wordmap)

  expected := "The vase is coloured beautifully"
  if result != expected {
    reportIncorrectReplacement(t, expected, result)
  }
}

func TestReplaceWordWithPrefix(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("mail", "email")

  source := "I send them a mail. And later another email."
  result := ReplaceAll(source, wordmap)

  expected := "I send them a email. And later another email."
  if result != expected {
    reportIncorrectReplacement(t, expected, result)
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
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("multiple instances of word", func(t *testing.T) {
    source := "This is a FOOO and this is a fooo as well."
    result := ReplaceAll(source, wordmap)

    expected := "This is a FOO and this is a foo as well."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
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
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("multiple instances of word", func(t *testing.T) {
    source := "This is a FO and this is a fo as well."
    result := ReplaceAll(source, wordmap)

    expected := "This is a FOO and this is a foo as well."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}


// func TestTestA(t *testing.T) {
//   var wordmap wordmap.WordMap
//   wordmap.AddOne("post compromise", "post-compromise")
//
//   source := "Post Compromise"
//   result := ReplaceAll(source, wordmap)
//
//   expected := "Post-Compromise"
//   if result != expected {
//     reportIncorrectReplacement(t, expected, result)
//   }
// }
//
// func TestTest2(t *testing.T) {
//   var wordmap wordmap.WordMap
//   wordmap.AddOne("hello awesome world", "hello world")
//
//   source := "Hello Awesome World!"
//   result := ReplaceAll(source, wordmap)
//
//   expected := "Hello World!"
//   if result != expected {
//     reportIncorrectReplacement(t, expected, result)
//   }
//
//   source = "Hello awesome World!"
//   result = ReplaceAll(source, wordmap)
//
//   expected = "Hello World!"
//   if result != expected {
//     reportIncorrectReplacement(t, expected, result)
//   }
// }
