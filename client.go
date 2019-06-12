package main

import (
	"context"
	"log"

	"github.com/yukpiz/go-protobuf-example/pb"
	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer con.Close()
	client := pb.NewUserClient(con)
	res, err := client.GetUser(context.Background(), &pb.UserMessage{
		Id: 1,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response: %+v\n", res)
}
