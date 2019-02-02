FROM golang:1.10 AS builder

MAINTAINER Mark St.Godard <markstgodard@gmail.com>

WORKDIR /go/src/github.com/markstgodard/linkerd2bug/foo

COPY . /go/src/github.com/markstgodard/linkerd2bug/foo

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -i -o /foo

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=builder /foo /app/

CMD /app/foo

EXPOSE 9000
