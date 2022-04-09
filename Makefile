all: build test run

build:
	@echo ""
	@echo "\033[32m==>\033[0m Building"
	go build -o zero-gtfo-overlay

test:
	@echo ""
	@echo "\033[33m==>\033[0m Running Tests"
	go test -v ./...

run:
	@echo ""
	@echo "\033[36m==>\033[0m Running"
	$(PWD)/zero-gtfo-overlay
