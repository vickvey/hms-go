# Define variables
BUILD_DIR := build
EXEC_NAME := hms-go

# Define targets
.PHONY: build_windows build_macos build_linux clean

# Default target
build: build_windows build_macos build_linux

# Build for Windows
build_windows:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(EXEC_NAME)-windows-amd64.exe

# Build for macOS
build_macos:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(EXEC_NAME)-macos-amd64

# Build for Linux
build_linux:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(EXEC_NAME)-linux-amd64

# Clean build directory
clean:
	rm -rf $(BUILD_DIR)
