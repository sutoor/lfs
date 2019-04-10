FROM golang:1.10-stretch

WORKDIR /go/src/lfscp

COPY main.go ./
COPY stock ./stock/

RUN go get ./...
RUN go install

ENTRYPOINT ["lfscp"]
