# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)

GITHUB_USERNAME=alokic
APPNAME=sdkgen
PROJECT_ROOT=${GOPATH}/src/github.com/${GITHUB_USERNAME}/${APPNAME}
SCRIPT_FOLDER=${PROJECT_ROOT}/scripts
CMD_FOLDER=${PROJECT_ROOT}/cmd
GOBIN=${GOPATH}/bin
GITCOMMIT=$(shell git rev-parse --short HEAD)
GITBRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GITVERSION=$(shell git describe --abbrev=4 --dirty --always --tags)
LDFLAGS=-s -w -X main.VERSION=${GITVERSION} -X main.GITCOMMIT=${GITCOMMIT} -X main.GITBRANCH=${GITBRANCH}

.PHONY: all
all: clean fmt test vet linux darwin 	## Builds the project.


.PHONY: test
test: 	## Tests the project except vendor and deployment folders
	@go test -v  $(shell go list ./... | grep -v /vendor/ | grep -v /deployment/ | grep -v /output/ ) 

.PHONY: lint
lint:														## lints the project except vendor and deployment folders
	@golint $(shell go list ./... | grep -v /vendor/ | grep -v /deployment/ |  grep -v /output/) | grep -v '.pb.go:' | tee /dev/stderr


.PHONY: vet
vet:														## Vets the project except vendor and deployment folders
	@go vet $(shell go list ./... | grep -v /vendor/ | grep -v /deployment/ |  grep -v /output/) | grep -v '.pb.go:' | tee /dev/stderr


.PHONY: fmt
fmt:														## Formats the project except vendor and deployment folders
	@go fmt  $(shell go list ./... | grep -v /vendor/ | grep -v /deployment/ |  grep -v /output/ | grep -v '.pb.go:') 


.PHONY: cover
cover: ## Runs go test with coverage
	@echo "" > coverage.txt
	@for d in $(shell go list ./... | grep -v /vendor/ | grep -v /deployment/ |  grep -v /output/ | grep -v '.pb.go:'); do \
		go test -race -coverprofile=profile.out "$$d"; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.txt; \
			rm profile.out; \
			rm coverage.txt; \
		fi; \
	done;


.PHONY: clean
clean:														## Clean any stray files formed during make


.PHONY: tag
tag: checkversion ## Create a new git tag to prepare to build a release
	git tag -sa $(GITVERSION) -m "$(GITVERSION)"
	@echo "Run git push origin $(GITVERSION) to push your new tag to GitHub and trigger a travis build."


.PHONY: build
build: sdkgen ctl ## Installs $APPNAME in $GOPATH/bin

.PHONY: sdkgen
sdkgen:		## Builds sdkgen - sdk generator. try running: sdkgen 
	@rm ${GOBIN}/sdkgen  2>/dev/null|| true
	@env GOARCH=amd64 GOGC=off go build -ldflags="${LDFLAGS}" -i -o ${GOBIN}/${APPNAME} ${CMD_FOLDER}/${APPNAME}/main.go

.PHONY: ctl
ctl:		## Builds ctl - golang based utilities for command line processing. try running: ctl 
	@rm ${GOBIN}/ctl  2>/dev/null|| true
	@env GOARCH=amd64 GOGC=off go build -ldflags="${LDFLAGS}" -i -o ${GOBIN}/ctl ${CMD_FOLDER}/ctl/main.go

.PHONY: sdk
sdk: ## Builds sdk in a programming language.  syntax: make sdk LANG=python


.PHONY: help
help:  ## Print help
	@echo "=================================================="
	@echo "Run: make <target_name> NAMESPACE=<namespace_name>"
	@echo "=================================================="
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: checkversion 
checkversion:
ifeq ($(GITVERSION),)
	@echo "Missing GITVERSION"
	@exit 1
endif