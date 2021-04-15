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

.PHONY: mockery
mockery:
	mockery -dir cmd/adapters/applications -name VampCloudApplicationsClient -case snake -output ./mocks/adaptersmocks -outpkg adaptersmocks
	mockery -dir cmd/adapters/clusters -name VampCloudClustersClient -case snake -output ./mocks/adaptersmocks -outpkg adaptersmocks
	mockery -dir cmd/adapters/services -name VampCloudServicesClient -case snake -output ./mocks/adaptersmocks -outpkg adaptersmocks
	mockery -dir cmd/adapters/ingresses -name VampCloudIngressesClient -case snake -output ./mocks/adaptersmocks -outpkg adaptersmocks
	mockery -dir cmd/adapters/policies -name VampCloudPoliciesClient -case snake -output ./mocks/adaptersmocks -outpkg adaptersmocks	
	mockery -dir cmd/adapters/releases -name VampCloudReleasesClient -case snake -output ./mocks/adaptersmocks -outpkg adaptersmocks