# Use the official Golang image as the base image
FROM golang:1.16 as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o inventory-service cmd/main.go

# Start a new stage with a minimal image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/inventory-service .

# Copy the config file
COPY config.yaml .

# Expose the necessary ports
EXPOSE 8000 9000 10000

# Run the application
CMD ["./inventory-service"]
