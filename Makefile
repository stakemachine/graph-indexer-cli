GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
GIT_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
GIT_TAG    ?= $(shell git describe --tags)

docker:
	DOCKER_BUILDKIT=1 docker build -t graph-indexer-cli:latest --build-arg git_commit='${GIT_COMMIT}' --build-arg git_branch='${GIT_BRANCH}' --build-arg git_tag='${GIT_TAG}' .
check:
	golangci-lint run
test:
	go test -v ./...