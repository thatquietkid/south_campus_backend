# Start from the official Go image
FROM golang:1.24.0

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 1323 to the outside world
EXPOSE 1323

# Command to run the executable
CMD ["./main"]
