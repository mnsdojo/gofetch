#!/bin/bash

# Define variables
PROGRAM_NAME="gofetch"
INSTALL_DIR="/usr/local/bin"
SOURCE_DIR="$(dirname "$0")/.."  # Directory containing main.go (parent of scripts directory)

# Navigate to the root directory where main.go is located
cd "$SOURCE_DIR" || { echo "Error: Unable to change directory to $SOURCE_DIR"; exit 1; }

# Build the Go program
go build -o "$PROGRAM_NAME"
if [ $? -ne 0 ]; then
    echo "Error: Build failed. Please check your Go code and try again."
    exit 1
fi

# Move the binary to the installation directory
sudo mv "$PROGRAM_NAME" "$INSTALL_DIR/"
if [ $? -ne 0 ]; then
    echo "Error: Installation failed. You may need superuser privileges."
    exit 1
fi

echo "Installation successful. You can now use '$PROGRAM_NAME' from the terminal."
