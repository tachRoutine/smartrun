#!/bin/bash

# SmartRun Installation Script

set -e

INSTALL_DIR="/usr/local/bin"
BINARY_NAME="smartrun"

echo "Installing SmartRun..."

# Check if binary exists
if [ ! -f "dist/$BINARY_NAME" ]; then
    echo "Binary not found. Please run build.sh first."
    exit 1
fi

# Copy binary to install directory
sudo cp "dist/$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"

# Make it executable
sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"

echo "SmartRun installed successfully!"
echo "You can now run: smartrun"