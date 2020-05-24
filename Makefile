executable_file=wordrow.o
coverage_file=coverage.out
test_root=./internal/...
fuzz_dir="./_fuzz"

build:
	go build -o $(executable_file) ./cmd/wordrow/

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
	rm -rf $(coverage_file)
	rm -rf $(executable_file)
	rm -rf **/*/*-fuzz.zip
	rm -rf **/*/_fuzz/
