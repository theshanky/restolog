FROM golang:1.17.5-alpine3.15 AS build

WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Run tests
RUN go test resto/resto_test.go resto/resto.go
# Build the application
RUN go build -o main .

# Create a new Docker image based on a lightweight Alpine Linux image
FROM alpine:3.15

WORKDIR /app

# Copy the built binary from the build image
COPY --from=build /app/main .

# Set the entrypoint command to run the binary
ENTRYPOINT ["./main"]
