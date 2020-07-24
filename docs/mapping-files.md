# Mapping Files

You can use *wordrow* to replace words in files, but to do that you need to
define a mapping file. This document gives an introduction to defining your own
mapping file. All examples in this document are specified in the _Comma
Separated Values_ (CSV) format. Note that *wordrow* supports other [mapping
formats] as well.

This document does not contain any mappings that are particularly useful in a
real-world scenario. Instead, it illustrate how you can use *wordrow* through
made-up examples.

In this document you can read about:

- [The Basics](#the-basics)
  - [Direction](#direction)
  - [Whitespace](#whitespace)
  - [Phrases](#phrases)
  - [Capitalisation](#capitalisation)
  - [Many-to-One](#many-to-one)
- [Prefixes and Suffixes](#prefixes-and-suffixes)
  - [The Preceding and Succeeding Word](#the-preceding-and-succeeding-word)
  - [Omitting Prefixes or Suffixes](#omitting-prefixes-or-suffixes)
  - [Escaping a Prefix or Suffix Dash](#escaping-a-prefix-or-suffix-dash)
- [Order Matters](#order-matters)
  - [Using Ordering to Your Advantage](#using-ordering-to-your-advantage)
- [Notes](#notes)

## The Basics

To define a mapping from one word to another word for *wordrow*, you must create
a file containing the words you want to map. For example, to replace the word
_"dog"_ to _"cat"_ you can write the following CSV file.

```csv
# mapping.csv

dog,cat
```

This mapping will replace all instances of the word _"dog"_ in a text, provided
it does not have a prefix or suffix, with the word _"cat"_. The following sample
text illustrates this behaviour. Notice that _"dog"_ is replaced by _"cat"_
because it has no prefix or suffix, whereas _"doggy"_ is not replaced by _"cat"_
(or _"catgy"_).

```diff
- I have a dog, but you have a small doggy.
+ I have a cat, but you have a small doggy.
```

A mapping file my contain any number of mappings. For example, when using the
following mapping file the words _"dog"_, _"canary"_, and _"horse"_ will be
replaced by _"cat"_, _"parrot"_, and _"donkey"_ respectively.

```csv
# mapping.csv

dog,cat
canary,parrot
horse,donkey
```

### Direction

As you have seen, a mapping will be interpreted left-to-right. That is, the
value on the left will be replaced by the value on the right in the input text.
So, if you use the following mapping file:

```csv
# mapping.csv

dog,cat
```

Then, for an input text containing both the word _"dog"_ and the word _"cat"_,
only the word _"dog"_ is replaced by _"cat"_. For example:

```diff
- I have a dog and you have a cat.
+ I have a cat and you have a cat.
```

A mapping can be inverted using the [*wordrow* CLI] as shown here.

```shell
$ wordrow input.txt --map-file mapping.csv --invert
```

Then, in the example before, only the word _"cat"_ will be replaced by _"dog"_:

```diff
- I have a dog and you have a cat.
+ I have a dog and you have a dog.
```

### Whitespace

Any whitespace before the first character and any whitespace after the last
character is ignored. Consider a CSV file with some whitespace, as in the
following example (dots are used to illustrate where the whitespace is).

```csv
..dog..,..cat..
```

This will still replace the word _"dog"_ by _"cat"_ in our text, even though
_"dog"_ is not preceded or followed by two spaces in this example. Also, it does
not add the two spaces surrounding _"cat"_ to the output.

```diff
- I have a dog, but you have a small doggy.
+ I have a cat, but you have a small doggy.
```

### Phrases

On the other hand, whitespace within a mapping value is not ignored. So, you can
replace a group of words, a phrase, in one mapping quite easily. For example, to
replace the phrase _"a dog"_  you can define the following mapping.

```csv
a dog, an amazing dog
```

This will replace the phrase _"a dog"_ in the text by _"an amazing dog"_. Also
in this scenario _"doggy"_ is not changed, as it does not match _"a dog"_.

```diff
- I have a dog, but you have a small doggy.
+ I have an amazing dog, but you have a small doggy.
```

There is no limitation on the number of words in a mapping phrase.

Every kind of whitespace (space, tab, newline) in the input text is considered
to match the spaces in the mapping file. This means that if the phrase that you
want to replaces appears at a line break, _wordrow_ will still replace it. For
example, given the mapping:

```csv
a hippo, an elephant
```

A text where _"a hippo"_ appears at a line break will be replaced with _"an
elephant"_ while preserving the line break.

```diff
- There once was a
- hippo in town.
+ There once was an
+ elephant in town.
```

### Capitalisation

The capitalisation present in a mapping is generally ignored, except when the
textual value (ignoring casing) is equal before and after.

#### Context-aware Capitalisation

For most mappings, _wordrow_ will try to apply the capitalisation of the words
as they appear in the text is maintained. Capital letters are maintained at the
start of a word and also if the original word appears in all capitals in the
text.

For example, if you have a mapping to change _"dog"_ into _"horse"_, the
capitalisation will be maintained as follows.

```diff
- Dog dog DOG
+ Horse horse HORSE
```

If a mapping consists of multiple words, the capitalisation of each individual
word is maintained. This also goes for, e.g., hyphenated words. For example, if
you use the following mapping file:

```csv
hello world, hey planet
so called, so-called
```

Then, a text containing _"So Called"_ or _"Hello World"_ will be updated to use
_"So-Called"_ and _"Hey Planet"_ with identical capitalisation:

```diff
- A So Called "Hello World" program is a program that prints "Hello world!".
+ A So-Called "Hey Planet" program is a program that prints "Hey planet!".
```

#### Explicit Capitalisation Mapping

If the mapping you define does not change the textual value of the phrase (i.e.
ignoring casing the values of the mapping are the same), then the provided
capitalisation is always used. This can be used for example to enforce the
correct capitalisation of proper names.

For example, if you have a mapping to change the name _"max"_ into _"Max"_, the
capitalisation will be maintained as follows.

```diff
- My dog is called max. Max is an awesome dog.
+ My dog is called Max. Max is an awesome dog.
```

### Many-to-One

In some cases you may want to replace multiple words by the same word. Instead
of defining a mapping for each word individually, you can define all of them in
a single mapping definition. Simply add all the values that should map to one
specific word on a single line and they will all be replaced by the last word in
the definition. For example:

```csv
cat, dog, horse
```

This will replace both the word _"dog"_ and _"cat"_ in a text with the word
_"horse"_.

```diff
- A cat is an animal and a dog is a mammal.
+ A horse is an animal and a horse is a mammal.
```

---

## Prefixes and Suffixes

You can define more advanced mappings by replacing words including a prefix, a
suffix or both. To do this, you can add a dash (`-`) before (prefix) or after
(suffix) the words in your mappings. If you do this, the word will be replaced
if it appears as is, or with a prefix/suffix in the text.

For example, if you want to replace all words ending in _"ize"_ with the same
word ending in _"ise"_ you can define the following mapping.

```csv
-ize, -ise
```

Then, if you use this mapping on a text containing words ending in _"ize"_, all
of them will be replaced by the same word, but ending in _"ise"_. For example:

```diff
- They realize that they should not idealize.
+ They realise that they should not idealise.
```

Similarly, if you want to replace the word _"color"_, and all its variants, with
the word _"colour"_, and all its variants, you can define the following mapping.

```csv
color-, colour-
```

Then, if you use this mapping on a text containing the word _"color"_ and also
words starting with _"color"_, all of them will be replaced by _"colour-"_. For
example:

```diff
- The colors on this colorful painting are amazing.
+ The colours on this colourful painting are amazing.
```

Note that it is not required for a word specified with a prefix or suffix to
appear with a prefix or suffix in the text. For example, you can use the
`color-` to `colour-` mapping to replace the word _"color"_ by itself as well,
for example:

```diff
- What color is the dog?
+ What colour is the dog?
```

It is also possible to match both prefixes and suffixes in the same mapping. To
do this, simply add both the prefix and suffix dash to the words in the
mapping.

```csv
-bloody-, -freaking-
```

In this example, if _"bloody"_ is used as an [expletive infixation] it will be
replaced by _"freaking"_.

```diff
- It is a fanbloodytastic movie.
+ It is a fanfreakingtastic movie.
````

It is important to remember that dashes only have a special meaning at the start
and end of mapping values. You can always use dashes in the middle of words. For
example, you can define the following mapping.

```csv
dog-like, cat-like
```

That will replace all instances of _"dog-like"_ in a text by _"cat-like"_. But
it won't affect any instances of _"dog"_ and _"like"_ with something in between
them, as in:

```diff
- I have a dog-like cat, and you have a dog I like.
+ I have a cat-like cat, and you have a dog I like.
```

### Omitting Prefixes or Suffixes

It is necessary to write the dash in the both words of the mapping. Otherwise
the prefix or suffix will be omitted when the word is replaced. This can,
however, be used if you want to restyle your text. Consider the following
mapping.

```csv
dog-, dog
```

In this example, any suffix the word _"dog"_ has will be removed and replaced by
just the word _"dog"_. So, a text that uses _"doggy"_ will be updated to use
_"dog"_ instead.

```diff
- I have a dog, but you have a small doggy.
+ I have a dog, but you have a small dog.
```

### The Preceding and Succeeding Word

One possible way to use the prefix and suffix dash is to match instances of the
word only if there is another word before or after it. You can do this by
putting a space between the word and the dash (remember that [whitespace
matters]).

```csv
- dogs, - cats
```

In this example, the word _"dog"_ is only replaced by _"cat"_ if there is a word
before _"dog"_, as illustrated by this text:

```diff
- Dogs are nice and dogs are cool.
+ Dogs are nice and cats are cool.
```

Again, it is necessary to specify the dash in both words. Otherwise the word
before or after the matched word is omitted from the result.

### Escaping a Prefix or Suffix Dash

Occasionally, you may need to replace a dash at the start or end of a word. In
this scenario you can escape the dash using a backslash (`\`).

```csv
world\-, world!
```

Given this mapping, any instance of the string _"world-"_ will be replaced by
_"world!"_, but words like _"worlds"_ will not not changed.

```diff
- Hello world- What is life like on other worlds?
+ Hello world! What is life like on other worlds?
```

---

## Order matters

It is important to note that the ordering in a mapping file matters. The
mappings defined in a file are applied to the input text top to bottom. To
illustrate this, consider the following mapping file.

```csv
dog, cat
cat, dog
```

Using this mapping on our example you will find that, in the end, it doesn't
change the input text.

```diff
#  1. applying "dog, cat"
- I have a dog, but you have a small doggy.
+ I have a cat, but you have a small doggy.

#  2. applying "cat, dog"
- I have a cat, but you have a small doggy.
+ I have a dog, but you have a small doggy.
```

This may be intuitive. However, let's now consider what happens if you swap the
two lines in the mapping file.

```csv
cat, dog
dog, cat
```

With this mapping, the input text is does change!

```diff
#  1. applying "cat, dog"
- I have a dog, but you have a small doggy.
+ I have a dog, but you have a small doggy.

#  2. applying "dog, cat"
- I have a dog, but you have a small doggy.
+ I have a cat, but you have a small doggy.
```

So, keep this in mind when you define a mapping to avoid any problems.

### Using Ordering to Your Advantage

This effect can also be used to your advantage. Consider the following scenario:
you define a mapping from a word that uses the article "a" to a word that uses
the article "an". For example, let's say you want to replace _"duck"_ with
_"owl"_ and you define the following mapping.

```csv
duck, owl
a owl, an owl
```

Then, a text containing the phrase _"a duck"_ will be transformed as follows.

```diff
#  1. applying "duck, owl"
- I see a duck, is it your duck?
+ I see a owl, is it your owl?

#  2. applying "a owl, an owl"
- I see a owl, is it your owl?
+ I see an owl, is it your owl?
```

## Notes

1. A mapping containing characters that are not in the [UTF-8 character set]
   won't be processed.

[expletive infixation]: https://www.youtube.com/watch?v=dt22yWYX64w
[mapping formats]: ./mapping-formats.md
[UTF-8 character set]: https://en.wikipedia.org/wiki/UTF-8
[whitespace matters]: #whitespace
[*wordrow* CLI]: ./cli.md
