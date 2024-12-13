# Build Stage
FROM golang:alpine AS builder

# Install necessary dependencies (git)
RUN apk update && apk add --no-cache git

# Set the working directory inside the container
WORKDIR /application

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./

RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application (output will be 'binary')
RUN go build -o binary .

# Final Stage (Minimal runtime image)
FROM alpine:latest

# Install necessary libraries (like ca-certificates for HTTPS requests)
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /application

# Copy the compiled binary from the build stage
COPY --from=builder /application/binary /application/

# Expose the port the app will listen on (if needed)
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["/application/binary"]