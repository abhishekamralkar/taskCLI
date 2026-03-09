.PHONY: build run clean docker-build docker-run docker-clean help fmt test trivy-scan sbom

# Variables
BINARY_NAME=taskcli
DOCKER_IMAGE=taskcli:latest
DOCKER_CONTAINER=taskcli-container
SCAN_REPORT=trivy-report.json
SBOM_CYCLONEDX=sbom-cyclonedx.json
SBOM_SPDX=sbom-spdx.json

help:
	@echo "Available targets:"
	@echo "  build         - Build the binary"
	@echo "  run           - Run the application"
	@echo "  clean         - Clean build artifacts"
	@echo "  fmt           - Format code"
	@echo "  docker-build  - Build Docker image"
	@echo "  docker-run    - Run Docker container"
	@echo "  docker-clean  - Remove Docker image and container"
	@echo "  trivy-scan    - Scan Docker image for vulnerabilities"
	@echo "  sbom          - Generate SBOM (CycloneDX and SPDX formats)"
	@echo "  test          - Run tests (if available)"

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) .
	@echo "✓ Build complete"

run: build
	@./$(BINARY_NAME) -list

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@echo "✓ Cleanup complete"

fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "✓ Code formatted"

docker-build:
	@echo "Building Docker image..."
	@sudo docker build -t $(DOCKER_IMAGE) .
	@echo "✓ Docker image built: $(DOCKER_IMAGE)"

docker-run: docker-build
	@echo "Running Docker container..."
	@sudo docker run --name $(DOCKER_CONTAINER) --rm $(DOCKER_IMAGE)

docker-clean:
	@echo "Cleaning Docker artifacts..."
	@sudo docker rm -f $(DOCKER_CONTAINER) 2>/dev/null || true
	@sudo docker rmi $(DOCKER_IMAGE) 2>/dev/null || true
	@echo "✓ Docker cleanup complete"

test:
	@echo "Running tests..."
	@go test -v ./...

trivy-scan: docker-build
	@echo "Running Trivy scan on $(DOCKER_IMAGE)..."
	@sudo trivy image --format json --output $(SCAN_REPORT) $(DOCKER_IMAGE)
	@echo "✓ Scan complete. Report saved to $(SCAN_REPORT)"
	@echo ""
	@echo "Vulnerability Summary:"

sbom: docker-build
	@echo "Generating SBOM for $(DOCKER_IMAGE)..."
	@sudo trivy image --format cyclonedx --output $(SBOM_CYCLONEDX) $(DOCKER_IMAGE)
	@echo "✓ CycloneDX SBOM generated: $(SBOM_CYCLONEDX)"
	@sudo trivy image --format spdx-json --output $(SBOM_SPDX) $(DOCKER_IMAGE)
	@echo "✓ SPDX SBOM generated: $(SBOM_SPDX)"
	@sudo trivy image --severity CRITICAL,HIGH $(DOCKER_IMAGE)
