# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to [Semantic
Versioning].

## [Unreleased]

- _No changes yet_

## [0.7.0-beta] - 2020-10-23

### Features

- Run on STDIN. ([#119])
- Process input files in parallel. ([#132])
- Log when a mapping is overwritten. ([#134])
- Exit with a non-zero exit code on errors. ([#139])

### Bug Fixes

- Allow for prefix notation in inline mappings. ([#136])
- Prevent runtime errors due to invalid mappings. ([#138])

### Documentation

- Provide a list of ready-to-use mapping files. ([#141])

## [0.6.0-beta] - 2020-08-18

### Breaking Changes

- Continue on non-blocking errors. ([#112])

### Features

- Add support for CLI arguments with "=". ([#102])
- Add support for many-to-one mapping syntax on the CLI. ([#118])
- Add `--strict`/`-S` flag to stop for any error. ([#112])

### Bug Fixes

- Fix bug preventing the usage of map files without file extension. ([#101])
- Fix logs of incorrect level being logged. ([#107], [#108])
- Fix crash due to empty values in the `--map` option. ([#110])
- Fix mistakes in the program help message. ([#114], [#115], [#116])

### Miscellaneous

- Removed (unused) `--config` option. ([#107])

## [0.5.0-beta] - 2020-07-24

### Features

- Add functionality to map many words to one word. ([#65])
- Add support for multiple flags in one CLI argument. ([#72])
- Allow mappings that change capitalization. ([#98])

### Bug Fixes

- Fix double newline in program logs. ([#78])
- Fix bug for mappings with Regular Expression syntax. ([#96])
- Avoid crash due to non UTF-8 characters in mappings. ([#96])

### Security

- Clean provided file paths before opening. ([#99])

### Documentation

- Update use of outdated CLI argument name. ([#67])
- Update use of outdated term in CLI help message. ([#71])

## [0.4.0-beta] - 2020-05-27

### Features

- Add option to explicitly state map file format. ([#46], [#57])
- Add support for more extensions as MarkDown map file. ([#50])
- Allow file extension of map files to be upper or mixed case. ([#55])
- Allow for replacing hyphens through escaping with backslash. ([#61])

### Bug Fixes

- Fix program crash due to empty CLI argument. ([#42])
- Fix program crash because of an incomplete MarkDown file. ([#58])

## [0.3.0-beta] - 2020-05-03

### Breaking Changes

- Rename `--map` and `-m` to `--map-file` and `-M` respectively. ([#30])
- Hard matching of individual whitespace characters no longer happens. ([#37])

### Features

- Maintain capitalization of all words in a mapping phrase. ([#26])
- Add option to specify a mapping from the CLI. ([#31])
- Replace phrases with spaces if a match is found with a newline. ([#37])

### Bug Fixes

- Add missing options `--invert` and `--verbose` to the usage message. ([#30])

## [0.2.0-beta] - 2020-04-18

### Features

- Add support for globs in arguments. ([#9], [#29])
- Implement the `--silent` flag, which reduced program output. ([#21])
- Implement the `--verbose` flag, which increases program output. ([#21])
- Add support for prefix and suffix matching. ([#22])

### Bug Fixes

- Fix issues due to empty values in mapping files. ([#14])

### Performance

- Improve performance of word replacement. ([#23])

## [0.1.0-beta] - 2020-04-04

### Features

- Replace instances of one word with another in multiple plaintext files.
- Define mappings of words in CSV or MarkDown files.
- Invert a mapping as it is defined in a file.

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html
[#9]: https://github.com/ericcornelissen/wordrow/pull/9
[#14]: https://github.com/ericcornelissen/wordrow/pull/14
[#22]: https://github.com/ericcornelissen/wordrow/pull/22
[#21]: https://github.com/ericcornelissen/wordrow/pull/21
[#23]: https://github.com/ericcornelissen/wordrow/pull/23
[#26]: https://github.com/ericcornelissen/wordrow/pull/26
[#29]: https://github.com/ericcornelissen/wordrow/pull/29
[#30]: https://github.com/ericcornelissen/wordrow/pull/30
[#31]: https://github.com/ericcornelissen/wordrow/pull/31
[#37]: https://github.com/ericcornelissen/wordrow/pull/37
[#42]: https://github.com/ericcornelissen/wordrow/pull/42
[#46]: https://github.com/ericcornelissen/wordrow/pull/46
[#50]: https://github.com/ericcornelissen/wordrow/pull/50
[#55]: https://github.com/ericcornelissen/wordrow/pull/55
[#57]: https://github.com/ericcornelissen/wordrow/pull/57
[#58]: https://github.com/ericcornelissen/wordrow/pull/58
[#61]: https://github.com/ericcornelissen/wordrow/pull/61
[#65]: https://github.com/ericcornelissen/wordrow/pull/65
[#67]: https://github.com/ericcornelissen/wordrow/pull/67
[#71]: https://github.com/ericcornelissen/wordrow/pull/71
[#72]: https://github.com/ericcornelissen/wordrow/pull/72
[#78]: https://github.com/ericcornelissen/wordrow/pull/78
[#96]: https://github.com/ericcornelissen/wordrow/pull/96
[#98]: https://github.com/ericcornelissen/wordrow/pull/98
[#99]: https://github.com/ericcornelissen/wordrow/pull/99
[#101]: https://github.com/ericcornelissen/wordrow/pull/101
[#102]: https://github.com/ericcornelissen/wordrow/pull/102
[#107]: https://github.com/ericcornelissen/wordrow/pull/107
[#108]: https://github.com/ericcornelissen/wordrow/pull/108
[#110]: https://github.com/ericcornelissen/wordrow/pull/110
[#112]: https://github.com/ericcornelissen/wordrow/pull/112
[#114]: https://github.com/ericcornelissen/wordrow/pull/114
[#115]: https://github.com/ericcornelissen/wordrow/pull/115
[#116]: https://github.com/ericcornelissen/wordrow/pull/116
[#118]: https://github.com/ericcornelissen/wordrow/pull/118
[#119]: https://github.com/ericcornelissen/wordrow/pull/119
[#132]: https://github.com/ericcornelissen/wordrow/pull/132
[#134]: https://github.com/ericcornelissen/wordrow/pull/134
[#136]: https://github.com/ericcornelissen/wordrow/pull/136
[#138]: https://github.com/ericcornelissen/wordrow/pull/138
[#139]: https://github.com/ericcornelissen/wordrow/pull/139
[#141]: https://github.com/ericcornelissen/wordrow/pull/141
