package main

import (
	"log"
	"net"

	"github.com/yukpiz/go-protobuf-example/pb"
	"github.com/yukpiz/go-protobuf-example/service"
	"google.golang.org/grpc"
)

func main() {
	log.Println("===> START: gRPC example server")

	port, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatal(err) // panic
	}

	server := grpc.NewServer()
	pb.RegisterUserServer(server, &service.UserService{}) // UserServer interfaceに準拠している型
	log.Println("===> Wait Your Connection...")
	server.Serve(port)
}
