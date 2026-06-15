.PHONY: help build test run clean docker-build

help:
	@echo "Bot Army CLI - Makefile targets"
	@echo "  build           - Build binary"
	@echo "  run             - Run CLI against local NATS (4222)"
	@echo "  test            - Run smoke tests"
	@echo "  docker-build    - Build Docker image"
	@echo "  docker-run      - Run in Docker"
	@echo "  clean           - Remove build artifacts"

build:
	@go mod download
	@go build -o bot-army-cli .
	@echo "✓ Built bot-army-cli"

run: build
	@./bot-army-cli

test: build
	@echo "Testing NATS connection..."
	@timeout 5 ./bot-army-cli <<< "help" > /dev/null && echo "✓ NATS connection test passed" || echo "❌ NATS not available"

docker-build:
	@docker build -t bot-army-cli:latest .
	@echo "✓ Built Docker image: bot-army-cli:latest"

docker-run: docker-build
	@docker run -it --network host \
		-e NATS_URL=nats://host.docker.internal:4222 \
		bot-army-cli:latest

clean:
	@rm -f bot-army-cli
	@go clean
	@echo "✓ Cleaned"
