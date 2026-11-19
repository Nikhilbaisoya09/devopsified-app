
# Stage 1: Build Go application
FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath -ldflags="-s -w" \
    -o /app/devopsified ./cmd/server

# Stage 2: Production runtime
FROM gcr.io/distroless/static-debian13:nonroot

WORKDIR /

COPY --from=builder /app/devopsified /devopsified

ENV PORT=8080
ENV APP_ENV=production
ENV APP_VERSION=1.0.0

EXPOSE 8080

USER 65532:65532

ENTRYPOINT ["/devopsified"]
