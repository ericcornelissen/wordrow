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
  - [Multiple Words](#multiple-words)
- [Order Matters](#order-matters)
  - [Using Ordering to Your Advantage](#using-ordering-to-your-advantage)

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
$ wordrow input.txt --mapping mapping.csv --invert
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

### Multiple Words

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

Do be aware that the amount of whitespace between words in a phrase matters. For
example, if you define the mapping with two spaces between _"a"_ an _"dog"_ as
in the following example (dots are used to illustrate where the whitespace is).

```csv
a..dog,an.amazing.dog
```

Then, the example text won't be changed, as it does not contain _"a..dog"_.

```diff
- I have a dog, but you have a small doggy.
+ I have a dog, but you have a small doggy.
```

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

[mapping formats]: ./mapping-formats.md
[*wordrow* CLI]: ./cli.md
