FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o stress-test ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/stress-test .
ENTRYPOINT ["./stress-test"]