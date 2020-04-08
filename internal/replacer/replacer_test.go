package replacer

import "testing"

import "github.com/ericcornelissen/wordrow/internal/dicts"


func reportIncorrectReplacement(t *testing.T, expected, actual string) {
  t.Errorf("Replacement did not work as intended\n expected : '%s'\n got      : '%s'", expected, actual)
}


func TestEmptyString(t *testing.T) {
  var wordmap dicts.WordMap

  source := ""
  result := ReplaceAll(source, wordmap)

  if result != source {
    t.Errorf("Result was not en empty string but: '%s'", result)
  }
}

func TestEmptyWordmap(t *testing.T) {
  var wordmap dicts.WordMap

  source := "Hello world!"
  result := ReplaceAll(source, wordmap)

  if result != source {
    reportIncorrectReplacement(t, result, source)
  }
}

func TestJustOneWord(t *testing.T) {
  from, to := "foo", "bar"

  var wordmap dicts.WordMap
  wordmap.AddOne(from, to)

  t.Run("word is in the WordMap", func(t *testing.T) {
    source := from
    result := ReplaceAll(source, wordmap)

    if result != to {
      reportIncorrectReplacement(t, result, to)
    }
  })
  t.Run("word is not in the WordMap", func(t *testing.T) {
    source := to
    result := ReplaceAll(source, wordmap)

    if result != to {
      reportIncorrectReplacement(t, result, to)
    }
  })
}

func TestOneWordOnceInText(t *testing.T) {
  var wordmap dicts.WordMap
  wordmap.AddOne("email", "e-mail")

  t.Run("One line", func(t *testing.T) {
    source := "This is an email."
    result := ReplaceAll(source, wordmap)

    expected := "This is an e-mail."
    if result != expected {
      reportIncorrectReplacement(t, result, expected)
    }
  })
  t.Run("Multiple lines", func(t *testing.T) {
    source := `
      This is an email. And this
      is an email as well. Emails
      are amazing, right?
    `
    result := ReplaceAll(source, wordmap)

    expected := `
      This is an e-mail. And this
      is an e-mail as well. E-mails
      are amazing, right?
    `
    if result != expected {
      reportIncorrectReplacement(t, result, expected)
    }
  })
}

func TestMultipleWords(t *testing.T) {
  var wordmap dicts.WordMap
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

func TestIgnoreInputCapitalization(t *testing.T) {
  var wordmap dicts.WordMap
  wordmap.AddOne("Foo", "Bar")

  source := "There once was a foo in the world."
  result := ReplaceAll(source, wordmap)

  expected := "There once was a bar in the world."
  if result != expected {
    reportIncorrectReplacement(t, result, expected)
  }
}

func TestMaintainCapitalization(t *testing.T) {
  var wordmap dicts.WordMap
  wordmap.AddOne("foo", "bar")

  source := "There once was a foo in the world. Foo did things."
  result := ReplaceAll(source, wordmap)

  expected := "There once was a bar in the world. Bar did things."
  if result != expected {
    reportIncorrectReplacement(t, result, expected)
  }
}

func TestWordWithSuffix(t *testing.T) {
  var wordmap dicts.WordMap
  wordmap.AddOne("color", "colour")

  source := "The vase is colored beautifully"
  result := ReplaceAll(source, wordmap)

  expected := "The vase is coloured beautifully"
  if result != expected {
    reportIncorrectReplacement(t, result, expected)
  }
}

func TestWordWithPrefix(t *testing.T) {
  var wordmap dicts.WordMap
  wordmap.AddOne("mail", "email")

  source := "I send them a mail. And later another email."
  result := ReplaceAll(source, wordmap)

  expected := "I send them a email. And later another email."
  if result != expected {
    reportIncorrectReplacement(t, result, expected)
  }
}

func TestWordAllCaps(t *testing.T) {
  var wordmap dicts.WordMap
  wordmap.AddOne("foo", "bar")

  source := "This is the FOO."
  result := ReplaceAll(source, wordmap)

  expected := "This is the BAR."
  if result != expected {
    reportIncorrectReplacement(t, result, expected)
  }
}

func TestReplaceByShorterString(t *testing.T) {
  var wordmap dicts.WordMap
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
  var wordmap dicts.WordMap
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

func TestReplaceByNothing(t *testing.T) {
  var wordmap dicts.WordMap
  wordmap.AddOne("cool", "")

  t.Run("replace all lowercase by nothing", func(t *testing.T) {
    source := "This is an awesome cool foo."
    result := ReplaceAll(source, wordmap)

    expected := "This is an awesome  foo."
    if result != expected {
      reportIncorrectReplacement(t, result, expected)
    }
  })
  t.Run("replace all uppercase by nothing", func(t *testing.T) {
    source := "This is an awesome COOL foo."
    result := ReplaceAll(source, wordmap)

    expected := "This is an awesome  foo."
    if result != expected {
      reportIncorrectReplacement(t, result, expected)
    }
  })
  t.Run("replace multiple instances by nothing", func(t *testing.T) {
    source := "This is a cool foo and that is a cool bar."
    result := ReplaceAll(source, wordmap)

    expected := "This is a  foo and that is a  bar."
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
