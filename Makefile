GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

LDFLAGS += -X "github.com/alimy/kr/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "github.com/alimy/kr/version.GitHash=$(shell git rev-parse HEAD)"

.PHONY: build
build: fmt
	go build  -ldflags '$(LDFLAGS)' -o kr main.go

.PHONY: generate
generate:
	-rm -f internal/generator/templates_gen.go
	go generate internal/generator/templates.go
	$(GOFMT) -w internal/generator/templates_gen.go

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)
