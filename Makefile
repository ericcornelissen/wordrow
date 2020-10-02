program_main:=./cmd/wordrow
executable_file:=wordrow

unit_test_root:=./internal/...
integration_test_root:=./cmd/wordrow/...
coverage_file:=coverage.out
fuzz_dir:=./_fuzz

markdown_files:=./*.md ./docs/*.md ./.github/**/*.md

go_install:=GO111MODULE=on go get -u


default: build

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
	$(go_install) github.com/ericcornelissen/goparamcount
	$(go_install) github.com/fzipp/gocyclo
	$(go_install) github.com/go-critic/go-critic/cmd/gocritic
	$(go_install) github.com/gordonklaus/ineffassign
	$(go_install) github.com/jgautheron/goconst/cmd/goconst
	$(go_install) github.com/kisielk/errcheck
	$(go_install) github.com/kyoh86/looppointer/cmd/looppointer
	$(go_install) github.com/mdempsky/unconvert
	$(go_install) github.com/nishanths/exhaustive/...
	$(go_install) github.com/remyoudompheng/go-misc/deadcode
	$(go_install) github.com/tommy-muehle/go-mnd/cmd/mnd
	$(go_install) golang.org/x/lint/golint
	$(go_install) honnef.co/go/tools/cmd/staticcheck
	$(go_install) mvdan.cc/unparam
	@echo INSTALLING MANUAL ANALYSIS TOOLS...
	$(go_install) github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

build:
	go build -o $(executable_file) $(program_main)

build-all:
	GOOS=windows GOARCH=amd64 go build -o $(executable_file)_win-amd64.exe $(program_main)
	GOOS=linux GOARCH=amd64 go build -o $(executable_file)_linux-amd64.o $(program_main)

test: test-unit test-integration

test-unit:
	go test $(unit_test_root)

test-integration:
	go test $(integration_test_root)

coverage:
	go test $(unit_test_root) -coverprofile $(coverage_file)
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
	@aligncheck ./...
	@dogsled -n 1 -set_exit_status ./...
	@exhaustive -maps ./...
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
