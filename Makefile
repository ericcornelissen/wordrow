executable_file=wordrow.o
coverage_file=coverage.out
test_root=./internal/...


build:
	go build -o $(executable_file) cmd/wordrow/main.go

test:
	go test $(test_root)

coverage:
	 go test $(test_root) -coverprofile $(coverage_file)
	 go tool cover -html=$(coverage_file)

benchmark:
	go test $(test_root) -bench=. -run=XXX

lint: lint-go lint-md

lint-go:
	golint -set_exit_status ./...

lint-md:
	npx markdownlint-cli -c .markdownlintrc.yml ./*.md ./**/*.md

clean:
	rm -rf $(coverage_file)
	rm -rf $(executable_file)
