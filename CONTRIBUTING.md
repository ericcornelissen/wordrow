# Contributing Guidelines

The _wordrow_ project welcomes contributions and corrections of all forms. This
includes improvements to the documentation or code base, new tests, bug fixes,
and implementations of new features. We recommend opening an issue before making
any significant changes so you can be sure your work won't be rejected. But for
changes such as fixing a typo you can open a Pull Request directly.

Before you continue, please do make sure to read through the relevant sections
of this document. In this document you can read about:

- [Bug Reports](#bug-reports)
- [Feature Requests](#feature-requests)
- [Workflow](#workflow)
- [Project Setup](#project-setup)
  - [Prerequisites](#prerequisites)
  - [Commands](#commands)
  - [Git Hooks](#git-hooks)
- [Testing](#testing)
  - [Fuzzing](#fuzzing)

---

## Bug Reports

If you have problems with _wordrow_ or think you've found a bug, please report
it to the developers; we cannot promise to do anything but we might well want to
fix it.

Before reporting a bug, make sure you've actually found a real bug. Carefully
read the documentation and see if it really says you can do what you're trying
to do. If it's not clear whether you should be able to do something or not,
report that too; it's a bug in the documentation!

Before reporting a bug, try to isolate it to a small working example that
reproduces the problem. Then, create a bug report including this example and the
results, including any error or warning messages. Please don't paraphrase these
messages: it's best to copy and paste them into your report. Finally, be sure to
explain what you expected to happen; this will help us decide whether it is a
bug or a problem with the documentation.

In addition to the information above, please be careful to include the version
number of _wordrow_ you are using. You can get this information with the command
`wordrow --version`. Be sure also to include the type of machine and operating
system you are using.

Once you have a precise problem you can report it online as a [Bug Report](https://github.com/ericcornelissen/wordrow/issues/new?labels=bug&template=bug_report.md).

## Feature Requests

New features can be requested through GitHub issues. Please provide as much
information as possible to help us understand the feature. This may include:

- Both a summary and detailed explanation of the feature.
- An example use case of the feature.
- A list of potential drawbacks of the feature.
- An outline for an implementation of the feature.
- Possible test cases for an implementation.

---

## Workflow

If you decide to make a contribution, please do use the following workflow:

- Fork the repository.
- Create a new branch from the latest `master`.
- Make your changes on the new branch.
- Commit to the new branch and push the commit(s).
- Make a Pull Request.

## Project Setup

This project is build for version `1.13` of Go and uses [GNU Make] as build
tool. In addition [golint] and [markdownlint] are used to lint the source files
and [gofmt] is used to format source files.

### Prerequisites

The prerequisites for contributing to this project are:

- Go; version `1.13`
- Git
- [GNU Make] (_Windows users can use [Make by GNUWin32]_)
- [golint]
- [gosec]
- [go-fuzz] (_only for fuzzing_)
- [NodeJS]; version `>=10` (_only needed for [markdownlint]_)

### Commands

The table below shows an overview of the commands available for the development
of this project. Note that the table is (intentionally) incomplete.

| Command          | Description                                          |
| ---------------- | ---------------------------------------------------- |
| `make`           | Compile a binary for the current OS called `wordrow` |
| `make format`    | Format the source files of the project               |
| `make test`      | Run all test suites for the project                  |
| `make coverage`  | Run all test suites and show the coverage results    |
| `make lint`      | Lint the source files of the project                 |
| `make analysis`  | Run static analysis tools on the code base           |
| `make build-all` | Compile all target binaries for the project          |
| `make clean`     | Delete all generated files                           |

### Git Hooks

We recommend [setting up a Git hook](https://githooks.com) on commits or pushes
to make sure your changes don't include any accidental mistakes. You can use the
following shell script as a template.

```shell
#!/bin/sh

# Stash unstaged changes
git stash -q --keep-index

# See if the project can be build and all tests pass.
make
make test

# Format the codebase and include (relevant) formatting changes in the commit.
make format
git update-index --again

# Check other formatting things. You can use `make lint-go` if you don't have
# NodeJS on your system.
make lint

# Restore unstaged changes
git stash pop -q
```

---

## Testing

In this project tests are written using Go's standard `testing` package. Tests
are located in files with a `_test.go` suffix and test function names always
start with the prefix `Test`. Helper methods for tests go into a file named
`helpers_test.go` in the same package.

All changes should be tested. New features should be tested thoroughly to ensure
they work correctly. Bug fixes should include at least one test case to verify
the bug has been fixed.

Public functions may have [testable examples]. The name of a testable example
function start with the prefix `Example` instead of `Test`, followed by the name
of the function of which the test is an example. All testable examples should
appear in the `_test.go` file before any _normal_ test.

_Below are guidelines regarding more advanced testing topics._

### Fuzzing

This project provides tooling to fuzz the code base using [go-fuzz]. You can run
the following commands to fuzz the packages of the project. (Any artefact that
may be generated by [go-fuzz] is automatically ignored by Git.)

```shell
# To fuzz the code base you must specify the target package using "PKG".
$ make fuzz PKG=internal/cli

# If there are multiple fuzzing functions for the package you must specify the
# function name (excluding the "Fuzz" prefix) using "FUNC".
$ make fuzz PKG=internal/wordmaps FUNC=MarkDown
```

In this project, fuzzing code must be located in a file with the `_fuzz.go`
suffix (similar to `_test.go` for tests) and fuzzing functions must be named
`FuzzXXX` (similar to `TestXXX` for tests). Use the following template for new
fuzzing files.

```go
// +build gofuzz

package xxx

func Fuzz(data []byte) int {
	// 1. Set the package name
	// 2. (optional) Rename this function to FuzzXXX
	// 3. Define the fuzzing logic
	// 4. Fuzz!
	return 0
}

```

You are welcome to use existing fuzzing functions to discover bugs. You can also
contribute by adding new fuzzing functions for previously unfuzzed code, or by
improving any of the existing fuzzing functions. If you discover a bug while
fuzzing, please follow the [Bug Reports](#bug-reports) steps.

[go-fuzz]: https://github.com/dvyukov/go-fuzz
[gofmt]: https://golang.org/cmd/gofmt/
[golint]: https://github.com/golang/lint
[gosec]: https://github.com/securego/gosec
[GNU Make]: https://www.gnu.org/software/make/
[Make by GNUWin32]: http://gnuwin32.sourceforge.net/packages/make.htm
[markdownlint]: https://github.com/DavidAnson/markdownlint
[NodeJS]: https://nodejs.org/en/
[testable examples]: https://blog.golang.org/examples
