GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

ASSETS_DATA_FILES := $(shell find assets/static | sed 's/  /\\ /g' | xargs)
CONFIG_DATA_FILES := $(shell find assets/config | sed 's/  /\\ /g' | xargs)
TEMPLATES_DATA_FILES := $(shell find templates | sed 's/  /\\ /g' | xargs)

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

.PHONY: gen-assets
gen-assets: $(ASSETS_DATA_FILES)
	-rm -f internal/assets/assets_gen.go
	go generate internal/assets/assets.go
	@$(GOFMT) internal/assets

.PHONY: gen-config
gen-config: $(CONFIG_DATA_FILES)
	-rm -f internal/config/config_gen.go
	go generate internal/config/config.go
	@$(GOFMT) internal/config

.PHONY: gen-templates
gen-templates: $(TEMPLATES_DATA_FILES)
	-rm -f internal/templates/templates_gen.go
	go generate internal/templates/templates.go
	@$(GOFMT) internal/templatess

.PHONY: generate
generate:
	@go generate mirc/main.go
	@$(GOFMT) ./

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)
