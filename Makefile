GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

TARGET = alexandrite

TAGS = release
LDFLAGS += -X "github.com/alimy/alexandrite/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "github.com/alimy/alexandrite/version.GitHash=$(shell git rev-parse HEAD)"

.PHONY: default
default: run

.PHONY: build
build: fmt
	go build -ldflags '$(LDFLAGS)' -tags '$(TAGS)' -o $(TARGET) main.go

.PHONY: build
run: fmt
	go run -ldflags '$(LDFLAGS)' -tags '$(TAGS)' main.go serve

.PHONY: generate
generate:
	@go generate mirc/main.go
	@$(GOFMT) ./

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)
