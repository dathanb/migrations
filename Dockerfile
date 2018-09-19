FROM golang:1.11

RUN mkdir -p /go/src/github.com/udacity/migration-demo
COPY . /go/src/github.com/udacity/migration-demo/
ENV GOPATH /go
WORKDIR /go/src/github.com/udacity/migration-demo
RUN make build

FROM golang:1.11

RUN mkdir /app
WORKDIR /app
COPY --from=0 /go/src/github.com/udacity/migration-demo/migration-demo /app/

CMD ["./migration-demo", "start"]
