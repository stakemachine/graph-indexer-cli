GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
GIT_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
GIT_TAG    ?= $(shell git describe --tags)

.PHONY: all build clean test coverage bench fmt vet docker check tidy

all: check test build

build:
	go build -ldflags "-X main.gitCommit=${GIT_COMMIT} -X main.gitBranch=${GIT_BRANCH} -X main.gitTag=${GIT_TAG}" -o graph-indexer-cli

clean:
	rm -f graph-indexer-cli
	rm -f coverage.out

docker:
	DOCKER_BUILDKIT=1 docker build -t graph-indexer-cli:latest --build-arg git_commit='${GIT_COMMIT}' --build-arg git_branch='${GIT_BRANCH}' --build-arg git_tag='${GIT_TAG}' .

check:
	golangci-lint run

test:
	go test -v ./...

coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

bench:
	go test -bench=. -benchmem ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy