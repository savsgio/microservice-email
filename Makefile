.PHONY: all get build clean run
.DEFAULT_GOAL: $(BIN_FILE)

PROJECT_NAME = microservice-email

BIN_DIR = ./bin
BIN_FILE = $(PROJECT_NAME)

CMD_DIR = ./cmd
CONFIG_DIR = ./config/

# Get version constant
VERSION := 1.4.0
BUILD := $(shell git rev-parse HEAD)

# Use linker flags to provide version/build settings to the binary
LDFLAGS=-ldflags "-s -w -X=main.version=$(VERSION) -X=main.build=$(BUILD)"


default: get build

get:
	@echo "[*] Downloading dependencies..."
	cd $(CMD_DIR)/$(PROJECT_NAME) && go get
	@echo "[*] Finish..."

vendor:
	@go mod vendor

build:
	@echo "[*] Building $(PROJECT_NAME)..."
	go build $(LDFLAGS) -o $(BIN_DIR)/$(BIN_FILE) $(CMD_DIR)/...
	@echo "[*] Finish..."

test:
	go test -v -race -cover ./...

run: build
	$(BIN_DIR)/$(BIN_FILE) -config-file $(CONFIG_DIR)/$(PROJECT_NAME).conf.yml

install:
	mkdir -p /etc/$(PROJECT_NAME)/
	cp $(BIN_DIR)/$(BIN_FILE) /usr/local/bin/
	cp $(CONFIG_DIR)/$(PROJECT_NAME).conf.yml /etc/

uninstall:
	rm -rf /usr/local/bin/$(BIN_FILE)
	rm -rf /etc/$(PROJECT_NAME)/

clean:
	rm -rf bin/
	rm -rf vendor/

docker_build:
	cd docker && docker-compose build

docker_run:
	cd docker && \
		docker-compose run $(PROJECT_NAME) -config-file /code/$(CONFIG_DIR)/$(PROJECT_NAME).dev.conf.yml
