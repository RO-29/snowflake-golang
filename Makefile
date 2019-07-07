
.PHONY: build
build: build/snowflake

GO_SOURCES=$(shell find . -type f -name "*.go") go.mod go.sum

build/snowflake: $(GO_SOURCES)
	$(call GO_BUILD,.,$@)
