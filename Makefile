OUT_DIR := ./build
OUT_BIN_NAME := main

.PHONY: test build-server clean

test:
	@go test -v ./...

build-server:
	@go build -o $(OUT_DIR)/server ./cmd/server

clean:
	@rm -r $(OUT_DIR)
