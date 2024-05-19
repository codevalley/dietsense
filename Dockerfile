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
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../main .

# Start a new stage from scratch using a minimal base image
FROM debian:bookworm-slim

# Set the working directory
WORKDIR /app

# Update CA certificates and install SQLite
RUN apt-get update && \
    apt-get install -y ca-certificates sqlite3 --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*
    
# Create a non-root user and switch to it
RUN useradd -m dietuser && \
    chown dietuser /app
USER dietuser

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Assuming your configuration file is named config.yaml and located in the root of your project
# Make sure this file does not contain sensitive information if it's being included in a public image
COPY --from=builder /app/config.yaml ./config.yaml

# Expose the port the app runs on
EXPOSE 8080

# Include a healthcheck (optional)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD [ "curl", "-f", "http://localhost:8080/health" ] || exit 1

# Set the entry point to the application executable
ENTRYPOINT ["/app/main"]
