program_main=./cmd/wordrow
executable_file=wordrow

test_root=./internal/...
coverage_file=coverage.out

fuzz_dir="./_fuzz"


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
	cd ${PKG}; go-fuzz-build; go-fuzz -func Fuzz${FUNC} -workdir ${fuzz_dir}

benchmark:
	go test $(test_root) -bench=. -run=XXX

analyze:
	go vet ./...

format:
	go fmt ./...

lint: lint-go lint-md

lint-go:
	golint -set_exit_status ./...

lint-md:
	npx markdownlint-cli -c .markdownlintrc.yml ./*.md ./**/*.md

clean:
	rm -rf $(executable_file)*
	rm -rf $(coverage_file)
	rm -rf **/*/*-fuzz.zip
	rm -rf **/*/_fuzz/
