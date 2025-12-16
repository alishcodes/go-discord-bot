#build stage
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bot ./cmd/bot

#final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bot .

RUN adduser -D botuser
USER botuser

CMD ["./bot"]
