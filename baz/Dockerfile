FROM golang:1.10 AS builder

MAINTAINER Mark St.Godard <markstgodard@gmail.com>

WORKDIR /go/src/github.com/markstgodard/baz

COPY . /go/src/github.com/markstgodard/baz

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -i -o /baz

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=builder /baz /app/

CMD /app/baz

EXPOSE 9000
