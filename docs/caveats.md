# Caveats

1. Mappings can be overwritten. The last definition of a mapping encountered by
   *wordrow* is the one that will be used. When using the `--verbose` flag you
   will get a warning when a mapping is overwritten.
1. A mapping containing characters that are not in the [UTF-8 character set]
   won't be processed.

[UTF-8 character set]: https://en.wikipedia.org/wiki/UTF-8
