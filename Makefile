program_main:=./cmd/wordrow
executable_file:=wordrow

unit_test_root:=./internal/...
integration_test_root:=./cmd/wordrow/...
coverage_file:=coverage.out
fuzz_dir:=./_fuzz

markdown_files:=./*.md ./docs/*.md ./.github/**/*.md

go_nomod:=GO111MODULE=off


default: build

install: install-deps install-dev-deps

install-deps:
	@echo "INSTALLLING DEPENDENCIES"
	go get -u github.com/yargevad/filepathx

install-dev-deps:
	@echo "INSTALLLING DEVELOPMENT TOOLS"
	$(go_nomod) go get golang.org/x/tools/cmd/goimports
	@echo "INSTALLLING STATIC ANALYSIS TOOLS"
	$(go_nomod) go get -u golang.org/x/lint/golint
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b ${GOPATH}/bin v2.3.0
	@echo "INSTALLLING MANUAL ANALYSIS TOOLS"
	$(go_nomod) go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

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
	@echo "VETTING"
	go vet ./...
	@echo "SECURITY SCAN"
	gosec -quiet ./...

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
