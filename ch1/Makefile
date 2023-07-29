.DEFAULT_GOAL := build

init:
	rm -rf go.mod
	go mod init ch1
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