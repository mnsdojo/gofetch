#!/bin/bash

# Set the program name and install directory
PROGRAM_NAME="gofetch"
INSTALL_DIR="/usr/local/bin"
REPO_URL="https://github.com/mnsdojo/gofetch.git"
TEMP_DIR="/tmp/gofetch"

# Create a temporary directory
mkdir -p "$TEMP_DIR"
cd "$TEMP_DIR" || exit

# Clone the repository
git clone "$REPO_URL" .
if [ $? -ne 0 ]; then
    echo "Error: Failed to clone repository."
    exit 1
fi

# Build the Go program
go build -o "$PROGRAM_NAME"
if [ $? -ne 0 ]; then
    echo "Error: Build failed. Please check your Go code and try again."
    exit 1
fi

# Install the program
sudo mv "$PROGRAM_NAME" "$INSTALL_DIR/"
if [ $? -ne 0 ]; then
    echo "Error: Installation failed. You may need superuser privileges."
    exit 1
fi

# Clean up
rm -rf "$TEMP_DIR"

echo "Installation successful. You can now use '$PROGRAM_NAME' from the terminal."
