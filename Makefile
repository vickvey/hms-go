# Define variables
BUILD_DIR := build
EXEC_NAME := hms-go
SRC_DIR := cmd
SRCS := $(wildcard $(SRC_DIR)/*.go)

# Determine the operating system
ifeq ($(OS),Windows_NT)
    OS_FLAG := windows
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        OS_FLAG := linux
    endif
    ifeq ($(UNAME_S),Darwin)
        OS_FLAG := darwin
    endif
endif

# Define targets
.PHONY: build clean run

# Default target
build: build_$(OS_FLAG)

# Build target for Windows
build_windows:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(EXEC_NAME)-windows-amd64.exe $(SRCS)

# Build target for macOS
build_darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(EXEC_NAME)-macos-amd64 $(SRCS)

# Build target for Linux
build_linux:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(EXEC_NAME)-linux-amd64 $(SRCS)

# Clean target
clean:
	rm -rf $(BUILD_DIR)

# Run target
run:
	@mkdir -p $(BUILD_DIR)
	@[ -f "$(BUILD_DIR)/$(EXEC_NAME)-$(OS_FLAG)-amd64$(if $(filter windows,$(OS_FLAG)),.exe,)" ] || make build_$(OS_FLAG)
	@./$(BUILD_DIR)/$(EXEC_NAME)-$(OS_FLAG)-amd64$(if $(filter windows,$(OS_FLAG)),.exe,)
