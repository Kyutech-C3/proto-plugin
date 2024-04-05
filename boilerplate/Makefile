# Variables
PROTO_DIR := ./sample
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)
PROTO_OUT_DIR := .
PROTO_PLUGIN_NAME := boilerplate
PROTO_PLUGIN_FILE := protoc-gen-$(PROTO_PLUGIN_NAME)

# Targets
.PHONY: all clean

all: build gen

gen:
	@mkdir -p $(PROTO_OUT_DIR)
	@protoc -I. --plugin=./$(PROTO_PLUGIN_FILE) --$(PROTO_PLUGIN_NAME)_out=. --go_out=$(PROTO_OUT_DIR) $(PROTO_FILES)

build:
	@go build -o $(PROTO_PLUGIN_FILE) .

clean:
	@rm -rf $(PROTO_OUT_DIR)
	@rm -f $(PROTO_PLUGIN_FILE)