# GOBIN needs to be set to ensure govendor can actually be found and executed
PATH=$(shell printenv PATH):$(GOBIN)

# If you need to overwrite PROTOTOOL_BIN, you can set this environment variable.
PROTOTOOL_BIN ?=$(shell which prototool)

all: install proto-all

.PHONY: help
help: ## Show this help message.
	@echo 'usage: make [target] ...'
	@echo
	@echo 'targets:'
	@egrep '^(.+)\:\ ##\ (.+)' ${MAKEFILE_LIST} | column -t -c 2 -s ':#'

install: ## Install dependencies required to generate bindings & documentation
install: install_prototool install_dep vendorinstall

install_prototool:
	mkdir -p ~/bin
	PROTOTOOL_VERSION=$(PROTOTOOL_VERSION) ./scripts/install_prototool.sh

install_dep:
	go get -u "github.com/golang/dep/..."
	dep ensure

vendorinstall: ## Installs all protobuf dependencies with go-vendorinstall
	go install github.com/centrifuge/centrifuge-protobufs/vendor/github.com/roboll/go-vendorinstall
	go-vendorinstall github.com/golang/protobuf/protoc-gen-go

lint: ## runs prototool lint
	$(PROTOTOOL_BIN) lint

proto-gen-go: ## generates the go bindings
	$(PROTOTOOL_BIN) gen

proto-all: ## runs prototool all
	$(PROTOTOOL_BIN) all
