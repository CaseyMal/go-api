# Use a minimal Go base image
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Copy Go files
COPY . .

# Build the Go binary
RUN go build -o go-api

# Run the binary
CMD ["./go-api"]

