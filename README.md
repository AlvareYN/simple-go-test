# Test Binary

A simple Go application that prints "Hello World" with version information in an infinite loop.

## Application

The application runs an infinite loop that prints "Hello World" along with the binary version every 5 seconds.

## Version

Current version: **1.0.0**

## Building and Running

### Local Development

```bash
go run main.go
```

### Docker

Build the image:
```bash
docker build -t test-binary:1.0.0 .
```

Run the container:
```bash
docker run test-binary:1.0.0
```

## GitHub Actions Workflow

The project includes a GitHub Actions workflow that automatically builds and pushes Docker images to Docker Hub.

### Setup

1. Create a Docker Hub account if you don't have one
2. Go to your GitHub repository settings
3. Add the following secrets:
   - `DOCKER_USERNAME`: Your Docker Hub username
   - `DOCKER_PASSWORD`: Your Docker Hub password or access token

### Triggering the Workflow

The workflow runs automatically on:
- Pushes to the `main` branch
- Tag pushes (e.g., `v1.0.0`)
- Manual trigger via workflow_dispatch

### Image Tags

The workflow creates multiple tags:
- Branch name (e.g., `main`)
- Git SHA (e.g., `sha-abc1234`)
- Semver tags for version tags (e.g., `1.0.0`, `1.0`, `1`)

## Files

- `main.go` - The main application code
- `go.mod` - Go module definition
- `Dockerfile` - Multi-stage Docker build using Alpine Linux
- `.github/workflows/docker-publish.yml` - GitHub Actions workflow for Docker Hub deployment
