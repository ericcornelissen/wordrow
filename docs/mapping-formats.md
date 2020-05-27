# Mapping Formats

A mapping file for *wordrow* can be defined in different formats. This document
contains the specifications for each available format.

This document covers the following format:

- [Comma Separated Values (CSV)](#comma-separated-values)
- [MarkDown](#markdown)

## Comma Separated Values

A Comma Separated Values (CSV) files can be used to define simple mappings. A
valid CSV for *wordrow* consists of two columns. The left column contains the
words in your text and the right column contains the respective replacements.
For example:

```csv
dog, cat
canary, parrot
horse, donkey
```

If any row contains more than two columns, the entire file is considered invalid
and will not be used by *wordrow*.

Any file with one of the following extension is considered to be a CSV file by
*wordrow*: `.csv`

## MarkDown

A MarkDown file can be used to define a mapping through two-column MarkDown
tables. Every table in a MarkDown file is considered to be a mapping. The table
header is ignored by *wordrow*, but it must be present as the MarkDown is
invalid otherwise. All other lines in a MarkDown file are ignored (you may use
it as comments). For example:

```markdown
# My mapping

This text, as well as the title, will be ignored by wordrow.

| From   | To     |
| ------ | ------ |
| dog    | cat    |
| canary | parrot |
| horse  | donkey |
```

If any row in any table contains more than two columns, the entire file is
considered invalid and will not be used by *wordrow*.

Any file with one of the following extension is considered to be a MarkDown
file by *wordrow*: `.md`
