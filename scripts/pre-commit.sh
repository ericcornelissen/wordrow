#!/bin/sh

set -e

# Stash unstaged changes
git stash -q -u --keep-index

# See if the project can be build and tests pass.
make
make test

# Format the code base and include (relevant) formatting changes in the commit.
make format
git update-index --again

# Check other formatting things. You can use `make lint-go` if you don't have
# NodeJS on your system.
make lint

# Restore unstaged changes
git stash pop -q
