# Use the official Golang image as the base image
FROM golang:1.16 as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the application
RUN make build

# Start a new stage from the scratch image
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/bin/inventory-service /inventory-service

# Copy the configuration file
COPY --from=builder /app/config.yaml /config.yaml

# Expose the necessary ports
EXPOSE 8000 9000 10000

# Run the application
CMD ["/inventory-service", "-conf=/config.yaml"]
