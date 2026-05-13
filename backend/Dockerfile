# =============================================================================
# Backend Dockerfile — Go + Echo v5
# Multi-stage build untuk ukuran image minimal dan keamanan optimal
# =============================================================================

# ---------------------------------------------------------------------------
# Stage 1: Builder — Kompilasi binary Go
# ---------------------------------------------------------------------------
FROM golang:1.26-alpine AS builder

WORKDIR /build

# Install dependensi sistem yang diperlukan untuk CGO (jika dibutuhkan)
RUN apk add --no-cache git ca-certificates tzdata

# Copy dependency files terlebih dahulu untuk memanfaatkan Docker layer cache
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy seluruh source code backend
COPY . .

# Build binary API server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X main.version=$(git describe --tags --always --dirty 2>/dev/null || echo dev)" \
    -o /build/api-server \
    ./cmd/api

# Build binary Worker (Asynq)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /build/worker \
    ./cmd/worker

# ---------------------------------------------------------------------------
# Stage 2: Runner — Image produksi minimal
# ---------------------------------------------------------------------------
FROM alpine:3.22 AS runner

WORKDIR /app

# Install ca-certificates dan timezone data
RUN apk add --no-cache ca-certificates tzdata && \
    addgroup -S appgroup && \
    adduser -S appuser -G appgroup

# Copy binary dari builder stage
COPY --from=builder /build/api-server /app/api-server
COPY --from=builder /build/worker /app/worker

# Copy file konfigurasi jika ada
COPY --from=builder /build/.env.example /app/.env.example

# Atur kepemilikan file
RUN chown -R appuser:appgroup /app

# Jalankan sebagai non-root user
USER appuser

# Expose port API (default Echo)
EXPOSE 8080

# Healthcheck
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD wget -qO- http://localhost:8080/health || exit 1

# Default: jalankan API server
CMD ["/app/api-server"]
