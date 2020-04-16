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
      an %s as well. And, ow,
      over there, another %s one!
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

func TestReplaceWhiteSpaceInPhrase(t *testing.T) {
  t.Run("single space", func(t *testing.T) {
    from, to := "foo bar", "foobar"

    var wordmap wordmap.WordMap
    wordmap.AddOne(from, to)

    source := from
    result := ReplaceAll(source, wordmap)
    if result != to {
      reportIncorrectReplacement(t, to, result)
    }

    source = "foo  bar"
    result = ReplaceAll(source, wordmap)
    if result != source {
      reportIncorrectReplacement(t, source, result)
    }
  })
  t.Run("mutliple spaces", func(t *testing.T) {
    from, to := "foo  bar", "foobar"

    var wordmap wordmap.WordMap
    wordmap.AddOne(from, to)

    source := from
    result := ReplaceAll(source, wordmap)
    if result != to {
      reportIncorrectReplacement(t, to, result)
    }

    source = "foo bar"
    result = ReplaceAll(source, wordmap)
    if result != source {
      reportIncorrectReplacement(t, source, result)
    }
  })
}

func TestReplaceIgnoreCapitalizationInMapping(t *testing.T) {
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

  source := "There once was a foo in the world. Foo did things."
  result := ReplaceAll(source, wordmap)

  expected := "There once was a bar in the world. Bar did things."
  if result != expected {
    reportIncorrectReplacement(t, expected, result)
  }
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

func TestReplaceWordWithPrefixes(t *testing.T) {
  t.Run("maintain prefix", func(t *testing.T) {
    var wordmap wordmap.WordMap
    wordmap.AddOne("-ize", "-ise")

    source := "They Realize that they should not idealize."
    result := ReplaceAll(source, wordmap)

    expected := "They Realise that they should not idealise."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("replace only if preceeded by another word", func(t *testing.T) {
    var wordmap wordmap.WordMap
    wordmap.AddOne("- dogs", "- cats")

    source := "Dogs are nice and dogs are cool."
    result := ReplaceAll(source, wordmap)

    expected := "Dogs are nice and cats are cool."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("omit prefix", func(t *testing.T) {
    var wordmap wordmap.WordMap
    wordmap.AddOne("-phone", "phone")

    source := "That cat has a telephone."
    result := ReplaceAll(source, wordmap)

    expected := "That cat has a phone."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("omit the preceding word", func(t *testing.T) {
    var wordmap wordmap.WordMap
    wordmap.AddOne("- people", "people")

    source := "Cool people are nice and nice people are cool."
    result := ReplaceAll(source, wordmap)

    expected := "People are nice and people are cool."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}

func TestReplaceWordWithSuffixes(t *testing.T) {
  t.Run("maintain suffix", func(t *testing.T) {
    var wordmap wordmap.WordMap
    wordmap.AddOne("color-", "colour-")

    source := "The colors on this colorful painting are amazing."
    result := ReplaceAll(source, wordmap)

    expected := "The colours on this colourful painting are amazing."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("replace only if succeeded by another word", func(t *testing.T) {
    var wordmap wordmap.WordMap
    wordmap.AddOne("dog -", "cat -")

    source := "I have a dog and you have a dog."
    result := ReplaceAll(source, wordmap)

    expected := "I have a cat and you have a dog."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("maintain the succeeding word", func(t *testing.T) {
    var wordmap wordmap.WordMap
    wordmap.AddOne("very -", "super -")

    source := "This is a very special day."
    result := ReplaceAll(source, wordmap)

    expected := "This is a super special day."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("omit suffix", func(t *testing.T) {
    var wordmap wordmap.WordMap
    wordmap.AddOne("dog-", "dog")

    source := "I have a dog, but you have a small doggy."
    result := ReplaceAll(source, wordmap)

    expected := "I have a dog, but you have a small dog."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("omit the succeeding word", func(t *testing.T) {
    var wordmap wordmap.WordMap
    wordmap.AddOne("a -", "a")

    source := "I have a particularly cool dog."
    result := ReplaceAll(source, wordmap)

    expected := "I have a cool dog."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}

func TestReplaceWordWithPrefixesAndSuffixes(t *testing.T) {
  t.Run("maintain both", func(t *testing.T) {
    var wordmap wordmap.WordMap
    wordmap.AddOne("-bloody-", "-freaking-")

    source := "It is a fanbloodytastic movie."
    result := ReplaceAll(source, wordmap)

    expected := "It is a fanfreakingtastic movie."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("omit prefix, maintain suffix", func(t *testing.T) {
    var wordmap wordmap.WordMap
    wordmap.AddOne("-b-", "b-")

    source := "abc"
    result := ReplaceAll(source, wordmap)

    expected := "bc"
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("maintain prefix, omit suffix", func(t *testing.T) {
    var wordmap wordmap.WordMap
    wordmap.AddOne("-b-", "-b")

    source := "abc"
    result := ReplaceAll(source, wordmap)

    expected := "ab"
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("omit both", func(t *testing.T) {
    var wordmap wordmap.WordMap
    wordmap.AddOne("-b-", "b")

    source := "abc"
    result := ReplaceAll(source, wordmap)

    expected := "b"
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}

func TestReplaceWordWithoutPrefixes(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("mail", "email")

  source := "I send them a mail. And later another email."
  result := ReplaceAll(source, wordmap)

  expected := "I send them a email. And later another email."
  if result != expected {
    reportIncorrectReplacement(t, expected, result)
  }
}

func TestReplaceWordWithoutSuffixes(t *testing.T) {
  var wordmap wordmap.WordMap
  wordmap.AddOne("commen", "common")

  source := "He game a comment that that is quite commen"
  result := ReplaceAll(source, wordmap)

  expected := "He game a comment that that is quite common"
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
