IMPORTPATH = github.com/cloustone/macaca
# V := 1 # When V is set, print cmd and build progress.
Q := $(if $V,,@)

VERSION          := $(shell git describe --tags --always --dirty="-dev")
DATE             := $(shell date -u '+%Y-%m-%d-%H%M UTC')
VERSION_FLAGS    := -ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"'

.PHONY: all
all: build

.PHONY: build
build: macaca benchmark

.PHONY: macaca 
macaca: 
	@echo "building macaca..."
	$Q CGO_ENABLED=0 go build -v -o bin/macaca $(IMPORTPATH)/

.PHONY: benchmark 
benchmark: 
	@echo "building benchmark..."
	$Q CGO_ENABLED=0 go build -v -o bin/macaca-benchmark $(IMPORTPATH)/benchmark


