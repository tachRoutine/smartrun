#!/bin/bash

# SmartRun Build Script

set -e

echo "Building SmartRun..."

# Clean previous builds
rm -rf dist/

# Create dist directory
mkdir -p dist/

# Build for current platform
go build -o dist/smartrun ./cmd/smartrun

echo "Build completed successfully!"
echo "Binary available at: dist/smartrun"