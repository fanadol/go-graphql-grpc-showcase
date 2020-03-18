package main

import (
	"log"
	"net"

	"github.com/fanadol/graphql-go-example/internal/constants"
	"github.com/fanadol/graphql-go-example/internal/storage/memory"
	"github.com/fanadol/graphql-go-example/internal/user"
	"google.golang.org/grpc"
)

func main() {
	memorystorage := new(memory.Storage)
	userservice := user.NewService(memorystorage)

	grpcsvc := grpc.NewServer()

	user.NewGRPCHandler(grpcsvc, userservice)

	log.Println("Starting GRPC at port", constants.GRPCPort)

	l, err := net.Listen("tcp", constants.GRPCPort)
	if err != nil {
		panic(err.Error())
	}

	log.Fatal(grpcsvc.Serve(l))
}
