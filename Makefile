.PHONY: all clean deps fmt vet test docker

EXECUTABLE ?= drone-gke
IMAGE ?= yurifl/$(EXECUTABLE)
COMMIT ?= $(shell git rev-parse --short HEAD)

LDFLAGS = -X "main.buildCommit=$(COMMIT)"
PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: deps build

clean:
	go clean -i ./...

deps:
	go get -t ./...

build: $(wildcard *.go)
	go build -ldflags '-s -w $(LDFLAGS)'

publish:
	docker build --rm -t yurifl/drone-gke .
	docker push yurifl/drone-gke
