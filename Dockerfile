# Stage 1: Build the Go app
FROM golang:1.19-alpine AS builder
RUN export GOPROXY=direct
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Build the Go app statically
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Stage 2: Final image with minimal dependencies
FROM ubuntu

# Copy the statically-linked Go binary from the builder stage
COPY --from=builder /app/main /app/main

# Expose the required port (replace with your app's port)
EXPOSE 8080
# Set the entrypoint to run the Go binary
ENTRYPOINT ["/app/main"]