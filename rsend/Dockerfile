FROM golang:1.22
RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go build -o rsend-app ./cmd/async_producer

CMD ["./rsend-app"]