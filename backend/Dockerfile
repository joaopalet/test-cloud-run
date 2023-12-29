# Use the official golang image with a specific version as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code from your host to the container's working directory
COPY . .

# Build the Go application inside the container
RUN go build -o test-app

# Expose the port the application runs on
EXPOSE 8080

# Set the entry point for the container
CMD ["./test-app"]
