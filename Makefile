OUT_DIR := ./build
OUT_BIN_NAME := main

.PHONY: build-server clean

build-server:
	@go build -o $(OUT_DIR)/server ./cmd/server

clean:
	@rm -r $(OUT_DIR)
