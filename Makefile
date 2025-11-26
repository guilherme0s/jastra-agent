GO ?= go

GOLANGCI_BIN := $(shell command -v golangci-lint 2>/dev/null || echo $(GOBIN)/golangci-lint)
GOLANGCI_ARGS ?= --timeout=5m

GOLANGCI_LINT_MODULE ?= github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: lint
lint:
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "golangci-lint not found. Installing $(GOLANGCI_LINT_MODULE)..."; \
		$(GO) install $(GOLANGCI_LINT_MODULE); \
	fi
	@$(GOLANGCI_BIN) run $(GOLANGCI_ARGS) ./...

.PHONY: run
run:
	@$(GO) run ./cmd/server/main.go server

.PHONY: migrate
migrate:
	@$(GO) run ./cmd/server/main.go migrate
