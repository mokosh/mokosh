package main

import (
	"log"

	"fmt"
	"github.com/mokosh/mokosh/pb"
	"github.com/mokosh/mokosh/pkg/mokoshsvc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var Version = "UNKNOWN"
var BuildTime = "-"

const (
	port = ":50051"
)

func main() {

	fmt.Printf("mokoshsrv %v - %v\n", Version, BuildTime)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := mokoshsvc.NewMokoshService(nil)
	pb.RegisterMokoshServer(s, server)
	pb.RegisterMergerServer(s, server)


	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
