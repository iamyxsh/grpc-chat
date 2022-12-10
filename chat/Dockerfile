# Choose whatever you want, version >= 1.16
FROM golang:1.19-alpine

WORKDIR /go/src/github.com/iamyxsh/grpc-chat/chat

RUN go install github.com/cosmtrek/air@latest

COPY go.mod ./

CMD ["air"]