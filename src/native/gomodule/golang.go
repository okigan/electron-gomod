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

import (
	"fmt"
	hardwaremonitoring "golangmodule/gen"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

//export Hello
func Hello() *C.char {
	return C.CString("Hello world!")
}

// required to build
func main() {

	fmt.Println("Welcome to streaming HW monitoring")
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		panic(err)
	}
	// Create a gRPC server
	gRPCserver := grpc.NewServer()

	// Create a server object of the type we created in server.go
	s := &Server{}
	hardwaremonitoring.RegisterHardwareMonitorServer(gRPCserver, s)

	go func() {
		log.Fatal(gRPCserver.Serve(lis))
	}()

	// We need to wrap the gRPC server with a multiplexer to enable
	// the usage of http2 over http1
	grpcWebServer := grpcweb.WrapServer(gRPCserver)

	multiplex := grpcMultiplexer{
		grpcWebServer,
	}

	// a regular http router
	r := http.NewServeMux()
	// Load our React application
	webapp := http.FileServer(http.Dir("webapp/hwmonitor/build"))
	// Host the web app at / and wrap it in a multiplexer
	r.Handle("/", multiplex.Handler(webapp))

	// create a http server with some defaults
	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// host it
	// go func() {
	log.Fatal(srv.ListenAndServe())
	// }()

	//gRPCserver.GracefulStop()
}

// https://github.com/charlieduong94/node-golang-native-addon-experiment
// maybe switch to https://www.electronjs.org/docs/tutorial/using-native-node-modules
// http://blog.cinan.sk/2018/02/22/integrate-native-node-dot-js-modules-into-an-electron-app-1-slash-2/

// grpcMultiplexer enables HTTP requests and gRPC requests to multiple on the same channel
// this is needed since browsers dont fully support http2 yet
type grpcMultiplexer struct {
	*grpcweb.WrappedGrpcServer
}

// Handler is used to route requests to either grpc or to regular http
func (m *grpcMultiplexer) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.IsGrpcWebRequest(r) {
			m.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}