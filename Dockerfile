#build stage
FROM golang:alpine AS builder

WORKDIR /go/src/github.com/alishcodes/go-discord-bot

# Install git in case dependencies need it
RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . /go/src/github.com/alishcodes/go-discord-bot

RUN go build -o bot ./cmd/bot

#final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /go/src/github.com/alishcodes/go-discord-bot/bot .

RUN adduser -D botuser
USER botuser

CMD ["./bot"]
