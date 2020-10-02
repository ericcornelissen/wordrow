# wordrow

[![GitHub Actions][ci-image]][ci-url]
[![Test Coverage][coverage-image]][coverage-url]
[![Maintainability][maintainability-image]][maintainability-url]
[![LGTM Alerts][lgtm-image]][lgtm-url]
[![BCH Compliance][bch-image]][bch-url]
[![Go Report Card][grc-image]][grc-url]

A small CLI tool to replace instances of words with other words in plaintext,
written in [Go].

**Quick Links**:
  [Documentation] |
  [Changelog] |
  [Contributing Guidelines] |
  [Code of Conduct]

## Why

Everything *wordrow* does can be achieved with existing tools, e.g. with [sed],
simple scripts, or even just a text editor. However, *wordrow* aims to make it
easy to do text replacements in batch and continuously. Using a simple syntax
to define mappings, you can improve what you writing quickly and continuously.

## Usage

First, download the [latest release] for your system and add the *wordrow*
binary to your PATH.

To get started, you can read [how to use the CLI](./docs/cli.md) and [how to
define a mapping](./docs/mapping-files.md). For quick access to help on the CLI
use:

```shell
$ wordrow --help
```

[changelog]: ./CHANGELOG.md
[code of conduct]: ./CODE_OF_CONDUCT.md
[contributing guidelines]: ./CONTRIBUTING.md
[documentation]: ./docs
[go]: https://golang.org/
[latest release]: https://github.com/ericcornelissen/wordrow/releases/latest
[sed]: https://www.gnu.org/software/sed/manual/sed.html

[ci-url]: https://github.com/ericcornelissen/wordrow/actions?query=workflow%3A%22wordrow+CI%22+branch%3Amaster
[ci-image]: https://github.com/ericcornelissen/wordrow/workflows/wordrow%20CI/badge.svg?branch=master
[coverage-url]: https://codeclimate.com/github/ericcornelissen/wordrow/test_coverage
[coverage-image]: https://api.codeclimate.com/v1/badges/36d32594ea2274cbf972/test_coverage
[maintainability-url]: https://codeclimate.com/github/ericcornelissen/wordrow/maintainability
[maintainability-image]: https://api.codeclimate.com/v1/badges/36d32594ea2274cbf972/maintainability
[lgtm-url]: https://lgtm.com/projects/g/ericcornelissen/wordrow/alerts/
[lgtm-image]: https://img.shields.io/lgtm/alerts/g/ericcornelissen/wordrow.svg?logo=lgtm&logoWidth=18
[bch-url]: https://bettercodehub.com/results/ericcornelissen/wordrow
[bch-image]: https://bettercodehub.com/edge/badge/ericcornelissen/wordrow?branch=master
[grc-url]: https://goreportcard.com/report/github.com/ericcornelissen/wordrow
[grc-image]: https://goreportcard.com/badge/github.com/ericcornelissen/wordrow
