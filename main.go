package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/dezkoat/xcntn/proto"
)

var (
	port = flag.Int("port", 50001, "Server port")
)

type contentServer struct {
	pb.UnimplementedContentServer
}

func (s *contentServer) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	ts := timestamppb.Now()

	return &pb.Post{
		Id:    "12355",
		Title: "Dummy title",
		Text:  "Dummy text",
		Metadata: &pb.Metadata{
			CreatedBy: "Dean",
			CreatedAt: ts,
			UpdatedBy: "asd",
			UpdatedAt: ts,
		},
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	} else {
		log.Printf("Listening to port %v", *port)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterContentServer(grpcServer, &contentServer{})
	grpcServer.Serve(lis)
}
