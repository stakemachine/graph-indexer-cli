FROM golang:1.16 as builder
RUN apt update && \
    apt-get install ca-certificates tzdata -y && \
    update-ca-certificates
ENV USER=appuser
ENV UID=10001

RUN adduser --disabled-password --gecos "" --home "/nonexistent" --shell "/sbin/nologin" --no-create-home --uid "${UID}" "${USER}"


COPY . /go/src/github.com/stakemachine/graph-indexer-cli
WORKDIR /go/src/github.com/stakemachine/graph-indexer-cli
RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o /go/bin/graph-indexer-cli .

FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/graph-indexer-cli /go/bin/graph-indexer-cli

USER appuser:appuser
WORKDIR /

ENTRYPOINT ["/go/bin/graph-indexer-cli"]