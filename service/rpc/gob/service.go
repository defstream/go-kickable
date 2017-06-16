package gob

import (
	"net"
	"net/http"
	"net/rpc"

	"github.com/defstream/go-kickable"
)

type Kickable struct {
}

//This procedure is invoked by rpc and calls rpcexample.Multiply which stores product of args.A and args.B in result pointer
func (s *Kickable) CanIKick(it string, result *string) error {
	*result = kickable.CanIKick(it)
	return nil
}

type Service struct {
	address  string
	kickable *Kickable
}

func NewService(address string) *Service {
	return &Service{
		address:  address,
		kickable: &Kickable{},
	}
}

func (s *Service) Run() error {

	server := rpc.NewServer()

	err := server.RegisterName("Kickable", s.kickable)
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	return http.Serve(l, nil)
}
