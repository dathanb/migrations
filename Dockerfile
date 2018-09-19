FROM golang:1.11

RUN mkdir -p /go/src/github.com/dathanb/fakestack
COPY . /go/src/github.com/dathanb/fakestack
ENV GOPATH /go
WORKDIR /go/src/github.com/dathanb/fakestack
RUN make build

FROM golang:1.11

RUN mkdir /app
WORKDIR /app
COPY --from=0 /go/src/github.com/dathanb/fakestack/fakestack /app/

CMD ["./fakestack", "start"]
