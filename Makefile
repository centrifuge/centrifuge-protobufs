all: install proto

.PHONY: help install proto-lint proto
help: ## Show this help message.
	@echo 'usage: make [target] ...'
	@echo
	@echo 'targets:'
	@egrep '^(.+)\:\ ##\ (.+)' ${MAKEFILE_LIST} | column -t -c 2 -s ':#'

install: ## Install dependencies required to generate bindings & documentation
	@go install github.com/ckaznocha/protoc-gen-lint@0.2.4
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@1.2.0

PROTO_FILES = $(shell  find . -maxdepth 2 -name '*.proto')

proto-lint: ## Lint protos
	@protoc -I. \
    		-Ivendor/github.com/centrifuge \
    		--lint_out=. \
    		${PROTO_FILES}

OUTDIR = gen/go

proto: ## Generate go bindings
	@rm -rf ${OUTDIR}
	@mkdir -p ${OUTDIR}
	@protoc -I. \
		-Ivendor/github.com/centrifuge \
		--go_out=paths=source_relative:${OUTDIR} \
		--go-grpc_out=paths=source_relative:${OUTDIR} \
		${PROTO_FILES}