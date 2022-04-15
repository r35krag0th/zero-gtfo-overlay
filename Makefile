all: build

build: generate test
	@echo ""
	@echo "\033[32m==>\033[0m Building (overlay)"
	go build -o zero-gtfo-overlay ./overlay/overlay.go
	@echo ""
	@echo "\033[32m==>\033[0m Building (rundown)"
	go build -o zero-gtfo-rundown ./rundown/rundown.go

test:
	@echo ""
	@echo "\033[33m==>\033[0m Running Tests"
	go test -v ./overlay/...
	go test -v ./rundown/...

generate:
	@echo ""
	@echo "\033[35m==>\033[0m Generating"
	go generate ./...

run-overlay: build
	@echo ""
	@echo "\033[36m==>\033[0m Running (overlay)"
	env APP_ENV=dev $(PWD)/zero-gtfo-overlay -f overlay/etc/overlay-api.yaml

run-rundown: build
	@echo ""
	@echo "\033[36m==>\033[0m Running (rundown)"
	env APP_ENV=dev $(PWD)/zero-gtfo-rundown -f rundown/etc/rundown-api.yaml
