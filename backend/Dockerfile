# ---------- STAGE 1: Build ----------
FROM golang:1.25-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Copy go mod files first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the backend code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# ---------- STAGE 2: Runtime ----------
FROM alpine:latest

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/server .

# Expose backend port
EXPOSE 8080

# Run the binary
CMD ["./server"]
