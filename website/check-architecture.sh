#!/bin/bash

# Architecture Compatibility Check Script
# This script helps diagnose Docker architecture compatibility issues

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

echo "=== Docker Architecture Compatibility Check ==="
echo ""

# Check system architecture
print_info "Checking system architecture..."
SYSTEM_ARCH=$(uname -m)
echo "System architecture: $SYSTEM_ARCH"

case $SYSTEM_ARCH in
    "x86_64")
        print_success "System is x86_64 (compatible with amd64 Docker images)"
        EXPECTED_DOCKER_ARCH="amd64"
        ;;
    "aarch64"|"arm64")
        print_warning "System is ARM64. Docker images need to be built for arm64 platform"
        EXPECTED_DOCKER_ARCH="arm64"
        ;;
    *)
        print_warning "Unknown system architecture: $SYSTEM_ARCH"
        EXPECTED_DOCKER_ARCH="unknown"
        ;;
esac
echo ""

# Check if Docker is available
if command -v docker &> /dev/null; then
    print_success "Docker is installed"

    # Check Docker daemon
    if docker info &> /dev/null; then
        print_success "Docker daemon is running"
    else
        print_error "Docker daemon is not running"
        exit 1
    fi
else
    print_error "Docker is not installed"
    exit 1
fi
echo ""

# Check if float-island-website image exists
if docker images | grep -q "float-island-website"; then
    print_info "Found float-island-website images:"
    docker images float-island-website --format "table {{.Repository}}\t{{.Tag}}\t{{.ID}}\t{{.Size}}\t{{.CreatedAt}}"

    echo ""
    print_info "Checking image architectures..."

    # Get all tags of the image
    docker images float-island-website --format "{{.Tag}}" | while read -r tag; do
        if [ "$tag" != "TAG" ]; then  # Skip header
            IMAGE_NAME="float-island-website:$tag"
            IMAGE_ARCH=$(docker inspect "$IMAGE_NAME" --format='{{.Architecture}}' 2>/dev/null || echo "unknown")

            echo -n "Image $IMAGE_NAME: "
            if [ "$IMAGE_ARCH" = "$EXPECTED_DOCKER_ARCH" ]; then
                print_success "architecture $IMAGE_ARCH (matches system)"
            elif [ "$IMAGE_ARCH" = "unknown" ]; then
                print_warning "architecture unknown (unable to determine)"
            else
                print_error "architecture $IMAGE_ARCH (does not match system $EXPECTED_DOCKER_ARCH)"
                echo "  This may cause 'exec format error' when running containers"
            fi
        fi
    done
else
    print_warning "No float-island-website images found"
fi
echo ""

# Check running containers
if docker ps -a | grep -q "float-island-website"; then
    print_info "Found float-island-website containers:"
    docker ps -a --filter "name=float-island-website" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

    echo ""
    print_info "Recent container logs (last 10 lines):"
    CONTAINER_ID=$(docker ps -a --filter "name=float-island-website" --format "{{.ID}}" | head -1)
    if [ -n "$CONTAINER_ID" ]; then
        docker logs --tail 10 "$CONTAINER_ID" 2>&1 | while read -r line; do
            echo "  $line"
        done
    fi
else
    print_info "No float-island-website containers found"
fi
echo ""

# Recommendations
print_info "Recommendations:"
if [ "$SYSTEM_ARCH" = "x86_64" ] && [ "$EXPECTED_DOCKER_ARCH" = "amd64" ]; then
    echo "  ✓ Your system is x86_64 - use images built for linux/amd64 platform"
    echo "  ✓ Run: ./package-docker-image.sh to build compatible image"
elif [ "$SYSTEM_ARCH" = "aarch64" ] || [ "$SYSTEM_ARCH" = "arm64" ]; then
    echo "  ! Your system is ARM64 - build images with --platform linux/arm64"
    echo "  ! Modify Dockerfile and build scripts for ARM64 compatibility"
else
    echo "  ? Unknown architecture - verify Docker platform compatibility"
fi

echo ""
echo "For more help, see: https://docs.docker.com/desktop/multi-arch/"