# This image includes Go and necessary build tools.
FROM golang:1.21.7 AS builder

# Set the working directory inside the container to /app.
# All subsequent commands will be run from this directory.
WORKDIR /app

# Copy all files from the current host directory into the /app directory in the container.
COPY . /app

# Download all Go module dependencies as defined in go.mod/go.sum.
RUN go mod download

# Expose port 3000 so it can be accessed from outside the container.
# This does not publish the port but documents the intended port.
EXPOSE 3000

# Define the command to run the application.
# This runs the main.go file using `go run`.
CMD ["go", "run", "main.go"]
