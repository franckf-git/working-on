BINARY_NAME=lite-api-crud
.DEFAULT_GOAL := all

all: install test build

.PHONY: build test cover run clean install help

build:
	go build -ldflags="-s -w" -o ${BINARY_NAME} main.go

test:
	go test -v ./...

cover:
	go test -cover ./...

run:
	make build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

install: ## Install dependencies
	go install

help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'