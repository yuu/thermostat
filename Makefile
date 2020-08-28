.DEFAULT_GOAL := help
.PHONY: help prepare generate vendor build

COMMAND_LIST := ${MAKEFILE_LIST}

# if .env file exists, including secret infomations
ifneq ("$(wildcard .env)","")
include .env
endif

help:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' ${COMMAND_LIST} | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

prepare:  ## download grpc plugin for golang
	go get -u github.com/golang/protobuf/protoc-gen-go

generate:  ## generate grpc
	protoc -I bto --go_out=plugins=grpc:bto --go_opt=module=github.com/yuu/thermostat bto/ir_service.proto

vendor:
	go mod vendor

build:
	go build -mod=vendor
