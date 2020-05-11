FROM golang:1.13 AS builder

ADD . /go/src/go-con-client-2

WORKDIR /go/src/go-con-client-2
ENV GO111MODULE=on

RUN go build -mod=vendor -o main

FROM debian:9.9-slim

ENV LANG C.UTF-8
ENV GOPATH /go

COPY --from=builder /go/src/go-con-client-2/main /app/main
WORKDIR /app

RUN chmod +x main

ENTRYPOINT ["./main"]
EXPOSE 8082