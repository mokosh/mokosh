package mokoshsvc

import (
	"log"
	"golang.org/x/net/context"

	"github.com/mokosh/mokosh/pb"
	//"io"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type MergerConfig struct {
	ServerUrl string
}

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

func (s *mokosh) setMerger (merger *MergerConfig) {

	s.addCapability(pb.Capability_MERGE)
}

func NewMokoshServer(merger *MergerConfig) pb.MokoshServer {
	m := mokosh{}
	if merger != nil {
		m.setMerger(merger)
	}
	return &m
}


func (s *mokosh) addCapability (capability pb.Capability) {
	// Todo: check for duplicates
	n := len(s.capabilities)
	if n == cap(s.capabilities) {
		newSlice := make([]pb.Capability, len(s.capabilities), 2*len(s.capabilities)+1)
		copy(newSlice, s.capabilities)
		s.capabilities = newSlice
	}
	s.capabilities = s.capabilities[0 : n+1]
	s.capabilities[n] = capability
}
