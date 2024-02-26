.DEFAUL_GOAL := run

run:
	@go run .

tidy:
	@go mod tidy

build:
	@go build -o emailValidator

clean:
	@rm -rf emailValidator

compile_test:
	@go build -o temp && rm -rf temp
	@echo "compile status: OK"

vendor:
	@go mod vendor

.PHONY: run tidy build clean compile_test vendor