#build stage
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o bot ./cmd/bot

#final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bot .

RUN adduser -D botuser
USER botuser

CMD ["./bot"]
