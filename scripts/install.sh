#!/bin/bash

# Define the URL of the installation script and the temporary file location
SCRIPT_URL="https://raw.githubusercontent.com/mnsdojo/gofetch/main/scripts/install.sh"
TEMP_SCRIPT="/tmp/install.sh"

# Download the script
echo "Downloading installation script..."
curl -fsSL "$SCRIPT_URL" -o "$TEMP_SCRIPT"

# Check if download was successful
if [ $? -ne 0 ]; then
    echo "Error: Failed to download the script."
    exit 1
fi

# Make the script executable
echo "Making the script executable..."
chmod +x "$TEMP_SCRIPT"

# Run the script
echo "Running the installation script..."
"$TEMP_SCRIPT"

# Check if script executed successfully
if [ $? -ne 0 ]; then
    echo "Error: Installation failed."
    exit 1
fi

# Optionally, clean up
echo "Cleaning up..."
rm "$TEMP_SCRIPT"

echo "Installation completed successfully."
