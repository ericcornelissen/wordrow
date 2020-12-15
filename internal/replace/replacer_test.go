package replace

import (
	"bytes"
	"fmt"
	"testing"
)

func ExampleReplaceAll() {
	mapping := make(map[string]string)
	mapping["hello"] = "hey"
	mapping["world"] = "planet"

	s := []byte("Hello world!")
	out := All(s, mapping)
	fmt.Print(string(out))
	// Output: Hey planet!
}

func TestReplaceEmptyString(t *testing.T) {
	mapping := make(map[string]string)

	source := []byte("")
	result := All(source, mapping)

	if !bytes.Equal(result, source) {
		t.Errorf("Result was not en empty string but: '%s'", result)
	}
}

func TestReplaceEmptyMapping(t *testing.T) {
	mapping := make(map[string]string)

	source := []byte("Hello world!")
	result := All(source, mapping)

	if !bytes.Equal(result, source) {
		reportIncorrectReplacement(t, source, result)
	}
}

func TestReplaceOneWordInMapping(t *testing.T) {
	from, to := "foo", "bar"

	mapping := make(map[string]string)
	mapping[from] = to

	t.Run("source is 'from' in the Mapping", func(t *testing.T) {
		source := []byte(from)
		result := All(source, mapping)

		if !bytes.Equal(result, []byte(to)) {
			reportIncorrectReplacement(t, []byte(to), result)
		}
	})
	t.Run("source is 'to' in the Mapping", func(t *testing.T) {
		source := []byte(to)
		result := All(source, mapping)

		if !bytes.Equal(result, source) {
			reportIncorrectReplacement(t, []byte(to), result)
		}
	})
	t.Run("One line", func(t *testing.T) {
		template := "This is a %s."
		source := []byte(fmt.Sprintf(template, from))
		result := All(source, mapping)

		expected := []byte(fmt.Sprintf(template, to))
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("Multiple lines", func(t *testing.T) {
		template := `
			This is a %s. And this is
			an %s as well. And, ow,
			over there, another %s one!
		`
		source := []byte(fmt.Sprintf(template, from, from, from))
		result := All(source, mapping)

		expected := []byte(fmt.Sprintf(template, to, to, to))
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceMultipleWordsInMapping(t *testing.T) {
	mapping := make(map[string]string)
	mapping["foo"] = "bar"
	mapping["color"] = "colour"

	t.Run("All words", func(t *testing.T) {
		source := []byte("A foo is a creature in this world. It can change its color.")
		result := All(source, mapping)

		expected := []byte("A bar is a creature in this world. It can change its colour.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("Only one word", func(t *testing.T) {
		source := []byte("A foo is a creature in this world.")
		result := All(source, mapping)

		expected := []byte("A bar is a creature in this world.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceWhitespaceInPhrase(t *testing.T) {
	t.Run("single space", func(t *testing.T) {
		from, to := "foo bar", "foobar"

		mapping := make(map[string]string)
		mapping[from] = to

		source := []byte(from)
		result := All(source, mapping)
		if !bytes.Equal(result, []byte(to)) {
			reportIncorrectReplacement(t, []byte(to), result)
		}

		source = []byte("foo  bar")
		result = All(source, mapping)
		if !bytes.Equal(result, []byte(to)) {
			reportIncorrectReplacement(t, []byte(to), result)
		}
	})
	t.Run("two spaces", func(t *testing.T) {
		from, to := "a  dog", "an amazing dog"

		mapping := make(map[string]string)
		mapping[from] = to

		source := []byte(from)
		result := All(source, mapping)
		if !bytes.Equal(result, []byte("an  amazing dog")) {
			reportIncorrectReplacement(t, []byte(to), result)
		}

		source = []byte("a dog")
		result = All(source, mapping)
		if !bytes.Equal(result, []byte(to)) {
			reportIncorrectReplacement(t, []byte(to), result)
		}
	})
}

func TestReplaceIgnoreCapitalizationInMapping(t *testing.T) {
	mapping := make(map[string]string)
	mapping["Foo"] = "Bar"

	source := []byte("There once was a foo in the world.")
	result := All(source, mapping)

	expected := []byte("There once was a bar in the world.")
	if !bytes.Equal(result, expected) {
		reportIncorrectReplacement(t, expected, result)
	}
}

func TestReplaceMaintainCapitalization(t *testing.T) {
	t.Run("single word mapping", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["foo"] = "bar"

		source := []byte("There once was a foo in the world. Foo did things.")
		result := All(source, mapping)

		expected := []byte("There once was a bar in the world. Bar did things.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("two word mapping", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["hello world"] = "hey planet"

		source := []byte("Hello World!")
		result := All(source, mapping)

		expected := []byte("Hey Planet!")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("two word to hyphenated word mapping", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["so called"] = "so-called"

		source := []byte("A So called 'hypnotoad'")
		result := All(source, mapping)

		expected := []byte("A So-called 'hypnotoad'")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}

		source = []byte("A So Called 'hypnotoad'")
		result = All(source, mapping)

		expected = []byte("A So-Called 'hypnotoad'")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceWordAllCaps(t *testing.T) {
	mapping := make(map[string]string)
	mapping["foo"] = "bar"

	source := []byte("This is the FOO.")
	result := All(source, mapping)

	expected := []byte("This is the BAR.")
	if !bytes.Equal(result, expected) {
		reportIncorrectReplacement(t, expected, result)
	}
}

func TestReplaceToChangeCapitalization(t *testing.T) {
	t.Run("To titlecase", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["foo"] = "Foo"

		s := []byte(`foo FOO Foo fOO FOo`)
		actual := All(s, mapping)

		expected := []byte(`Foo Foo Foo Foo Foo`)
		if !bytes.Equal(actual, expected) {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
	t.Run("To lowercase", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["bar"] = "bar"

		s := []byte(`bar BAR Bar bAR BAr`)
		actual := All(s, mapping)

		expected := []byte(`bar bar bar bar bar`)
		if !bytes.Equal(actual, expected) {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
	t.Run("To all-caps", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["r2-d2"] = "R2-D2"

		s := []byte(`r2-d2 R2-d2 r2-D2`)
		actual := All(s, mapping)

		expected := []byte(`R2-D2 R2-D2 R2-D2`)
		if !bytes.Equal(actual, expected) {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
	t.Run("Multiple words", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["hello world"] = "Hello World"

		s := []byte(`hello world HELLO WORLD hElLo WoRlD HeLlO wOrLd`)
		actual := All(s, mapping)

		expected := []byte(`Hello World Hello World Hello World Hello World`)
		if !bytes.Equal(actual, expected) {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
	t.Run("With newline", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["hello world"] = "Hello World"

		s := []byte(`
			hello world
			hello
			world
		`)
		actual := All(s, mapping)

		expected := []byte(`
			Hello World
			Hello
			World
		`)
		if !bytes.Equal(actual, expected) {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
}

func TestReplaceWordWithPrefixes(t *testing.T) {
	t.Run("maintain prefix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["-ize"] = "-ise"

		source := []byte("They Realize that they should not idealize.")
		result := All(source, mapping)

		expected := []byte("They Realise that they should not idealise.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("replace only if preceded by another word", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["- dogs"] = "- cats"

		source := []byte("Dogs are nice and dogs are cool.")
		result := All(source, mapping)

		expected := []byte("Dogs are nice and cats are cool.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit prefix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["-phone"] = "phone"

		source := []byte("That cat has a telephone.")
		result := All(source, mapping)

		expected := []byte("That cat has a phone.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit the preceding word", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["- people"] = "people"

		source := []byte("Cool people are nice and nice people are cool.")
		result := All(source, mapping)

		expected := []byte("People are nice and people are cool.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceWordWithSuffixes(t *testing.T) {
	t.Run("maintain suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["color-"] = "colour-"

		source := []byte("The colors on this colorful painting are amazing.")
		result := All(source, mapping)

		expected := []byte("The colours on this colourful painting are amazing.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("replace only if succeeded by another word", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["dog -"] = "cat -"

		source := []byte("I have a dog and you have a dog.")
		result := All(source, mapping)

		expected := []byte("I have a cat and you have a dog.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("maintain the succeeding word", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["very -"] = "super -"

		source := []byte("This is a very special day.")
		result := All(source, mapping)

		expected := []byte("This is a super special day.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["dog-"] = "dog"

		source := []byte("I have a dog, but you have a small doggy.")
		result := All(source, mapping)

		expected := []byte("I have a dog, but you have a small dog.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit the succeeding word", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["a -"] = "a"

		source := []byte("I have a particularly cool dog.")
		result := All(source, mapping)

		expected := []byte("I have a cool dog.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceWordWithPrefixesAndSuffixes(t *testing.T) {
	t.Run("maintain both", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["-bloody-"] = "-freaking-"

		source := []byte("It is a fanbloodytastic movie.")
		result := All(source, mapping)

		expected := []byte("It is a fanfreakingtastic movie.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit prefix, maintain suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["-b-"] = "b-"

		source := []byte("abc")
		result := All(source, mapping)

		expected := []byte("bc")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("maintain prefix, omit suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["-b-"] = "-b"

		source := []byte("abc")
		result := All(source, mapping)

		expected := []byte("ab")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit both", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["-b-"] = "b"

		source := []byte("abc")
		result := All(source, mapping)

		expected := []byte("b")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceWordWithoutPrefixes(t *testing.T) {
	mapping := make(map[string]string)
	mapping["mail"] = "email"

	source := []byte("I send them a mail. And later another email.")
	result := All(source, mapping)

	expected := []byte("I send them a email. And later another email.")
	if !bytes.Equal(result, expected) {
		reportIncorrectReplacement(t, expected, result)
	}
}

func TestReplaceWordWithoutSuffixes(t *testing.T) {
	mapping := make(map[string]string)
	mapping["commen"] = "common"

	source := []byte("He game a comment that that is quite commen")
	result := All(source, mapping)

	expected := []byte("He game a comment that that is quite common")
	if !bytes.Equal(result, expected) {
		reportIncorrectReplacement(t, expected, result)
	}
}

func TestReplaceByShorterString(t *testing.T) {
	mapping := make(map[string]string)
	mapping["fooo"] = "foo"

	t.Run("one instance of word", func(t *testing.T) {
		source := []byte("This is a fooo.")
		result := All(source, mapping)

		expected := []byte("This is a foo.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("multiple instances of word", func(t *testing.T) {
		source := []byte("This is a FOOO and this is a fooo as well.")
		result := All(source, mapping)

		expected := []byte("This is a FOO and this is a foo as well.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceByLongerString(t *testing.T) {
	mapping := make(map[string]string)
	mapping["fo"] = "foo"

	t.Run("one instance of word", func(t *testing.T) {
		source := []byte("This is a fo.")
		result := All(source, mapping)

		expected := []byte("This is a foo.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("multiple instances of word", func(t *testing.T) {
		source := []byte("This is a FO and this is a fo as well.")
		result := All(source, mapping)

		expected := []byte("This is a FOO and this is a foo as well.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplacePhraseNewlineInSource(t *testing.T) {
	t.Run("newline without indentation", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["hello world"] = "hey planet"

		source := []byte("lorem ipsum hello\nworld dolor sit amet.")
		result := All(source, mapping)

		expected := []byte("lorem ipsum hey\nplanet dolor sit amet.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("newline with indentation", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["hello world"] = "hey planet"

		source := []byte("lorem ipsum hello\n  world dolor sit amet.")
		result := All(source, mapping)

		expected := []byte("lorem ipsum hey\n  planet dolor sit amet.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("space in from but not in to", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["foo bar"] = "foobar"

		source := []byte("lorem ipsum foo\nbar dolor sit amet.")
		result := All(source, mapping)

		expected := []byte("lorem ipsum foobar\ndolor sit amet.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("space in from but not in to, with indentation", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["foo bar"] = "foobar"

		source := []byte("lorem ipsum foo\n  bar dolor sit amet.")
		result := All(source, mapping)

		expected := []byte("lorem ipsum foobar\n  dolor sit amet.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("less spaces in from than in to", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["a dog"] = "an amazing dog"

		source := []byte("lorem ipsum a\ndog dolor sit amet.")
		result := All(source, mapping)

		expected := []byte("lorem ipsum an\namazing dog dolor sit amet.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("more spaces in from than in to", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["hello beautiful world"] = "hey planet"

		source := []byte("lorem ipsum hello\nbeautiful world dolor sit amet.")
		result := All(source, mapping)

		expected := []byte("lorem ipsum hey\nplanet dolor sit amet.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}

		source = []byte("lorem ipsum hello beautiful\nworld dolor sit amet.")
		result = All(source, mapping)

		expected = []byte("lorem ipsum hey planet\ndolor sit amet.")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceEscapeHyphen(t *testing.T) {
	t.Run("prefix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`\-foobar`] = `foobar`

		source := []byte(`-foobar`)
		result := All(source, mapping)

		expected := []byte(`foobar`)
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`world\-`] = `world!`

		source := []byte(`Hello world-`)
		result := All(source, mapping)

		expected := []byte(`Hello world!`)
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceEscapeEscapeCharacter(t *testing.T) {
	t.Run("prefix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`\\bar`] = `bar`

		source := []byte(`foo \bar`)
		result := All(source, mapping)

		expected := []byte(`foo bar`)
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`foo\\`] = `foo`

		source := []byte(`foo\ bar`)
		result := All(source, mapping)

		expected := []byte(`foo bar`)
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceEmptyFromValue(t *testing.T) {
	mapping := make(map[string]string)
	mapping[""] = "foobar"

	s := []byte("Hello world!")
	result := All(s, mapping)

	if !bytes.Equal(s, result) {
		t.Errorf("Unexpected result (got '%s')", result)
	}
}

func TestKeepNonExistentAffix(t *testing.T) {
	t.Run("keep non-existent prefix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`r`] = `-x`

		source := []byte(`foo r bar`)
		result := All(source, mapping)

		expected := []byte(`foo x bar`)
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("keep non-existent suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`b`] = `x-`

		source := []byte(`foo b bar`)
		result := All(source, mapping)

		expected := []byte(`foo x bar`)
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("keep non-existent prefix and suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`a`] = `-x-`

		source := []byte(`foo a bar`)
		result := All(source, mapping)

		expected := []byte(`foo x bar`)
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceCornerCases(t *testing.T) {
	t.Run("empty search string", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["foo"] = "bar"

		source := []byte{}
		result := All(source, mapping)

		if !bytes.Equal(result, source) {
			reportIncorrectReplacement(t, source, result)
		}
	})
	t.Run("empty from string", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[""] = "bar"

		source := []byte("foobar")
		result := All(source, mapping)

		if !bytes.Equal(result, source) {
			reportIncorrectReplacement(t, source, result)
		}
	})
	t.Run("empty to string", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["foo"] = ""

		source := []byte("foo bar foo")
		result := All(source, mapping)

		if !bytes.Equal(result, source) {
			reportIncorrectReplacement(t, source, result)
		}
	})
	t.Run("search string contains UTF-8 character", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["\xbf"] = "pikachu"

		source := []byte("foobar")
		result := All(source, mapping)

		if !bytes.Equal(result, source) {
			reportIncorrectReplacement(t, source, result)
		}
	})
}

func TestReplaceAffixInFromCornerCases(t *testing.T) {
	t.Run("only a hyphen", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`-`] = `x`

		source := []byte(`Hello world!`)
		result := All(source, mapping)

		if !bytes.Equal(result, source) {
			reportIncorrectReplacement(t, source, result)
		}
	})
	t.Run("only two hyphens", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`--`] = `x`

		source := []byte(`Hello world!`)
		result := All(source, mapping)

		if !bytes.Equal(result, source) {
			reportIncorrectReplacement(t, source, result)
		}
	})
	t.Run("two hyphens with a space", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`- -`] = `-`

		source := []byte(`Hello world!`)
		result := All(source, mapping)

		if !bytes.Equal(result, source) {
			reportIncorrectReplacement(t, source, result)
		}
	})
}

func TestReplaceAffixInToCornerCases(t *testing.T) {
	t.Run("only a hyphen", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`world`] = `-`

		source := []byte(`Hello world!`)
		result := All(source, mapping)

		if !bytes.Equal(result, source) {
			reportIncorrectReplacement(t, source, result)
		}
	})
	t.Run("only two hyphens", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`hello`] = `--`

		source := []byte(`Hello world!`)
		result := All(source, mapping)

		if !bytes.Equal(result, source) {
			reportIncorrectReplacement(t, source, result)
		}
	})
	t.Run("two hyphens with a space", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping[`hello`] = `- -`

		source := []byte(`Hello world!`)
		result := All(source, mapping)

		if !bytes.Equal(result, source) {
			reportIncorrectReplacement(t, source, result)
		}
	})
}

func TestDoublePrefixSuffixMatch(t *testing.T) {
	t.Run("double prefix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["b -"] = "bulbasaur"

		source := []byte("b b\nfoobar")
		result := All(source, mapping)

		expected := []byte("bulbasaur\nfoobar")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("double suffix", func(t *testing.T) {
		mapping := make(map[string]string)
		mapping["- b"] = "bulbasaur"

		source := []byte("a\nb b")
		result := All(source, mapping)

		expected := []byte("bulbasaur\nbulbasaur")
		if !bytes.Equal(result, expected) {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

var s = []byte(`
	Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec arcu est,
	consequat at scelerisque accumsan, mollis quis justo. Curabitur velit risus,
	vulputate vitae turpis at, lacinia posuere nunc. Nam ornare non quam.
`)

func BenchmarkAllWithStringMap(b *testing.B) {
	m := make(map[string]string, 3)
	m["Lorem"] = "Lroem"
	m["amet"] = "tema"
	m["mollis"] = "millos"
	for n := 0; n < b.N; n++ {
		All(s, m)
	}
}

func BenchmarkAllWithByteMap(b *testing.B) {
	m := [][]byte{
		{76, 111, 114, 101, 109}, {76, 114, 111, 101, 109},
		{97, 109, 101, 116}, {116, 101, 109, 97},
		{109, 111, 108, 108 ,105, 115}, {109 ,105, 108, 108 ,111, 115},
	}
	for n := 0; n < b.N; n++ {
		AllBytes(s, m)
	}
}

func BenchmarkAllWithByteMap2(b *testing.B) {
	m := ByteMap{
		from: [][]byte{
			[]byte("Lorem"),
			[]byte("amet"),
			[]byte("mollis"),
		},
		to: [][]byte{
			[]byte("Lroem"),
			[]byte("tema"),
			[]byte("millos"),
		},
	}
	for n := 0; n < b.N; n++ {
		AllBytes2(s, m)
	}
}
