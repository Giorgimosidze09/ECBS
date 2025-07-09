# Use Go image for building the application
FROM golang:1.23-alpine as builder

# Set the working directory for the build
WORKDIR /build

# Accept a build argument for the service name
ARG SERVICE_NAME

# Copy source
COPY . .

# Set the working directory to the specific service
WORKDIR /build/services/${SERVICE_NAME}

# Download dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o /bin/${SERVICE_NAME}

# Use a lightweight image for the runtime
FROM alpine:latest

# Set the environment variable for the service name
ARG SERVICE_NAME
ENV SERVICE_NAME=${SERVICE_NAME}

# Create a service-specific directory
WORKDIR /${SERVICE_NAME}

# Copy the pre-built binary from the builder stage into the service directory
COPY --from=builder /bin/${SERVICE_NAME} ./

# Expose port 8080 (can be adjusted for each service if needed)
EXPOSE 8080

# Use the environment variable to set the command dynamically
CMD ["sh", "-c", "./$SERVICE_NAME"]
