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

.PHONY: run