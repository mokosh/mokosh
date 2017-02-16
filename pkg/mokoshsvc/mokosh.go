package mokoshsvc

import (
	"log"
	"golang.org/x/net/context"

	"github.com/mokosh/mokosh/pb"
	//"io"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type mokosh struct {
	capabilities []pb.Capability
}

func (s *mokosh) Capabilities(ctx context.Context, req *pb.CapabilitiesRequest) (*pb.CapabilitiesReply, error) {
	log.Print("Received mokosh.Capabilities request")
	return &pb.CapabilitiesReply{Capabilities: s.capabilities}, nil
}

func (s *mokosh) Merge(stream pb.Mokosh_MergeServer) error {
	return grpc.Errorf(codes.Unimplemented, "We do not yet support merging.")
	/*
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MergeResult{})
		}
	}*/
}

func NewMokoshServer() pb.MokoshServer {
	m := mokosh{}
	return &m
}
