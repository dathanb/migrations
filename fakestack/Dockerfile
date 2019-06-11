FROM golang:1.11

RUN mkdir -p /go/src/github.com/dathanb/migrations/fakestack
COPY . /go/src/github.com/dathanb/migrations/fakestack
ENV GOPATH /go
WORKDIR /go/src/github.com/dathanb/migrations/fakestack
RUN make build

FROM golang:1.11

RUN mkdir /app
WORKDIR /app
COPY --from=0 /go/src/github.com/dathanb/migrations/fakestack/fakestack /app/

CMD ["./fakestack", "start"]
