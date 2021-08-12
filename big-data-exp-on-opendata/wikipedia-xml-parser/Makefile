BINARY_NAME=wikipedia-xml-parser

.PHONY: help build test cover clean install
.DEFAULT_GOAL := all

all: test build ## test and build

build: ## create binary
	go build -ldflags="-s -w" -o ${BINARY_NAME} ./...

test: ## run the tests
	go test ./... -v

cover: ## check tests coverage
	go test ./... -cover

clean: ## remove unnecessary items
	go clean
	rm ${BINARY_NAME}

install: ## install binary
	make build
	mkdir -p ~/.local/bin
	sudo cp ${BINARY_NAME} /usr/bin/

help: ## display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
