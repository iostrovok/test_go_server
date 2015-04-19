GP := $(shell dirname $(realpath $(lastword $(GOPATH))))
ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
export GOPATH := ${ROOT}:${GOPATH}

install:
	go get -u github.com/bmizerany/pat
	go get -u gopkg.in/check.v1

build:
	go build -a -v ./src/scripts/run.go
	go build -a -v ./src/test_script/add_word_file.go
	go build -a -v ./src/test_script/add_word_std.go

run:
	go run ./src/scripts/run.go

# Tests
test-store:
	go test ./src/Storage/

test-store-cover:
	go test ./src/Storage/ -cover -coverprofile ./tmp.out; go tool cover -html=./tmp.out -o cover.html; rm ./tmp.out


test-common:
	go test ./src/Common/

test-common-cover:
	go test ./src/Common/ -cover -coverprofile ./tmp.out; go tool cover -html=./tmp.out -o cover.html; rm ./tmp.out


