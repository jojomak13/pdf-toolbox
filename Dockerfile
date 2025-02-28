# Build stage
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Install necessary build tools
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:3.19

WORKDIR /app

# Create storage directory
RUN mkdir -p storage \
    && chmod -R 777 storage

# Add necessary runtime dependencies
RUN apk add --no-cache \
    ca-certificates \
    chromium \
    harfbuzz \
    nss \
    freetype \
    ttf-freefont \
    font-noto-emoji \
    wqy-zenhei

# Set environment variables for Chrome
ENV CHROME_BIN=/usr/bin/chromium-browser \
    CHROME_PATH=/usr/lib/chromium/

# Set default environment variables
ENV APP_NAME="PDF-Toolbox" \
    PORT=8080 \
    OUTPUT_FILE_NAME="out.pdf"

# S3 configuration will be provided at runtime
# S3_KEY
# S3_SECRET
# S3_REGION
# S3_BUCKET

# Copy binary from builder
COPY --from=builder /app/main .

# Expose port (using the PORT environment variable)
EXPOSE ${PORT}

# Run the application
CMD ["./main"]
