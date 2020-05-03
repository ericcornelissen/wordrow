package replacer

import "fmt"
import "testing"

import "github.com/ericcornelissen/wordrow/internal/wordmaps"


func reportIncorrectReplacement(t *testing.T, expected, actual string) {
  t.Errorf(`Replacement did not work as intended
    expected : '%s'
    got      : '%s'
  `, expected, actual)
}


func TestReplaceEmptyString(t *testing.T) {
  var wm wordmaps.WordMap

  source := ""
  result := ReplaceAll(source, wm)

  if result != source {
    t.Errorf("Result was not en empty string but: '%s'", result)
  }
}

func TestReplaceEmptyWordmap(t *testing.T) {
  var wm wordmaps.WordMap

  source := "Hello world!"
  result := ReplaceAll(source, wm)

  if result != source {
    reportIncorrectReplacement(t, source, result)
  }
}

func TestReplaceOneWordInWordMap(t *testing.T) {
  from, to := "foo", "bar"

  var wm wordmaps.WordMap
  wm.AddOne(from, to)

  t.Run("source is 'from' in the WordMap", func(t *testing.T) {
    source := from
    result := ReplaceAll(source, wm)

    if result != to {
      reportIncorrectReplacement(t, to, result)
    }
  })
  t.Run("source is 'to' in the WordMap", func(t *testing.T) {
    source := to
    result := ReplaceAll(source, wm)

    if result != source {
      reportIncorrectReplacement(t, to, result)
    }
  })
  t.Run("One line", func(t *testing.T) {
    template := "This is a %s."
    source := fmt.Sprintf(template, from)
    result := ReplaceAll(source, wm)

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
    result := ReplaceAll(source, wm)

    expected := fmt.Sprintf(template, to, to, to)
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}

func TestReplaceMultipleWordsInWordMap(t *testing.T) {
  var wm wordmaps.WordMap
  wm.AddOne("foo", "bar")
  wm.AddOne("color", "colour")

  t.Run("All words", func(t *testing.T) {
    source := "A foo is a creature in this world. It can change its color."
    result := ReplaceAll(source, wm)

    expected := "A bar is a creature in this world. It can change its colour."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("Only one word", func(t *testing.T) {
    source := "A foo is a creature in this world."
    result := ReplaceAll(source, wm)

    expected := "A bar is a creature in this world."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}

func TestReplaceWhiteSpaceInPhrase(t *testing.T) {
  t.Run("single space", func(t *testing.T) {
    from, to := "foo bar", "foobar"

    var wm wordmaps.WordMap
    wm.AddOne(from, to)

    source := from
    result := ReplaceAll(source, wm)
    if result != to {
      reportIncorrectReplacement(t, to, result)
    }

    source = "foo  bar"
    result = ReplaceAll(source, wm)
    if result != to {
      reportIncorrectReplacement(t, to, result)
    }
  })
  t.Run("two spaces", func(t *testing.T) {
    from, to := "a  dog", "an amazing dog"

    var wm wordmaps.WordMap
    wm.AddOne(from, to)

    source := from
    result := ReplaceAll(source, wm)
    if result != "an  amazing dog" {
      reportIncorrectReplacement(t, to, result)
    }

    source = "a dog"
    result = ReplaceAll(source, wm)
    if result != to {
      reportIncorrectReplacement(t, to, result)
    }
  })
}

func TestReplaceIgnoreCapitalizationInMapping(t *testing.T) {
  var wm wordmaps.WordMap
  wm.AddOne("Foo", "Bar")

  source := "There once was a foo in the world."
  result := ReplaceAll(source, wm)

  expected := "There once was a bar in the world."
  if result != expected {
    reportIncorrectReplacement(t, expected, result)
  }
}

func TestReplaceMaintainCapitalization(t *testing.T) {
  var wm wordmaps.WordMap
  wm.AddOne("foo", "bar")
  wm.AddOne("hello world", "hey planet")
  wm.AddOne("so called", "so-called")

  t.Run("single word mapping", func(t *testing.T) {
    source := "There once was a foo in the world. Foo did things."
    result := ReplaceAll(source, wm)

    expected := "There once was a bar in the world. Bar did things."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("two word mapping", func(t *testing.T) {
    source := "Hello World!"
    result := ReplaceAll(source, wm)

    expected := "Hey Planet!"
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("two word to hyphenated word mapping", func(t *testing.T) {
    source := "A So called 'hypnotoad'"
    result := ReplaceAll(source, wm)

    expected := "A So-called 'hypnotoad'"
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }

    source = "A So Called 'hypnotoad'"
    result = ReplaceAll(source, wm)

    expected = "A So-Called 'hypnotoad'"
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}

func TestReplaceWordAllCaps(t *testing.T) {
  var wm wordmaps.WordMap
  wm.AddOne("foo", "bar")

  source := "This is the FOO."
  result := ReplaceAll(source, wm)

  expected := "This is the BAR."
  if result != expected {
    reportIncorrectReplacement(t, expected, result)
  }
}

func TestReplaceWordWithPrefixes(t *testing.T) {
  t.Run("maintain prefix", func(t *testing.T) {
    var wm wordmaps.WordMap
    wm.AddOne("-ize", "-ise")

    source := "They Realize that they should not idealize."
    result := ReplaceAll(source, wm)

    expected := "They Realise that they should not idealise."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("replace only if preceeded by another word", func(t *testing.T) {
    var wm wordmaps.WordMap
    wm.AddOne("- dogs", "- cats")

    source := "Dogs are nice and dogs are cool."
    result := ReplaceAll(source, wm)

    expected := "Dogs are nice and cats are cool."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("omit prefix", func(t *testing.T) {
    var wm wordmaps.WordMap
    wm.AddOne("-phone", "phone")

    source := "That cat has a telephone."
    result := ReplaceAll(source, wm)

    expected := "That cat has a phone."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("omit the preceding word", func(t *testing.T) {
    var wm wordmaps.WordMap
    wm.AddOne("- people", "people")

    source := "Cool people are nice and nice people are cool."
    result := ReplaceAll(source, wm)

    expected := "People are nice and people are cool."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}

func TestReplaceWordWithSuffixes(t *testing.T) {
  t.Run("maintain suffix", func(t *testing.T) {
    var wm wordmaps.WordMap
    wm.AddOne("color-", "colour-")

    source := "The colors on this colorful painting are amazing."
    result := ReplaceAll(source, wm)

    expected := "The colours on this colourful painting are amazing."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("replace only if succeeded by another word", func(t *testing.T) {
    var wm wordmaps.WordMap
    wm.AddOne("dog -", "cat -")

    source := "I have a dog and you have a dog."
    result := ReplaceAll(source, wm)

    expected := "I have a cat and you have a dog."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("maintain the succeeding word", func(t *testing.T) {
    var wm wordmaps.WordMap
    wm.AddOne("very -", "super -")

    source := "This is a very special day."
    result := ReplaceAll(source, wm)

    expected := "This is a super special day."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("omit suffix", func(t *testing.T) {
    var wm wordmaps.WordMap
    wm.AddOne("dog-", "dog")

    source := "I have a dog, but you have a small doggy."
    result := ReplaceAll(source, wm)

    expected := "I have a dog, but you have a small dog."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("omit the succeeding word", func(t *testing.T) {
    var wm wordmaps.WordMap
    wm.AddOne("a -", "a")

    source := "I have a particularly cool dog."
    result := ReplaceAll(source, wm)

    expected := "I have a cool dog."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}

func TestReplaceWordWithPrefixesAndSuffixes(t *testing.T) {
  t.Run("maintain both", func(t *testing.T) {
    var wm wordmaps.WordMap
    wm.AddOne("-bloody-", "-freaking-")

    source := "It is a fanbloodytastic movie."
    result := ReplaceAll(source, wm)

    expected := "It is a fanfreakingtastic movie."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("omit prefix, maintain suffix", func(t *testing.T) {
    var wm wordmaps.WordMap
    wm.AddOne("-b-", "b-")

    source := "abc"
    result := ReplaceAll(source, wm)

    expected := "bc"
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("maintain prefix, omit suffix", func(t *testing.T) {
    var wm wordmaps.WordMap
    wm.AddOne("-b-", "-b")

    source := "abc"
    result := ReplaceAll(source, wm)

    expected := "ab"
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("omit both", func(t *testing.T) {
    var wm wordmaps.WordMap
    wm.AddOne("-b-", "b")

    source := "abc"
    result := ReplaceAll(source, wm)

    expected := "b"
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}

func TestReplaceWordWithoutPrefixes(t *testing.T) {
  var wm wordmaps.WordMap
  wm.AddOne("mail", "email")

  source := "I send them a mail. And later another email."
  result := ReplaceAll(source, wm)

  expected := "I send them a email. And later another email."
  if result != expected {
    reportIncorrectReplacement(t, expected, result)
  }
}

func TestReplaceWordWithoutSuffixes(t *testing.T) {
  var wm wordmaps.WordMap
  wm.AddOne("commen", "common")

  source := "He game a comment that that is quite commen"
  result := ReplaceAll(source, wm)

  expected := "He game a comment that that is quite common"
  if result != expected {
    reportIncorrectReplacement(t, expected, result)
  }
}

func TestReplaceByShorterString(t *testing.T) {
  var wm wordmaps.WordMap
  wm.AddOne("fooo", "foo")

  t.Run("one instance of word", func(t *testing.T) {
    source := "This is a fooo."
    result := ReplaceAll(source, wm)

    expected := "This is a foo."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("multiple instances of word", func(t *testing.T) {
    source := "This is a FOOO and this is a fooo as well."
    result := ReplaceAll(source, wm)

    expected := "This is a FOO and this is a foo as well."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}

func TestReplaceByLongerString(t *testing.T) {
  var wm wordmaps.WordMap
  wm.AddOne("fo", "foo")

  t.Run("one instance of word", func(t *testing.T) {
    source := "This is a fo."
    result := ReplaceAll(source, wm)

    expected := "This is a foo."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("multiple instances of word", func(t *testing.T) {
    source := "This is a FO and this is a fo as well."
    result := ReplaceAll(source, wm)

    expected := "This is a FOO and this is a foo as well."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}

func TestReplacePhraseNewlineInSource(t *testing.T) {
  var wm wordmaps.WordMap
  wm.AddOne("foo bar", "foobar")
  wm.AddOne("hello world", "hey planet")
  wm.AddOne("hello beautiful world", "hey planet")
  wm.AddOne("a dog", "an amazing dog")

  t.Run("newline without indentation", func(t *testing.T) {
    source := "lorem ipsum hello\nworld dolor sit amet."
    result := ReplaceAll(source, wm)

    expected := "lorem ipsum hey\nplanet dolor sit amet."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("newline with indentation", func(t *testing.T) {
    source := "lorem ipsum hello\n  world dolor sit amet."
    result := ReplaceAll(source, wm)

    expected := "lorem ipsum hey\n  planet dolor sit amet."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("space in from but not in to", func(t *testing.T) {
    source := "lorem ipsum foo\nbar dolor sit amet."
    result := ReplaceAll(source, wm)

    expected := "lorem ipsum foobar\ndolor sit amet."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("space in from but not in to, with indentation", func(t *testing.T) {
    source := "lorem ipsum foo\n  bar dolor sit amet."
    result := ReplaceAll(source, wm)

    expected := "lorem ipsum foobar\n  dolor sit amet."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("less spaces in from than in to", func(t *testing.T) {
    source := "lorem ipsum a\ndog dolor sit amet."
    result := ReplaceAll(source, wm)

    expected := "lorem ipsum an\namazing dog dolor sit amet."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
  t.Run("more spaces in from than in to", func(t *testing.T) {
    source := "lorem ipsum hello\nbeautiful world dolor sit amet."
    result := ReplaceAll(source, wm)

    expected := "lorem ipsum hey\nplanet dolor sit amet."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }

    source = "lorem ipsum hello beautiful\nworld dolor sit amet."
    result = ReplaceAll(source, wm)

    expected = "lorem ipsum hey planet\ndolor sit amet."
    if result != expected {
      reportIncorrectReplacement(t, expected, result)
    }
  })
}
