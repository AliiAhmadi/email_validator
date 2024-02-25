.DEFAUL_GOAL := run

run:
	@go run .

tidy:
	@go mod tidy

build:
	@go build -o emailValidator

.PHONY: run