.PHONY: lint test static install uninstall cross
VERSION := $(shell git describe --tags --dirty --always)
BIN_DIR := $(GOPATH)/bin
GOX := $(BIN_DIR)/gox

lint:
	test -z $$(gofmt -s -l cmd/ pkg/)
	go vet ./...

test:
	go test ./...

# Compilation
dev:
	go build -ldflags "-X main.Version=dev-${VERSION}" ./cmd/grr-list

LDFLAGS := '-s -w -extldflags "-static" -X main.Version=${VERSION}'
static:
	CGO_ENABLED=0 GOOS=linux go build -ldflags=${LDFLAGS} ./cmd/grr-list

install:
	CGO_ENABLED=0 go install -ldflags=${LDFLAGS} ./cmd/grr-list

uninstall:
	go clean -i ./cmd/grr-list

$(GOX):
	go get -u github.com/mitchellh/gox
cross: $(GOX)
	CGO_ENABLED=0 gox -output="dist/{{.Dir}}-{{.OS}}-{{.Arch}}" -ldflags=${LDFLAGS} -arch="amd64 arm64 arm" -os="linux" -osarch="darwin/amd64" ./cmd/grr-list

# Docker container
container: static
	docker build -t malcolmholmes/grizzly-lister .

# CI
drone:
	# Render the YAML from the jsonnet
	drone jsonnet --source .drone/drone.jsonnet --target .drone/drone.yml --stream --format
	# Sign the config
	drone --server https://drone.grafana.net sign --save malcolmholmes/grizzly-lister .drone/drone.yml
