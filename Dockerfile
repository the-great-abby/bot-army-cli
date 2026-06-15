FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bot-army-cli .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root

COPY --from=builder /app/bot-army-cli .

ENV NATS_URL=nats://nats:4222

ENTRYPOINT ["./bot-army-cli"]
