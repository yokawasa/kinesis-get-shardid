.PHONY: clean all kinesis-get-shardid

.DEFAULT_GOAL := all

TARGETS=kinesis-get-shardid

CUR := $(shell pwd)
OS := $(shell uname)
VERSION := $(shell cat ${CUR}/VERSION)

kinesis-get-shardid:
	golint ${CUR}
	GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=${VERSION}" -o ${CUR}/dist/kinesis-get-shardid_linux_amd64 ${CUR}/src
	GOOS=darwin GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=${VERSION}" -o ${CUR}/dist/kinesis-get-shardid_darwin_amd64 ${CUR}/src
	GOOS=windows GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=${VERSION}" -o ${CUR}/dist/kinesis-get-shardid_windows_amd64 ${CUR}/src

all: $(TARGETS)

clean:
	rm -rf dist
