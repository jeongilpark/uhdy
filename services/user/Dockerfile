# Stage 1: Build the Go binary
FROM golang:1.24.0-bookworm AS build

# Copy the source code
COPY ./user /app/user
COPY ./utils /app/utils

# Set the working directory
WORKDIR /app/user

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o user-service ./main.go

# Stage 2: Create a minimal image and copy the binary
FROM gcr.io/distroless/base-debian12

# Set the working directory
WORKDIR /

# Copy the binary and any other necessary files
COPY --from=build /app/user/user-service /user-service

# COPY the default configuration
COPY ./utils/config/config.yaml /

# Expose the port the service listens on
EXPOSE 8001

# Command to run the service
ENTRYPOINT ["/user-service"]