# Float Island Website

A modern, beautiful website for Float Island - a personal finance management app with glass morphism design.

## Features

- 🎨 Glass morphism UI design
- 📱 Responsive design
- 🔒 Privacy policy and terms of use in modals
- ⚡ Fast loading with Vite
- 🐳 Docker containerization
- 🚀 Easy deployment with docker-compose

## Development

### Prerequisites

- Node.js 18+
- npm or yarn
- Docker (for containerization)

### Installation

```bash
npm install
```

### Development Server

```bash
npm run dev
```

### Build for Production

```bash
npm run build
npm run preview
```

## Docker Deployment

### Quick Build and Run

```bash
# Build Docker image
./build.sh

# Run with docker-compose
docker-compose up -d
```

The website will be available at `http://localhost:8080`

### Package Docker Image for Deployment

For deploying to systems without Docker registry access, you can package the Docker image into a tar file:

```bash
# Package Docker image for Linux x86
./package-docker-image.sh

# Package with specific tag
./package-docker-image.sh v1.0.0

# Package with custom output filename
./package-docker-image.sh latest my-deployment.tar
```

**Important:** The packaged image is built specifically for Linux x86 (amd64) architecture. Make sure your target deployment server is running on x86 architecture. If you encounter "exec format error", it usually means there's an architecture mismatch.

### Advanced Build Options

```bash
# Build with specific tag
./build.sh v1.0.0

# Build and push to registry
./build.sh latest username/float-island
```

### Production Deployment

For production deployment with nginx reverse proxy:

```bash
# Start with nginx
docker-compose --profile production up -d

# Or build and deploy
./build.sh
docker-compose --profile production up -d
```

### Docker Commands

```bash
# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Rebuild and restart
docker-compose up -d --build

# Clean up
docker-compose down -v --rmi all
```

## Project Structure

```
website/
├── src/
│   ├── components/        # HTML component files
│   │   ├── header.html
│   │   ├── hero.html
│   │   ├── features.html
│   │   ├── showcase.html
│   │   ├── reviews.html
│   │   ├── faq.html
│   │   ├── footer.html
│   │   ├── mesh-bg.html
│   │   ├── privacy-modal.html
│   │   └── terms-modal.html
│   ├── main.js           # Main JavaScript file
│   ├── style.css         # Styles
│   ├── index.html        # HTML template
│   ├── privacy.html      # Legacy privacy page (deprecated)
│   └── terms.html        # Legacy terms page (deprecated)
├── Dockerfile            # Docker build configuration
├── docker-compose.yml    # Docker compose configuration
├── nginx.conf           # Nginx configuration for production
├── build.sh             # Build script
├── package-docker-image.sh # Docker image packaging script
├── check-architecture.sh    # Architecture compatibility checker
├── package.json
└── README.md
```

## Architecture Changes

### HTML Componentization

The original monolithic `main.js` file with inline HTML has been refactored into separate component files for better maintainability:

- **Before**: All HTML was embedded in `main.js` as a template string
- **After**: HTML is split into logical components loaded dynamically

### Modal Implementation

Privacy policy and terms of use are now displayed in modals instead of separate pages:

- **Before**: Links navigated to `privacy.html` and `terms.html`
- **After**: Links open modal dialogs with the same content

## Deployment Options

### Option 1: Direct Docker Run

```bash
docker build -t float-island-website .
docker run -p 8080:4173 float-island-website
```

### Option 2: Docker Compose (Recommended)

```bash
docker-compose up -d
```

### Option 3: Production with Nginx

```bash
docker-compose --profile production up -d
```

## Environment Variables

- `NODE_ENV`: Set to `production` for production builds

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test with `npm run dev`
5. Build with `npm run build`
6. Test Docker build with `./build.sh`
7. Submit a pull request

## Troubleshooting

### Architecture Compatibility Issues

If you encounter `exec /usr/local/bin/docker-entrypoint.sh: exec format error`:

1. **Check server architecture:**
   ```bash
   uname -m  # Should output "x86_64" for x86 systems
   ```

2. **Check Docker image architecture:**
   ```bash
   docker inspect float-island-website:latest --format='{{.Architecture}}'
   # Should output "amd64"
   ```

3. **Rebuild the image for correct architecture:**
   ```bash
   # Clean old images
   docker rmi float-island-website:latest

   # Rebuild for Linux x86
   ./package-docker-image.sh

   # Load on target server
   docker load < float-island-website-image-*.tar
   ```

4. **Alternative: Use Docker buildx for cross-platform builds:**
   ```bash
   docker buildx build --platform linux/amd64 -t float-island-website:latest .
   ```

### Common Docker Issues

- **Container exits immediately:** Check logs with `docker logs <container_id>`
- **Port already in use:** Change port mapping, e.g., `-p 8081:4173`
- **Permission denied:** Make sure Docker daemon is running and user has permissions

## License

© 2025 Float Island Inc. All rights reserved.