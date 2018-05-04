package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	pb "github.com/alextanhongpin/go-video/proto"
	"google.golang.org/grpc"
)

const port = ":9090"

type server struct{}

func (s *server) GetVideos(ctx context.Context, in *pb.GetVideosRequest) (*pb.GetVideosResponse, error) {
	return &pb.GetVideosResponse{
		Query: in.Query,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen to: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVideoServiceServer(s, &server{})

	// Register reflection service on gRPC server
	reflection.Register(s)
	log.Printf("listening to port *%v. press ctrl+c to cancel.\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
