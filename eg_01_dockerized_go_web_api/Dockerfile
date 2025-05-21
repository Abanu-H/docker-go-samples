# Use official Go image with version 1.21.7
FROM golang:1.21.7 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy dependency definitions
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the entire application source code
COPY . .

# Build a statically linked Linux binary for Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o dockerized_go_web_api .

# ----------- Run Stage (minimal image) -----------
FROM alpine:latest

# Copy the statically linked binary from the build stage
COPY --from=builder /app/dockerized_go_web_api /dockerized_go_web_api

# Set working directory (optional)
WORKDIR /app

# Expose the port your app listens on
EXPOSE 3000

# Run the app
CMD ["/dockerized_go_web_api"]
