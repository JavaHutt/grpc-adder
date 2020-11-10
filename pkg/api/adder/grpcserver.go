package adder

import (
	"context"

	"github.com/JavaHutt/grpc-adder/pkg/api"
)

// GRPCServer ...
type GRPCServer struct{}

// Add ...
func (server *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
