.DEFAULT_GOAL := build

init:
	rm -rf go.mod
	go mod init github.com/easyone-jwlee/learning_go/ch2
	go mod tidy
.PHONY:init

fmt: init
# fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build hello.go
.PHONY:build

hey_inst:
	go install github.com/rakyll/hey@latest
.PHONY:hey_install

goimports_inst:
	go install golang.org/x/tools/cmd/goimports@latest
.PHONY:goimports_inst

golint_inst:
	go install golang.org/x/lint/golint@latest
.PHONY:golint_inst

lint_inst:
	# binary will be $(go env GOPATH)/bin/golangci-lint
	sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.53.3

.PHONY:lint_inst