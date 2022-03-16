FROM golang:1.17.8-alpine as builder

COPY . /go/src/github.com/stakemachine/graph-indexer-cli
WORKDIR /go/src/github.com/stakemachine/graph-indexer-cli
RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o /go/bin/graph-indexer-cli .

FROM gcr.io/distroless/static
COPY --from=builder /go/bin/graph-indexer-cli /go/bin/graph-indexer-cli

USER nonroot:nonroot
WORKDIR /

ENTRYPOINT ["/go/bin/graph-indexer-cli"]