package replace

import (
	"fmt"
	"testing"
)

func ExampleReplaceAll() {
	mapping := make(map[string]string)
	mapping["hello"] = "hey"
	mapping["world"] = "planet"

	out := All("Hello world!", mapping)
	fmt.Print(out)
	// Output: Hey planet!
}

func TestReplaceEmptyString(t *testing.T) {
	mapping := make(map[string]string)

	source := ""
	result := All(source, mapping)

	if result != source {
		t.Errorf("Result was not en empty string but: '%s'", result)
	}
}

func TestReplaceEmptyWordmap(t *testing.T) {
	mapping := make(map[string]string)

	source := "Hello world!"
	result := All(source, mapping)

	if result != source {
		reportIncorrectReplacement(t, source, result)
	}
}

func TestReplaceOneWordInWordMap(t *testing.T) {
	from, to := "foo", "bar"

	mapping := make(map[string]string)
	mapping[from] = to

	t.Run("source is 'from' in the WordMap", func(t *testing.T) {
		source := from
		result := All(source, mapping)

		if result != to {
			reportIncorrectReplacement(t, to, result)
		}
	})
	t.Run("source is 'to' in the WordMap", func(t *testing.T) {
		source := to
		result := All(source, mapping)

		if result != source {
			reportIncorrectReplacement(t, to, result)
		}
	})
	t.Run("One line", func(t *testing.T) {
		template := "This is a %s."
		source := fmt.Sprintf(template, from)
		result := All(source, mapping)

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
		result := All(source, mapping)

		expected := fmt.Sprintf(template, to, to, to)
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceMultipleWordsInWordMap(t *testing.T) {
	mapping := make(map[string]string)
	mapping["foo"] = "bar"
	mapping["color"] = "colour"

	t.Run("All words", func(t *testing.T) {
		source := "A foo is a creature in this world. It can change its color."
		result := All(source, mapping)

		expected := "A bar is a creature in this world. It can change its colour."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("Only one word", func(t *testing.T) {
		source := "A foo is a creature in this world."
		result := All(source, mapping)

		expected := "A bar is a creature in this world."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceWhitespaceInPhrase(t *testing.T) {
	t.Run("single space", func(t *testing.T) {
		from, to := "foo bar", "foobar"

		mapping := make(map[string]string)
		mapping[from] = to

		source := from
		result := All(source, mapping)
		if result != to {
			reportIncorrectReplacement(t, to, result)
		}

		source = "foo  bar"
		result = All(source, mapping)
		if result != to {
			reportIncorrectReplacement(t, to, result)
		}
	})
	t.Run("two spaces", func(t *testing.T) {
		from, to := "a  dog", "an amazing dog"

		mapping := make(map[string]string)
		mapping[from] = to

		source := from
		result := All(source, mapping)
		if result != "an  amazing dog" {
			reportIncorrectReplacement(t, to, result)
		}

		source = "a dog"
		result = All(source, mapping)
		if result != to {
			reportIncorrectReplacement(t, to, result)
		}
	})
}

func TestReplaceIgnoreCapitalizationInMapping(t *testing.T) {
	mapping := make(map[string]string)
	mapping["Foo"] = "Bar"

	source := "There once was a foo in the world."
	result := All(source, mapping)

	expected := "There once was a bar in the world."
	if result != expected {
		reportIncorrectReplacement(t, expected, result)
	}
}

func TestReplaceMaintainCapitalization(t *testing.T) {
	t.Run("single word mapping", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["foo"] = "bar"

		source := "There once was a foo in the world. Foo did things."
		result := All(source, mapping)

		expected := "There once was a bar in the world. Bar did things."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("two word mapping", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["hello world"] = "hey planet"

		source := "Hello World!"
		result := All(source, mapping)

		expected := "Hey Planet!"
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("two word to hyphenated word mapping", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["so called"] = "so-called"

		source := "A So called 'hypnotoad'"
		result := All(source, mapping)

		expected := "A So-called 'hypnotoad'"
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}

		source = "A So Called 'hypnotoad'"
		result = All(source, mapping)

		expected = "A So-Called 'hypnotoad'"
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceWordAllCaps(t *testing.T) {
	mapping := make(map[string]string)
	mapping["foo"] = "bar"

	source := "This is the FOO."
	result := All(source, mapping)

	expected := "This is the BAR."
	if result != expected {
		reportIncorrectReplacement(t, expected, result)
	}
}

func TestReplaceToChangeCapitalization(t *testing.T) {
	t.Run("To titlecase", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["foo"] = "Foo"

		s := `foo FOO Foo fOO FOo`
		actual := All(s, mapping)

		expected := `Foo Foo Foo Foo Foo`
		if actual != expected {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
	t.Run("To lowercase", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["bar"] = "bar"

		s := `bar BAR Bar bAR BAr`
		actual := All(s, mapping)

		expected := `bar bar bar bar bar`
		if actual != expected {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
	t.Run("To all-caps", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["r2-d2"] = "R2-D2"

		s := `r2-d2 R2-d2 r2-D2`
		actual := All(s, mapping)

		expected := `R2-D2 R2-D2 R2-D2`
		if actual != expected {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
	t.Run("Multiple words", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["hello world"] = "Hello World"

		s := `hello world HELLO WORLD hElLo WoRlD HeLlO wOrLd`
		actual := All(s, mapping)

		expected := `Hello World Hello World Hello World Hello World`
		if actual != expected {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
	t.Run("With newline", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["hello world"] = "Hello World"

		s := `
			hello world
			hello
			world
		`
		actual := All(s, mapping)

		expected := `
			Hello World
			Hello
			World
		`
		if actual != expected {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
}

func TestReplaceWordWithPrefixes(t *testing.T) {
	t.Run("maintain prefix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["-ize"] = "-ise"

		source := "They Realize that they should not idealize."
		result := All(source, mapping)

		expected := "They Realise that they should not idealise."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("replace only if preceded by another word", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["- dogs"] = "- cats"

		source := "Dogs are nice and dogs are cool."
		result := All(source, mapping)

		expected := "Dogs are nice and cats are cool."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit prefix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["-phone"] = "phone"

		source := "That cat has a telephone."
		result := All(source, mapping)

		expected := "That cat has a phone."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit the preceding word", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["- people"] = "people"

		source := "Cool people are nice and nice people are cool."
		result := All(source, mapping)

		expected := "People are nice and people are cool."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceWordWithSuffixes(t *testing.T) {
	t.Run("maintain suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["color-"] = "colour-"

		source := "The colors on this colorful painting are amazing."
		result := All(source, mapping)

		expected := "The colours on this colourful painting are amazing."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("replace only if succeeded by another word", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["dog -"] = "cat -"

		source := "I have a dog and you have a dog."
		result := All(source, mapping)

		expected := "I have a cat and you have a dog."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("maintain the succeeding word", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["very -"] = "super -"

		source := "This is a very special day."
		result := All(source, mapping)

		expected := "This is a super special day."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["dog-"] = "dog"

		source := "I have a dog, but you have a small doggy."
		result := All(source, mapping)

		expected := "I have a dog, but you have a small dog."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit the succeeding word", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["a -"] = "a"

		source := "I have a particularly cool dog."
		result := All(source, mapping)

		expected := "I have a cool dog."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceWordWithPrefixesAndSuffixes(t *testing.T) {
	t.Run("maintain both", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["-bloody-"] = "-freaking-"

		source := "It is a fanbloodytastic movie."
		result := All(source, mapping)

		expected := "It is a fanfreakingtastic movie."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit prefix, maintain suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["-b-"] = "b-"

		source := "abc"
		result := All(source, mapping)

		expected := "bc"
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("maintain prefix, omit suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["-b-"] = "-b"

		source := "abc"
		result := All(source, mapping)

		expected := "ab"
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit both", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["-b-"] = "b"

		source := "abc"
		result := All(source, mapping)

		expected := "b"
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceWordWithoutPrefixes(t *testing.T) {
	mapping := make(map[string]string)
	mapping["mail"] = "email"

	source := "I send them a mail. And later another email."
	result := All(source, mapping)

	expected := "I send them a email. And later another email."
	if result != expected {
		reportIncorrectReplacement(t, expected, result)
	}
}

func TestReplaceWordWithoutSuffixes(t *testing.T) {
	mapping := make(map[string]string)
	mapping["commen"] = "common"

	source := "He game a comment that that is quite commen"
	result := All(source, mapping)

	expected := "He game a comment that that is quite common"
	if result != expected {
		reportIncorrectReplacement(t, expected, result)
	}
}

func TestReplaceByShorterString(t *testing.T) {
	mapping := make(map[string]string)
	mapping["fooo"] = "foo"

	t.Run("one instance of word", func(t *testing.T) {
		source := "This is a fooo."
		result := All(source, mapping)

		expected := "This is a foo."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("multiple instances of word", func(t *testing.T) {
		source := "This is a FOOO and this is a fooo as well."
		result := All(source, mapping)

		expected := "This is a FOO and this is a foo as well."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceByLongerString(t *testing.T) {
	mapping := make(map[string]string)
	mapping["fo"] = "foo"

	t.Run("one instance of word", func(t *testing.T) {
		source := "This is a fo."
		result := All(source, mapping)

		expected := "This is a foo."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("multiple instances of word", func(t *testing.T) {
		source := "This is a FO and this is a fo as well."
		result := All(source, mapping)

		expected := "This is a FOO and this is a foo as well."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplacePhraseNewlineInSource(t *testing.T) {
	t.Run("newline without indentation", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["hello world"] = "hey planet"

		source := "lorem ipsum hello\nworld dolor sit amet."
		result := All(source, mapping)

		expected := "lorem ipsum hey\nplanet dolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("newline with indentation", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["hello world"] = "hey planet"

		source := "lorem ipsum hello\n  world dolor sit amet."
		result := All(source, mapping)

		expected := "lorem ipsum hey\n  planet dolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("space in from but not in to", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["foo bar"] = "foobar"

		source := "lorem ipsum foo\nbar dolor sit amet."
		result := All(source, mapping)

		expected := "lorem ipsum foobar\ndolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("space in from but not in to, with indentation", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["foo bar"] = "foobar"

		source := "lorem ipsum foo\n  bar dolor sit amet."
		result := All(source, mapping)

		expected := "lorem ipsum foobar\n  dolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("less spaces in from than in to", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["a dog"] = "an amazing dog"

		source := "lorem ipsum a\ndog dolor sit amet."
		result := All(source, mapping)

		expected := "lorem ipsum an\namazing dog dolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("more spaces in from than in to", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["hello beautiful world"] = "hey planet"

		source := "lorem ipsum hello\nbeautiful world dolor sit amet."
		result := All(source, mapping)

		expected := "lorem ipsum hey\nplanet dolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}

		source = "lorem ipsum hello beautiful\nworld dolor sit amet."
		result = All(source, mapping)

		expected = "lorem ipsum hey planet\ndolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceEscapeHyphen(t *testing.T) {
	t.Run("prefix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`\-foobar`] = `foobar`

		source := `-foobar`
		result := All(source, mapping)

		expected := `foobar`
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`world\-`] = `world!`

		source := `Hello world-`
		result := All(source, mapping)

		expected := `Hello world!`
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceEscapeEscapeCharacter(t *testing.T) {
	t.Run("prefix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`\\bar`] = `bar`

		source := `foo \bar`
		result := All(source, mapping)

		expected := `foo bar`
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`foo\\`] = `foo`

		source := `foo\ bar`
		result := All(source, mapping)

		expected := `foo bar`
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}
