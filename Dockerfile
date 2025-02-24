FROM golang:1.24.0-alpine AS builder

ARG TARGETARCH=arm64

RUN wget https://github.com/upx/upx/releases/download/v5.0.0/upx-5.0.0-${TARGETARCH}_linux.tar.xz -O /tmp/upx.tar.xz && \
    tar -xJOf /tmp/upx.tar.xz upx-5.0.0-${TARGETARCH}_linux/upx > /usr/local/bin/upx && \
    chmod +x /usr/local/bin/upx

COPY . /go/src/github.com/stakemachine/graph-indexer-cli
WORKDIR /go/src/github.com/stakemachine/graph-indexer-cli
RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -ldflags='-w -s -extldflags "-static"' -a -o /go/bin/graph-indexer-cli .

RUN /usr/local/bin/upx --overlay=strip --best /go/bin/graph-indexer-cli

FROM gcr.io/distroless/static
COPY --from=builder /go/bin/graph-indexer-cli /go/bin/graph-indexer-cli

USER nonroot:nonroot
WORKDIR /

ENTRYPOINT ["/go/bin/graph-indexer-cli"]