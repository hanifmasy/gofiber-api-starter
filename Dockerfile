# ---- Build stage ----
FROM golang:1.22-alpine AS builder

# Enable go modules
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the binary (static build for Alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -o fiber-api main.go

# ---- Runtime stage ----
FROM alpine:3.20

WORKDIR /root/
COPY --from=builder /app/fiber-api .
COPY --from=builder /app/.env . # optional if you want env in container

# For debugging/logging inside container
RUN apk add --no-cache tzdata curl

EXPOSE 8080

CMD ["./fiber-api"]
