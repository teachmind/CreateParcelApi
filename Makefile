SHELL=/bin/bash
APP=CreateParcelApi
APP_EXECUTABLE="./build/$(APP)"
APP_COMMIT=$(shell git rev-parse HEAD)
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
SOURCE_DIRS=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
COVERAGE_MIN=95

.PHONY: build

all: clean test

clean:
	@echo "> cleaning up the mess"
	@rm -rf build && mkdir -p build

lint:
	@echo "> running linter $(SOURCE_DIRS)/..."
	@golangci-lint run -v --timeout 5m $(SOURCE_DIRS)/...

server: build
	@echo "> running server command"
	@${APP_EXECUTABLE} server