
# Use a minimal base image
FROM alpine:latest

# Set environment variables
ENV NODE_VERSION=18.16.1
ENV GO_VERSION=1.22.3
ENV PATH="/usr/local/go/bin:${PATH}"

# Install dependencies
RUN apk add --no-cache \
    curl \
    git \
    build-base \
    ca-certificates \
    bash

# Install Node.js
RUN apk add --no-cache nodejs npm

# Remove any previous Go installation
RUN rm -rf /usr/local/go

# Install Go
RUN curl -fsSL https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xz \
    && ln -s /usr/local/go/bin/go /usr/local/bin/go

# Set the working directory
WORKDIR /app

# Copy the project files
COPY . .

# Install Node.js dependencies and build the project
RUN npm install && npm run build

# Build the Go application
RUN go build -tags netgo -ldflags '-s -w' -o app

# Expose port 8000
EXPOSE 8080

# Start the application
CMD ["./app"]
