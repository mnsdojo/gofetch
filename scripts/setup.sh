#!/bin/bash

# Step 1: Build the project
echo "Building the project..."
go build -o bin/main cmd/cli/main.go

# Step 2: Move the binary to a directory in PATH
echo "Moving the binary to /usr/local/bin..."
sudo mv bin/main /usr/local/bin/gofetch

# Step 3: Define the script file to execute
SCRIPT_FOLDER="scripts"
SCRIPT_FILE="your_script.sh"  # Update this with your actual script file name

# Step 4: Move the script file to /usr/local/bin
echo "Moving the script file to /usr/local/bin..."
sudo mv "$SCRIPT_FOLDER/$SCRIPT_FILE" /usr/local/bin

echo "Installation completed."
