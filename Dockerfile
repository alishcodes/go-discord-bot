#build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o bot ./cmd/bot

#final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bot .

RUN adduser -D botuser
USER botuser

CMD ["./bot"]
