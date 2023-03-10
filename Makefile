BINARY=userstore

TARGET_DIR=target
BRANCH ?= $(shell git branch | grep "^\*" | sed 's/^..//')
COMMIT ?= $(shell git rev-parse --short HEAD)
VERSION=$(BRANCH)-$(COMMIT)
BUILDTIME ?= $(shell date -u +%FT%T)

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST ?=$(GOCMD) test
GORUN=$(GOCMD) run
GOLANG_LINTER ?= golangci-lint

.PHONY: lint
lint:
	$(GOLANG_LINTER) run ./...

.PHONY: lint-fix
lint-fix:
	$(GOLANG_LINTER) run --fix ./...

.PHONY: test-unit
test-unit:
	$(GOTEST) ./... -count=1 -race -cover
