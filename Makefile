program_main:=./cmd/wordrow
executable_file:=wordrow

default_test_root:=.
coverage_root:=./internal/...
coverage_file:=coverage.out
fuzz_dir:=./_fuzz

markdown_files:=./*.md ./docs/*.md ./.github/**/*.md

go_install:=GO111MODULE=on go get -u


default: help

init: hooks install

hooks:
	@echo SETTING UP GIT HOOKS...
	@cp ./scripts/pre-commit.sh ./.git/hooks/pre-commit

install: install-deps install-dev-deps

install-deps:
	@echo INSTALLING DEPENDENCIES...
	$(go_install) github.com/yargevad/filepathx
	$(go_install) github.com/ericcornelissen/stringsx

install-dev-deps:
	@echo INSTALLING DEVELOPMENT TOOLS...
	$(go_install) golang.org/x/tools/cmd/goimports
	@echo INSTALLING STATIC ANALYSIS TOOLS...
	$(go_install) 4d63.com/gochecknoinits
	$(go_install) gitlab.com/opennota/check/cmd/aligncheck
	$(go_install) gitlab.com/opennota/check/cmd/structcheck
	$(go_install) gitlab.com/opennota/check/cmd/varcheck
	$(go_install) github.com/alexkohler/dogsled/cmd/dogsled
	$(go_install) github.com/alexkohler/nakedret
	$(go_install) github.com/alexkohler/prealloc
	$(go_install) github.com/alexkohler/unimport
	$(go_install) github.com/client9/misspell/cmd/misspell
	$(go_install) github.com/ericcornelissen/goparamcount
	$(go_install) github.com/fzipp/gocyclo/cmd/gocyclo
	$(go_install) github.com/go-critic/go-critic/cmd/gocritic
	$(go_install) github.com/gordonklaus/ineffassign
	$(go_install) github.com/jgautheron/goconst/cmd/goconst
	$(go_install) github.com/kisielk/errcheck
	$(go_install) github.com/kyoh86/looppointer/cmd/looppointer
	$(go_install) github.com/mdempsky/unconvert
	$(go_install) github.com/nishanths/exhaustive/...
	$(go_install) github.com/remyoudompheng/go-misc/deadcode
	$(go_install) github.com/sanposhiho/wastedassign/cmd/wastedassign
	$(go_install) github.com/tommy-muehle/go-mnd/cmd/mnd
	$(go_install) golang.org/x/lint/golint
	$(go_install) golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
	$(go_install) honnef.co/go/tools/cmd/staticcheck
	$(go_install) mvdan.cc/unparam
	@echo INSTALLING MANUAL ANALYSIS TOOLS...
	$(go_install) github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

build:
	go build -o $(executable_file) $(program_main)

build-all:
	GOOS=windows GOARCH=amd64 go build -o $(executable_file)_win-amd64.exe $(program_main)
	GOOS=linux GOARCH=amd64 go build -o $(executable_file)_linux-amd64.o $(program_main)

test%: PKG?=$(default_test_root)
test: test-unit

test-unit:
	go test ${PKG}/...

coverage:
	go test $(coverage_root) -coverprofile $(coverage_file)
	go tool cover -html=$(coverage_file)

fuzz%: FUNC?=Fuzz  # Set default fuzzing function to "Fuzz"
fuzz: fuzz-build fuzz-run

fuzz-build:
	@echo BUILDING FUZZING BINARY FOR ${PKG}...
	@cd ${PKG}; go-fuzz-build

fuzz-run:
	@echo FUZZING ${PKG}::${FUNC}...
	@cd ${PKG}; go-fuzz -func ${FUNC} -workdir ${fuzz_dir}

benchmark:
	go test $(test_root) -bench=. -run=XXX

analysis:
	@echo VETTING...
	@go vet ./...
	@go vet -vettool=`which wastedassign` ./...
	@go vet -vettool=`which shadow` ./...
	@aligncheck ./...
	@dogsled -n 1 -set_exit_status ./...
	@exhaustive ./...
	@goconst -ignore-tests -set-exit-status ./...
	@gocritic check -enableAll ./...
	@gocyclo -over 15 ./
	@goparamcount -set_exit_status -max 3 ./...
	@looppointer ./...
	@mnd -ignored-numbers "0,1" ./...
	@prealloc -set_exit_status ./...
	@staticcheck -show-ignored ./...
	@structcheck -a -e -t ./...
	@unconvert -v ./...
	@unparam -exported -tests ./...
	@varcheck -e ./...

	@echo VERIFYING ERRORS ARE CHECKED...
	@errcheck -asserts -blank -ignoretests -exclude .errcheckrc.txt ./...

	@echo CHECKING FOR DEAD CODE...
	@ineffassign ./*
	@deadcode ./internal/*
	@deadcode ./cmd/*

	@echo CHECKING SPELLING...
	@misspell -error .

format:
	go fmt ./...
	go mod tidy
	goimports -w .

lint: lint-go lint-md

lint-go:
	@echo LINTING GO...
	@golint -set_exit_status ./...
	@gochecknoinits ./...
	@nakedret -l 0 ./...
	@unimport ./...

lint-md:
	@echo LINTING MARKDOWN...
	@npx markdownlint-cli -c .markdownlintrc.yml $(markdown_files)

clean:
	rm -rf $(executable_file)*
	rm -rf $(coverage_file)
	rm `find ./ -name '_fuzz'` -rf
	rm `find ./ -name '*-fuzz.zip'` -rf

.PHONY: default init hooks install build clean format lint analysis test fuzz

help:
	@echo "USAGE:"
	@echo "  make [SUBCOMMANDS] <OPTIONS>"
	@echo
	@echo "SUBCOMMANDS:"
	@echo "  init       Initialze the environment for development."
	@echo "  hooks      Install git hooks."
	@echo "  install    Install all project dependencies"
	@echo "  build      Build the wordrow binary (for the current OS)."
	@echo "  clean      Remove all generated files."
	@echo "  test       Run all tests. Use the option PKG to specify the package"
	@echo "               to test."
	@echo "  coverage   Run all tests and create a coverage report."
	@echo "  fuzz       Run go-fuzz on the source code. Use the option PKG to"
	@echo "               specifythe package to fuzz and option FUNC to specify"
	@echo "               the fuzzing function."
	@echo "  format     Automatically format the source code."
	@echo "  lint       Lint the source code."
	@echo "  analysis   Analyze the source code."
	@echo
	@echo "EXAMPLES:"
	@echo "  make init"
	@echo "  make build"
	@echo "  make analysis"
	@echo "  make test PKG=./internal"
	@echo "  make fuzz PKG=./internal/replace FUNC=FuzzReplaceAll"

