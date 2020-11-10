package main

import (
	"log"
	"net"

	"github.com/JavaHutt/grpc-adder/pkg/api"
	"github.com/JavaHutt/grpc-adder/pkg/api/adder"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
