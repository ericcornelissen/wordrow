# The wordrow CLI

The Command Line Interface (CLI) of *wordrow* can be used to replace words in
files, large or small, using mappings that are simple to define. This document
gives an introduction on how to use the CLI of *wordrow*.

This document is not an in-depth documentation of all CLI options, instead it
provides an introduction to how it can be used to achieve common tasks. You can
always run the following command to get documentation on all the options.

```shell
$ wordrow --help
```

In this document you can read about:

- [The Basics](#the-basics)
- [Converting Multiple Files](#converting-multiple-files)
- [Inverting a Mapping File](#inverting-a-mapping-file)
- [Controlling the Output](#controlling-the-output)

## The Basics

If you have a file called `input.txt` that contains some plaintext, for example:

```text
# input.txt

I have a dog named Quark, a horse named
Proton, and a canary named Atom.
```

Then you can use *wordrow* to automatically and quickly replace certain words in
that text by other words. The quickest way to do this is by specifying a word
that you want to change on the CLI. To do this you can use the `--map` option as
follows:

```shell
$ wordow input.txt --map dog,cat
```

Then, `input.txt` will be updated as follows:

```diff
- I have a dog named Quark, a horse named
- Proton, and a canary named Atom.
+ I have a cat named Quark, a horse named
+ Proton, and a canary named Atom.
```

Another way is to define a [mapping file] in the Comma Separated Values (CSV)
format. With a mapping file you can specify any number of mappings. For example,
to change all the animals in the text you can create a file named `animals.csv`.

```csv
# animals.csv

dog, cat
canary, parrot
horse, donkey
```

Now you can run *wordrow* on `input.txt` with `mapping.csv` as shown below. This
tells *wordrow* that you want to replace words in `input.txt` using the
`--map-file` file `animals.csv`

```shell
$ wordow input.txt --map-file animals.csv
```

Then, `input.txt` will be updated as follows:

```diff
- I have a dog named Quark, a horse named
- Proton, and a canary named Atom.
+ I have a cat named Quark, a donkey named
+ Proton, and a parrot named Atom.
```

## Converting Multiple Files

In a typically scenario, you may want to run *wordrow* on multiple files. To run
*wordrow* on multiple files you can either specify multiple files or use a
[glob]. Specifying multiple input files is straightforward. For example, using
our mapping file from before, we can run *wordrow* on two input files using:

```shell
$ wordrow input-1.txt input-2.txt --map-file animals.csv
```

If, instead, you want to run *wordrow* on all `.txt` files in the current
folder, you can use a [glob]. A glob is a special expression that allows you
to refer to a bunch of files in a single argument. For example:

```shell
$ wordrow ./*.txt --map-file animals.csv
```

## Inverting a Mapping File

It may happen that you have a (large) mapping file that, instead of using it
left-to-right (i.e. replace _"dog"_ by _"cat"_, in `animals.csv`), you want to
use right-to-left (i.e. replace _"cat"_ by _"dog"_, in `animals.csv`). This can
be achieved simply by adding the `--invert` argument when you run *wordrow*.

For example, if you have an input file `input-alt.txt` that looks like:

```text
# input-alt.txt

I have a cat, a donkey, and a parrot.
```

And you run *wordrow* using the `--invert` argument as:

```shell
$ wordow input-alt.txt --map-file animals.csv --invert
```

Then, `input-alt.txt` will be updated as follows:

```diff
- I have a cat, a donkey, and a parrot.
+ I have a dog, a horse, and a canary.
```

## Controlling the Output

You may control the output behaviour of the CLI through some flag. First, you
can prevent *wordrow* from making any changes to the input files by using the
dry run flag:

```shell
$ wordrow input.txt --map-file animals.csv --dry-run
```

Second, if you want to control how much is printed to the console when you run
*wordrow*, you can use the following flags to reduce or increase the amount of
logging:

```shell
# Less printing
$ wordrow input.txt --map-file animals.csv --silent

# More printing
$ wordrow input.txt --map-file animals.csv --verbose
```

[glob]: https://mincong.io/2019/04/16/glob-expression-understanding/
[mapping file]: ./mapping-files.md
