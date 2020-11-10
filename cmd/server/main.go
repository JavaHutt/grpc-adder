package main

import (
	"github.com/JavaHutt/grpc-adder/pkg/api"
	"github.com/JavaHutt/grpc-adder/pkg/api/adder"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)
}
