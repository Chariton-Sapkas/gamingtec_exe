# Use the official Golang image as the build image
FROM golang:1.22.6 as builder

WORKDIR /app

# Copy go.mod and go.sum files first for dependency resolution
COPY go.mod go.sum ./

RUN go mod download

# Copy the source code
COPY . .

# Build the Go application as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Use a smaller base image for the runtime
FROM gcr.io/distroless/base-debian11

# Set the working directory
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main /app/main

# Expose gRPC port and health check port
EXPOSE 5050

# Run the application
ENTRYPOINT [ "/app/main" ]
