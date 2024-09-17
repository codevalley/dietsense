# Use the official Go image as a base image for the build stage
FROM golang:1.21 AS builder

WORKDIR /app

# Copy the Go mod and sum files, and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code into the container
COPY . .

# Change to the directory containing main.go
WORKDIR /app/cmd

# Build the application
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ../main .

# Start a new stage from scratch using a minimal base image
FROM debian:bookworm-slim

# Set the working directory
WORKDIR /app

# Install CA certificates and runtime dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    sqlite3 \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Create a non-root user and group with UID and GID 1000
RUN groupadd -g 1000 dietuser && \
    useradd -u 1000 -g dietuser dietuser

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Create the /app/data directory and set permissions
RUN mkdir -p /app/data && \
    chown -R dietuser:dietuser /app && \
    chmod -R 755 /app

# Switch to non-root user
USER dietuser

# Expose the port the app runs on
EXPOSE 8080

# Include a healthcheck (optional)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1

# Set the entry point to the application executable
ENTRYPOINT ["/app/main"]