# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./build/news_service ./cmd/news_service.go
COPY ./config ./build

FROM alpine:latest
RUN apk add curl \
    && adduser -D appuser
USER appuser
WORKDIR /app
EXPOSE 4000
COPY --from=builder /app/build/ .
CMD ["./news_service", "--config", "core.yaml"]
