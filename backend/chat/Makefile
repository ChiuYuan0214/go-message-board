GOTOOLCHAIN ?= go1.26.1
GOLANGCI_LINT ?= go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8
STATICCHECK ?= go run honnef.co/go/tools/cmd/staticcheck@2025.1.1

.PHONY: build test lint lint-style

build:
	go build ./...

test:
	go test ./...

lint:
	GOTOOLCHAIN=$(GOTOOLCHAIN) $(GOLANGCI_LINT) run ./...
	GOTOOLCHAIN=$(GOTOOLCHAIN) $(STATICCHECK) ./...
	./scripts/lint_style.sh

lint-style:
	./scripts/lint_style.sh
