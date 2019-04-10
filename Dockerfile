FROM golang:1.10-stretch

WORKDIR /go/src/lfs

COPY main.go ./
COPY stock ./stock/

RUN go get ./...
RUN go install
WORKDIR /go/src/lfs/stock
RUN go test
WORKDIR /go/src/lfs

