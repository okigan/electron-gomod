default: compile-go

setup-deps:
	brew install protobuf
	brew install protoc-gen-grpc-web

compile-go:
	mkdir -p ./gen
	PATH="${PATH}:$(go env GOPATH)/bin" go generate
	go build -buildmode=c-archive -o ./build/gomodule.so golang.go
