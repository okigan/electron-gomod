// go get -u google.golang.org/protobuf/cmd/protoc-gen-go
// go install google.golang.org/protobuf/cmd/protoc-gen-go

// go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

//go:generate protoc -I=./proto service.proto --go_out=./gen --go_opt=paths=source_relative
//go:generate protoc -I=./proto service.proto --go-grpc_out=./gen
//go:generate protoc -I=./proto service.proto --js_out=import_style=commonjs:./gen
//go:generate protoc -I=./proto service.proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./gen

package main

import "C"

//export Hello
func Hello() *C.char {
	return C.CString("Hello world!")
}

// required to build
func main() {
}

// https://github.com/charlieduong94/node-golang-native-addon-experiment
// maybe switch to https://www.electronjs.org/docs/tutorial/using-native-node-modules
// http://blog.cinan.sk/2018/02/22/integrate-native-node-dot-js-modules-into-an-electron-app-1-slash-2/
