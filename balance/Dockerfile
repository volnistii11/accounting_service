FROM golang:1.22
RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go build -o balance-app ./cmd/balance

CMD ["./balance-app"]