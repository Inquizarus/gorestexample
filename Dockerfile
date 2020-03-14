FROM golang:alpine as builder

RUN apk update && apk add ca-certificates && apk add make

COPY . $GOPATH/src/github.com/inquizarus/gorestexample/

WORKDIR $GOPATH/src/github.com/inquizarus/gorestexample
RUN CGO111MODULE=on make build-linux
RUN mv gorestexample_unix /go/bin/gorestexample

FROM busybox:latest
COPY --from=builder /go/bin/restexample /go/bin/restexample
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
EXPOSE 8080
WORKDIR /go/bin
CMD ./restexample
