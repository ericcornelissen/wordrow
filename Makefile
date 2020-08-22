program_main:=./cmd/wordrow
executable_file:=wordrow

unit_test_root:=./internal/...
integration_test_root:=./cmd/wordrow/...
coverage_file:=coverage.out
fuzz_dir:=./_fuzz

markdown_files:=./*.md ./docs/*.md ./.github/**/*.md

go_install:=go get -u
go_install_dev:=GO111MODULE=off $(go_install)


default: build

install: install-deps install-dev-deps

install-deps:
	@echo "INSTALLLING DEPENDENCIES"
	$(go_install) github.com/yargevad/filepathx

install-dev-deps:
	@echo "INSTALLLING DEVELOPMENT TOOLS"
	$(go_install_dev) golang.org/x/tools/cmd/goimports
	@echo "INSTALLLING STATIC ANALYSIS TOOLS"
	$(go_install_dev) golang.org/x/lint/golint
	$(go_install_dev) github.com/alexkohler/nakedret
	$(go_install_dev) github.com/alexkohler/unimport
	$(go_install_dev) 4d63.com/gochecknoinits
	$(go_install_dev) github.com/kisielk/errcheck
	$(go_install_dev) github.com/gordonklaus/ineffassign
	$(go_install_dev) github.com/remyoudompheng/go-misc/deadcode
	$(go_install_dev) github.com/mdempsky/unconvert
	$(go_install_dev) github.com/mdempsky/maligned
	$(go_install_dev) github.com/jgautheron/goconst/cmd/goconst
	$(go_install_dev) mvdan.cc/unparam
	$(go_install_dev) github.com/tommy-muehle/go-mnd/cmd/mnd
	$(go_install_dev) github.com/alexkohler/prealloc
	$(go_install_dev) github.com/alexkohler/dogsled/cmd/dogsled
	$(go_install_dev) github.com/nishanths/exhaustive/...
	$(go_install_dev) github.com/kyoh86/looppointer/cmd/looppointer
	$(go_install_dev) github.com/go-critic/go-critic/cmd/gocritic
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b ${GOPATH}/bin v2.3.0
	@echo "INSTALLLING MANUAL ANALYSIS TOOLS"
	$(go_install_dev) github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

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
	@echo "VETTING..."
	@go vet ./...
	@unconvert -v ./...
	@maligned ./...
	@goconst -ignore-tests ./...
	@unparam -exported -tests ./...
	@mnd ./...
	@prealloc -set_exit_status ./...
	@nakedret -l 0 ./...
	@unimport ./...
	@gochecknoinits ./...
	@dogsled -n 1 -set_exit_status ./...
	@exhaustive -maps ./...
	@looppointer ./...
	@gocritic check ./...
	@echo "SECURITY SCAN..."
	@gosec -conf .gosecrc.json -quiet ./...
	@echo "VERIFYING ERRORS ARE CHECKED..."
	@errcheck -asserts -blank -ignoretests -exclude errcheck_excludes.txt ./...
	@echo "CHECKING FOR DEAD CODE..."
	@ineffassign ./*
	@deadcode ./internal/*
	@deadcode ./cmd/*

format:
	go fmt ./...
	goimports -w .

lint: lint-go lint-md

lint-go:
	golint -set_exit_status ./...

lint-md:
	npx markdownlint-cli -c .markdownlintrc.yml $(markdown_files)

clean:
	rm -rf $(executable_file)*
	rm -rf $(coverage_file)
	rm `find ./ -name '_fuzz'` -rf
	rm `find ./ -name '*-fuzz.zip'` -rf

.PHONY: default install build clean format lint analysis test fuzz
