package mokoshsvc_test

import (
	"testing"
	"github.com/mokosh/mokosh/pkg/mokoshsvc"
	"github.com/mokosh/mokosh/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func TestMokosh_Capabilities(t *testing.T) {
	capabilities, err := mokoshsvc.NewMokoshServer().Capabilities(nil, &pb.CapabilitiesRequest{})
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(capabilities.Capabilities) != 0 {
		t.Logf("we do not any empty capabilites, but we got some: %v", capabilities.Capabilities)
		t.Fail()
	}
}

func TestMokosh_Merge(t *testing.T) {
	err := mokoshsvc.NewMokoshServer().Merge(nil)
	t.Logf("received error: %v", err)
	if err == nil {
		t.Log("we expect an unconfigured Server to return nil.")
		t.Fail()
	}
	if grpc.Code(err) != codes.Unimplemented {
		t.Logf("we expect error code of %v - unimplemented", codes.Unimplemented)
		t.Fail()
	}
}
