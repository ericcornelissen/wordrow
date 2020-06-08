program_main:=./cmd/wordrow
executable_file:=wordrow

test_root:=./internal/...
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

test:
	go test $(test_root)

coverage:
	go test $(test_root) -coverprofile $(coverage_file)
	go tool cover -html=$(coverage_file)

fuzz:
	cd ${PKG}; \
	go-fuzz-build; \
	go-fuzz -func Fuzz${FUNC} -workdir ${fuzz_dir}

benchmark:
	go test $(test_root) -bench=. -run=XXX

analysis:
	@echo "VETTING"
	go vet ./...
	@echo "SECURITY SCAN"
	gosec -quiet ./...

format:
	go fmt ./...

lint: lint-go lint-md

lint-go:
	golint -set_exit_status ./...

lint-md:
	npx markdownlint-cli -c .markdownlintrc.yml $(markdown_files)

clean:
	rm -rf $(executable_file)*
	rm -rf $(coverage_file)
	rm -rf **/*/*-fuzz.zip
	rm -rf **/*/_fuzz/

.PHONY: default install build clean format lint analysis test fuzz
