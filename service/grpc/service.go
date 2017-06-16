package grpc

import (
	"fmt"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	proto "github.com/defstream/go-kickable/service/grpc/proto"

	"github.com/defstream/go-kickable"
)

type Service struct {
	address  string
	certFile string
	keyFile  string
}

func NewService(address string, certFile, keyFile string) *Service {
	return &Service{
		address:  address,
		certFile: certFile,
		keyFile:  keyFile,
	}
}

func (g *Service) CanIKick(ctx context.Context, req *proto.CanIKickRequest) (*proto.CanIKickResponse, error) {
	it := kickable.CanIKick(req.It)

	res := &proto.CanIKickResponse{
		Response: it,
	}
	return res, nil
}

func (s *Service) Run() error {

	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	var opts []grpc.ServerOption
	if s.certFile != "" || s.keyFile != "" {
		creds, err := credentials.NewServerTLSFromFile(s.certFile, s.keyFile)
		if err != nil {
			return err
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	server := grpc.NewServer(opts...)

	proto.RegisterKickableServer(server, s)
	fmt.Printf("gRPC server listening on %s\n", s.address)
	return server.Serve(lis)
}
