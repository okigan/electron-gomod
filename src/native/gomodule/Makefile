default: compile-go

setup-deps:
	brew install protobuf
	brew install protoc-gen-grpc-web

generate:
	mkdir -p ./gen
	PATH="${PATH}:$(go env GOPATH)/bin" go generate

compile-go:
	go build -buildmode=c-archive -o ./build/gomodule.so
