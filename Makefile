SHELL             := bash
.SHELLFLAGS       := -eu -o pipefail -c
.DEFAULT_GOAL     := default
.DELETE_ON_ERROR  :
.SUFFIXES         :# Go parameters

GOCMD	:= go
GOTEST 	:= $(GOCMD)	test
SWAGGERCMD = docker run --rm -it -e GOPATH=$${HOME}/go:/go -v $${HOME}:$${HOME} -w $(shell pwd) quay.io/goswagger/swagger:v0.25.0

.PHONY: default
default: clean build

.PHONY: build
build:
	./build.sh

.PHONY: build-local
build-local:
	./build.sh local

.PHONY: install
install:
	./install.sh

.PHONY: test
test:
	$(GOTEST) ./...

test-cover:
	$(GOTEST) ./... -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: swagger-generate
swagger-generate:
	$(SWAGGERCMD) generate client -f ./api/swagger.yml -A anansi -P models.Principal

.PHONY: swagger-validate
swagger-validate:
	$(SWAGGERCMD) validate ./api/swagger.yml