#!/bin/bash

# Configuration
IMAGE_NAME="float-manager"

# 1. Try to get version from command line argument
VERSION=$1

# 2. If no argument, try to get from package.json
if [ -z "$VERSION" ]; then
    # Use grep to find the version line, expected format: "version": "0.0.0",
    # We restrict to lines starting with spaces/tabs to avoid finding nested dependency versions if possible,
    # though usually the main version is at the top.
    VERSION=$(grep -m1 '"version":' package.json | cut -d '"' -f 4)
fi

# 3. Fallback to latest if still empty
if [ -z "$VERSION" ]; then
    echo "⚠️  Could not detect version. Using 'latest'."
    VERSION="latest"
fi

OUTPUT_FILE="${IMAGE_NAME}_${VERSION}.tar.gz"

echo "========================================"
echo "Start building Docker image: $IMAGE_NAME:$VERSION"
echo "========================================"

# Add --platform linux/amd64 for cross-platform compatibility (e.g., M1 Mac -> Linux Server)
# We tag both :latest and :$VERSION
if docker build --platform linux/amd64 -t $IMAGE_NAME:latest -t $IMAGE_NAME:$VERSION .; then
    echo "✅ Build successful."
else
    echo "❌ Build failed."
    exit 1
fi

echo "========================================"
echo "Exporting and compressing image..."
echo "========================================"

# Save the specific versioned image
if docker save $IMAGE_NAME:$VERSION | gzip > $OUTPUT_FILE; then
    echo "✅ Image saved to $OUTPUT_FILE"
    ls -lh $OUTPUT_FILE
else
    echo "❌ Failed to save image."
    exit 1
fi

echo "Done."
