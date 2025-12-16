#build stage
#golang:alpine AS builder
FROM golang:1.25

WORKDIR /usr/src/app

# Install git in case dependencies need it
# RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v -o bot ./cmd/bot

#final stage
#FROM alpine:latest

#WORKDIR /app

#COPY --from=builder /go/src/github.com/alishcodes/go-discord-bot/bot .

#RUN adduser -D botuser
#USER botuser

CMD ["./bot"]
