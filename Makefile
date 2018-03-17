.PHONY: all check-path build clean install uninstall fmt simplify run
.DEFAULT_GOAL: $(BIN_FILE)

SHELL := /bin/bash

export GOPATH := $(shell pwd)

PROJECT_NAME = microservice-email

BUILD_DIR = build
BIN_DIR = bin
BIN_FILE = $(PROJECT_NAME)

SRC_DIR = ./src/$(PROJECT_NAME)
SRC_PKGS = $$(go list $(SRC_DIR)/...)
SRC_FILES = $(shell find . -type f -name '*.go' -path "$(SRC_DIR)/*")

VERSION := 1.0.2
BUILD := `git rev-parse HEAD`

# Use linker flags to provide version/build settings to the binary
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

DOCKER_COMPOSE_CMD = docker-compose -p $(PROJECT_NAME) -f environment/docker-compose.yml


all: dep build

check-path:
ifndef GOPATH
	@echo "FATAL: you must declare GOPATH environment variable, for more"
	@echo "       details, please check"
	@echo "       http://golang.org/doc/code.html#GOPATH"
	@exit 1
endif

dep: check-path
	@echo "Downloading dependencies..."
	@cd $(SRC_DIR)/ && dep ensure
	@echo "Finish..."

build: check-path
	@echo "Compiling $(BUILD_DIR)/$(BIN_FILE)..."
	@go build $(LDFLAGS) -i -o $(BUILD_DIR)/$(BIN_FILE) $(SRC_DIR)/main.go
	@echo "Finish..."

clean:
	@rm -rf $(BIN_DIR)/$(BIN_FILE)
	@rm -rf $(BUILD_DIR)
	@rm -rf pkg/
	@rm -rf bin/

install: check-path
	@go install $(LDFLAGS) $(SRC_PKGS)
	@cp $(BIN_DIR)/$(BIN_FILE) /usr/local/bin/
	@cp etc/config.yml /etc/microservice-email.yml
	@echo "Instalation complete..."

uninstall:
	@rm -f $$(which $(BIN_FILE))

fmt: check-path
	@gofmt -l -w $(SRC_FILES)

simplify: check-path
	@gofmt -s -l -w $(SRC_FILES)

run: build
	$(BUILD_DIR)/$(BIN_FILE) -config-file=etc/config-dev.yml -log-level=debug

# -------------------------------------------------------------------
# -								Docker								-
# -------------------------------------------------------------------

docker_build:
	@$(DOCKER_COMPOSE_CMD) build $(PROJECT_NAME)

docker_shell:
	@$(DOCKER_COMPOSE_CMD) run --rm $(PROJECT_NAME) /bin/bash

docker_run:
	@$(DOCKER_COMPOSE_CMD) run --rm --service-ports --name $(PROJECT_NAME) $(PROJECT_NAME) \
		/bin/bash -ci "make dep && make run"
