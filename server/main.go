package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/nyaruka/courier/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedFlowsMesssageServer
}

func (s *server) CreateMessage(ctx context.Context, in *pb.CreateFlowsMessageRequest) (*pb.CreateFlowsMessageResponse, error) {
	return &pb.CreateFlowsMessageResponse{Message: "ok"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFlowsMesssageServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
