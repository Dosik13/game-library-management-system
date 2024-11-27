# Step 1: Build the application
FROM golang:1.23 AS builder

WORKDIR /app

# Copy dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the project files
COPY . .

# Cross-compile for Linux
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Step 2: Use a minimal runtime image
FROM scratch

WORKDIR /app

# Copy the statically built binary
COPY --from=builder /app/cmd/main .
COPY --from=builder /app/.env .env

# Expose the application port
# 8080 is the default port for the application
EXPOSE 8080

# Run the application
CMD ["./main"]
