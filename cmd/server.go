package cmd

import (
	"context"

	grpc "go-grpc-server/pkg/grpc"
)

// RunServer runs grpc server
func RunServer() error {
	ctx := context.Background()

	// TODO: Run other processes here

	return grpc.RunServer(ctx)
}
