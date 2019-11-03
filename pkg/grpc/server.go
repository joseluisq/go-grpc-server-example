package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	pb "go-grpc-server/helloworld"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// RunServer runs grpc server
func RunServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return err
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			// Sig is a ^C, handle it
			log.Println("Shutting down gRPC server...")

			s.GracefulStop()

			<-ctx.Done()
		}
	}()

	// Start gRPC server
	log.Println("Starting gRPC server at", port)

	return s.Serve(lis)
}
