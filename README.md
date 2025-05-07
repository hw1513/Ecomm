# E-Commerce Microservices Platform

This is a microservices-based e-commerce platform built with Go. The project follows a clean architecture pattern and uses gRPC for inter-service communication.

## Project Structure

```
.
├── api/            # API definitions (Protocol Buffers, OpenAPI)
├── cmd/            # Application entry points
├── internal/       # Private application code
│   ├── common/     # Shared utilities and configurations
│   └── order/      # Order service implementation
├── scripts/        # Build and utility scripts
└── Makefile       # Build automation
```

## Prerequisites

- Go 1.24.2 or later
- Protocol Buffers compiler
- gRPC tools

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/hw1513/ecomm.git
   cd ecomm
   ```

2. Install dependencies:

   ```bash
   make deps
   ```

3. Generate code from proto definitions:

   ```bash
   make generate
   ```

4. Build the project:
   ```bash
   make build
   ```

## Available Make Commands

- `make generate` - Generate code from proto and OpenAPI definitions
- `make build` - Build the project
- `make test` - Run tests
- `make clean` - Clean generated files
- `make deps` - Install dependencies
- `make help` - Show help message

## Services

### Order Service

The Order Service handles order management functionality. It exposes both HTTP and gRPC endpoints.

## Configuration

Configuration is managed through YAML files located in `internal/common/config/`. Different environments (development, production) have their own configuration files.

## Development

1. Make sure all dependencies are installed:

   ```bash
   make deps
   ```

2. Generate code after making changes to proto files:

   ```bash
   make generate
   ```

3. Run tests:
   ```bash
   make test
   ```

## License

[Add your license information here]

## Contributing

[Add contribution guidelines here]
