PROJ=dex-cli
ORG_PATH=github.com/PwC-ASD
REPO_PATH=$(ORG_PATH)/$(PROJ)
export PATH := $(PWD)/bin:$(PATH)

VERSION ?= $(shell ./scripts/git-version)

DOCKER_REPO=cnero/dex-cli
#DOCKER_IMAGE=$(DOCKER_REPO):$(VERSION)
DOCKER_IMAGE=$(DOCKER_REPO)

$( shell mkdir -p bin )

user=$(shell id -u -n)
group=$(shell id -g -n)

export GOBIN=$(PWD)/bin

LD_FLAGS="-w -X $(REPO_PATH)/version.Version=$(VERSION)"

build: bin/dex-cli

bin/dex-cli: check-go-version
	@go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)

.PHONY: release-binary
release-binary:
	@go build -o /go/bin/dex-cli -v -ldflags $(LD_FLAGS) $(REPO_PATH)

.PHONY: revendor
revendor:
	@glide up -v
	@glide-vc --use-lock-file --no-tests --only-code

vet:
	@go vet $(shell go list ./... | grep -v '/vendor/')

fmt:
	@./scripts/gofmt $(shell go list ./... | grep -v '/vendor/')

lint:
	@for package in $(shell go list ./... | grep -v '/vendor/' | grep -v '/api' | grep -v '/server/internal'); do \
      golint -set_exit_status $$package $$i || exit 1; \
	done

.PHONY: docker-image
docker-image:
	@sudo docker build -t $(DOCKER_IMAGE) .

.PHONY: check-go-version
check-go-version:
	@./scripts/check-go-version

clean:
	@rm -rf bin/

FORCE:

.PHONY: vet fmt lint
