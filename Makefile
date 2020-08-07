GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

ASSETS_DATA_FILES := $(shell find assets/static | sed 's/  /\\ /g' | xargs)
TEMPLATES_DATA_FILES := $(shell find templates | sed 's/  /\\ /g' | xargs)

.PHONY: default
default: run

.PHONY: build
build: fmt
	go build -o mir-examples main.go

.PHONY: build
run: fmt
	go run main.go

.PHONY: gen-assets
gen-assets: $(ASSETS_DATA_FILES)
	-rm -f internal/assets/assets_gen.go
	go generate internal/assets/assets.go
	@$(GOFMT) internal/assets

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
