package replace

import (
	"fmt"
	"testing"

	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

func ExampleReplaceAll() {
	var wm wordmaps.WordMap
	wm.AddOne("hello", "hey")
	wm.AddOne("world", "planet")

	out := All("Hello world!", wm)
	fmt.Print(out)
	// Output: Hey planet!
}

func TestReplaceEmptyString(t *testing.T) {
	var wm wordmaps.WordMap

	source := ""
	result := All(source, wm)

	if result != source {
		t.Errorf("Result was not en empty string but: '%s'", result)
	}
}

func TestReplaceEmptyWordmap(t *testing.T) {
	var wm wordmaps.WordMap

	source := "Hello world!"
	result := All(source, wm)

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
		result := All(source, wm)

		if result != to {
			reportIncorrectReplacement(t, to, result)
		}
	})
	t.Run("source is 'to' in the WordMap", func(t *testing.T) {
		source := to
		result := All(source, wm)

		if result != source {
			reportIncorrectReplacement(t, to, result)
		}
	})
	t.Run("One line", func(t *testing.T) {
		template := "This is a %s."
		source := fmt.Sprintf(template, from)
		result := All(source, wm)

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
		result := All(source, wm)

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
		result := All(source, wm)

		expected := "A bar is a creature in this world. It can change its colour."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("Only one word", func(t *testing.T) {
		source := "A foo is a creature in this world."
		result := All(source, wm)

		expected := "A bar is a creature in this world."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceWhitespaceInPhrase(t *testing.T) {
	t.Run("single space", func(t *testing.T) {
		from, to := "foo bar", "foobar"

		var wm wordmaps.WordMap
		wm.AddOne(from, to)

		source := from
		result := All(source, wm)
		if result != to {
			reportIncorrectReplacement(t, to, result)
		}

		source = "foo  bar"
		result = All(source, wm)
		if result != to {
			reportIncorrectReplacement(t, to, result)
		}
	})
	t.Run("two spaces", func(t *testing.T) {
		from, to := "a  dog", "an amazing dog"

		var wm wordmaps.WordMap
		wm.AddOne(from, to)

		source := from
		result := All(source, wm)
		if result != "an  amazing dog" {
			reportIncorrectReplacement(t, to, result)
		}

		source = "a dog"
		result = All(source, wm)
		if result != to {
			reportIncorrectReplacement(t, to, result)
		}
	})
}

func TestReplaceIgnoreCapitalizationInMapping(t *testing.T) {
	var wm wordmaps.WordMap
	wm.AddOne("Foo", "Bar")

	source := "There once was a foo in the world."
	result := All(source, wm)

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
		result := All(source, wm)

		expected := "There once was a bar in the world. Bar did things."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("two word mapping", func(t *testing.T) {
		source := "Hello World!"
		result := All(source, wm)

		expected := "Hey Planet!"
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("two word to hyphenated word mapping", func(t *testing.T) {
		source := "A So called 'hypnotoad'"
		result := All(source, wm)

		expected := "A So-called 'hypnotoad'"
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}

		source = "A So Called 'hypnotoad'"
		result = All(source, wm)

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
	result := All(source, wm)

	expected := "This is the BAR."
	if result != expected {
		reportIncorrectReplacement(t, expected, result)
	}
}

func TestReplaceToChangeCapitalization(t *testing.T) {
	var wm wordmaps.WordMap
	wm.AddOne(`foo`, `Foo`)
	wm.AddOne(`bar`, `bar`)
	wm.AddOne(`r2-d2`, `R2-D2`)
	wm.AddOne(`hello world`, `Hello World`)

	t.Run("To title", func(t *testing.T) {
		s := `foo FOO Foo fOO FOo`
		actual := All(s, wm)

		expected := `Foo Foo Foo Foo Foo`
		if actual != expected {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
	t.Run("To lower", func(t *testing.T) {
		s := `bar BAR Bar bAR BAr`
		actual := All(s, wm)

		expected := `bar bar bar bar bar`
		if actual != expected {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
	t.Run("To all-caps", func(t *testing.T) {
		s := `r2-d2 R2-d2 r2-D2`
		actual := All(s, wm)

		expected := `R2-D2 R2-D2 R2-D2`
		if actual != expected {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
	t.Run("Multiple words", func(t *testing.T) {
		s := `hello world HELLO WORLD hElLo WoRlD HeLlO wOrLd`
		actual := All(s, wm)

		expected := `Hello World Hello World Hello World Hello World`
		if actual != expected {
			reportIncorrectReplacement(t, expected, actual)
		}
	})
	t.Run("With newline", func(t *testing.T) {
		s := `
			hello world
			hello
			world
		`
		actual := All(s, wm)

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
		var wm wordmaps.WordMap
		wm.AddOne("-ize", "-ise")

		source := "They Realize that they should not idealize."
		result := All(source, wm)

		expected := "They Realise that they should not idealise."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("replace only if preceded by another word", func(t *testing.T) {
		var wm wordmaps.WordMap
		wm.AddOne("- dogs", "- cats")

		source := "Dogs are nice and dogs are cool."
		result := All(source, wm)

		expected := "Dogs are nice and cats are cool."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit prefix", func(t *testing.T) {
		var wm wordmaps.WordMap
		wm.AddOne("-phone", "phone")

		source := "That cat has a telephone."
		result := All(source, wm)

		expected := "That cat has a phone."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit the preceding word", func(t *testing.T) {
		var wm wordmaps.WordMap
		wm.AddOne("- people", "people")

		source := "Cool people are nice and nice people are cool."
		result := All(source, wm)

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
		result := All(source, wm)

		expected := "The colours on this colourful painting are amazing."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("replace only if succeeded by another word", func(t *testing.T) {
		var wm wordmaps.WordMap
		wm.AddOne("dog -", "cat -")

		source := "I have a dog and you have a dog."
		result := All(source, wm)

		expected := "I have a cat and you have a dog."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("maintain the succeeding word", func(t *testing.T) {
		var wm wordmaps.WordMap
		wm.AddOne("very -", "super -")

		source := "This is a very special day."
		result := All(source, wm)

		expected := "This is a super special day."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit suffix", func(t *testing.T) {
		var wm wordmaps.WordMap
		wm.AddOne("dog-", "dog")

		source := "I have a dog, but you have a small doggy."
		result := All(source, wm)

		expected := "I have a dog, but you have a small dog."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit the succeeding word", func(t *testing.T) {
		var wm wordmaps.WordMap
		wm.AddOne("a -", "a")

		source := "I have a particularly cool dog."
		result := All(source, wm)

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
		result := All(source, wm)

		expected := "It is a fanfreakingtastic movie."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit prefix, maintain suffix", func(t *testing.T) {
		var wm wordmaps.WordMap
		wm.AddOne("-b-", "b-")

		source := "abc"
		result := All(source, wm)

		expected := "bc"
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("maintain prefix, omit suffix", func(t *testing.T) {
		var wm wordmaps.WordMap
		wm.AddOne("-b-", "-b")

		source := "abc"
		result := All(source, wm)

		expected := "ab"
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("omit both", func(t *testing.T) {
		var wm wordmaps.WordMap
		wm.AddOne("-b-", "b")

		source := "abc"
		result := All(source, wm)

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
	result := All(source, wm)

	expected := "I send them a email. And later another email."
	if result != expected {
		reportIncorrectReplacement(t, expected, result)
	}
}

func TestReplaceWordWithoutSuffixes(t *testing.T) {
	var wm wordmaps.WordMap
	wm.AddOne("commen", "common")

	source := "He game a comment that that is quite commen"
	result := All(source, wm)

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
		result := All(source, wm)

		expected := "This is a foo."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("multiple instances of word", func(t *testing.T) {
		source := "This is a FOOO and this is a fooo as well."
		result := All(source, wm)

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
		result := All(source, wm)

		expected := "This is a foo."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("multiple instances of word", func(t *testing.T) {
		source := "This is a FO and this is a fo as well."
		result := All(source, wm)

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
		result := All(source, wm)

		expected := "lorem ipsum hey\nplanet dolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("newline with indentation", func(t *testing.T) {
		source := "lorem ipsum hello\n  world dolor sit amet."
		result := All(source, wm)

		expected := "lorem ipsum hey\n  planet dolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("space in from but not in to", func(t *testing.T) {
		source := "lorem ipsum foo\nbar dolor sit amet."
		result := All(source, wm)

		expected := "lorem ipsum foobar\ndolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("space in from but not in to, with indentation", func(t *testing.T) {
		source := "lorem ipsum foo\n  bar dolor sit amet."
		result := All(source, wm)

		expected := "lorem ipsum foobar\n  dolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("less spaces in from than in to", func(t *testing.T) {
		source := "lorem ipsum a\ndog dolor sit amet."
		result := All(source, wm)

		expected := "lorem ipsum an\namazing dog dolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("more spaces in from than in to", func(t *testing.T) {
		source := "lorem ipsum hello\nbeautiful world dolor sit amet."
		result := All(source, wm)

		expected := "lorem ipsum hey\nplanet dolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}

		source = "lorem ipsum hello beautiful\nworld dolor sit amet."
		result = All(source, wm)

		expected = "lorem ipsum hey planet\ndolor sit amet."
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceEscapeHyphen(t *testing.T) {
	var wm wordmaps.WordMap
	wm.AddOne(`\-foobar`, `foobar`)
	wm.AddOne(`world\-`, `world!`)

	t.Run("prefix", func(t *testing.T) {
		source := `-foobar`
		result := All(source, wm)

		expected := `foobar`
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("suffix", func(t *testing.T) {
		source := `Hello world-`
		result := All(source, wm)

		expected := `Hello world!`
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}

func TestReplaceEscapeEscapeCharacter(t *testing.T) {
	var wm wordmaps.WordMap
	wm.AddOne(`\\bar`, `bar`)
	wm.AddOne(`foo\\`, `foo`)

	t.Run("prefix", func(t *testing.T) {
		source := `foo \bar`
		result := All(source, wm)

		expected := `foo bar`
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
	t.Run("suffix", func(t *testing.T) {
		source := `foo\ bar`
		result := All(source, wm)

		expected := `foo bar`
		if result != expected {
			reportIncorrectReplacement(t, expected, result)
		}
	})
}
