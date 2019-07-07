
.PHONY: build
build: build/snowflake

GO_SOURCES=$(shell find . -type f -name "*.go") go.mod go.sum

build/snowflake: $(GO_SOURCES)
	$(call GO_BUILD,.,$@)

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint: \
	golangci-lint

GOLANGCI_LINT_VERSION=v1.17.1
GOLANGCI_LINT_DIR=$(shell go env GOPATH)/pkg/golangci-lint/$(GOLANGCI_LINT_VERSION)
$(GOLANGCI_LINT_DIR):
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOLANGCI_LINT_DIR) $(GOLANGCI_LINT_VERSION)

.PHONY: install-golangci-lint
install-golangci-lint: $(GOLANGCI_LINT_DIR)

.PHONY: golangci-lint
golangci-lint: install-golangci-lint
	$(GOLANGCI_LINT_DIR)/golangci-lint run
