# ============================
# Stage 1: Build Go binary
# ============================
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git gcc musl-dev

WORKDIR /build

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o cl_system main.go

# ============================
# Stage 2: Run
# ============================
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /build/cl_system .
COPY --from=builder /build/manifest ./manifest
COPY --from=builder /build/resource ./resource

EXPOSE 8099

CMD ["./cl_system"]
