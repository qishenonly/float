#!/bin/bash

# Test script for Docker image packaging
# This script tests the package-docker-image.sh functionality

set -e

echo "Testing Docker image packaging script..."

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

# Check if script exists and is executable
if [ ! -f "package-docker-image.sh" ]; then
    print_error "package-docker-image.sh not found"
    exit 1
fi

if [ ! -x "package-docker-image.sh" ]; then
    print_error "package-docker-image.sh is not executable"
    exit 1
fi

print_success "Script exists and is executable"

# Check script syntax
if bash -n package-docker-image.sh; then
    print_success "Script syntax is valid"
else
    print_error "Script has syntax errors"
    exit 1
fi

# Check if help text is shown (by checking script comments)
if grep -q "Usage:" package-docker-image.sh; then
    print_success "Usage documentation found in script"
else
    print_error "Usage documentation missing"
    exit 1
fi

# Check if platform specification exists
if grep -q "linux/amd64" package-docker-image.sh; then
    print_success "Linux x86 platform specification found"
else
    print_error "Platform specification missing"
    exit 1
fi

# Check if docker save command exists in script
if grep -q "docker save" package-docker-image.sh; then
    print_success "Docker save command found in script"
else
    print_error "Docker save command missing"
    exit 1
fi

echo ""
print_success "All tests passed! The packaging script is ready to use."
echo ""
echo "To package the Docker image, run:"
echo "  ./package-docker-image.sh"
echo ""
echo "To package with a specific version:"
echo "  ./package-docker-image.sh v1.0.0"