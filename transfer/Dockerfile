FROM golang:1.22
RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go build -o transfer-app ./cmd/transfer

CMD ["./transfer-app"]