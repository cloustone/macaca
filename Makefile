IMPORTPATH = github.com/cloustone/macaca
#V := 1 # When V is set, print cmd and build progress.
Q := $(if $V,,@)

VERSION          := $(shell git describe --tags --always --dirty="-dev")
DATE             := $(shell date -u '+%Y-%m-%d-%H%M UTC')
VERSION_FLAGS    := -ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"'
DEBUG_FLAGS 	 := -gcflags "-N -l"

.PHONY: all
all: build

.PHONY: build
build: macaca benchmark macc

.PHONY: macaca 
macaca: 
	@echo "building macaca..."
	$Q CGO_ENABLED=0 go build  $(DEBUG_FLAGS) -v

.PHONY: benchmark 
benchmark: 
	@echo "building benchmark..."
	$Q CGO_ENABLED=0 go build -v $(DEBUG_FLAGS) -o bin/macaca-benchmark $(IMPORTPATH)/tools/benchmark

.PHONY: macc 
macc: 
	@echo "building macc..."
	$Q CGO_ENABLED=0 go build -v $(DEBUG_FLAGS) -o bin/macc $(IMPORTPATH)/tools/macc



