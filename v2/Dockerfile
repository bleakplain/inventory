# Use the official Golang image as the base image
FROM golang:1.16 as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN make build

# Start a new stage from the scratch image
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/inventory-service /inventory-service

# Copy the config.yaml file
COPY --from=builder /app/config.yaml /config.yaml

# Expose the ports for HTTP, gRPC, and Thrift
EXPOSE 8080 9000

# Run the inventory service
ENTRYPOINT ["/inventory-service"]
