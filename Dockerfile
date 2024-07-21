FROM golang:1.22.3-alpine AS builder

COPY .  /github.com/ruknez/chat-server
WORKDIR  /github.com/ruknez/chat-server

RUN go mod download
RUN go build -o ./bin/chat_server cmd/chat_server/app.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/ruknez/chat-server/bin/chat_server .

CMD ["./chat_server"]
