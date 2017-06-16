package kite

import (
	"net"
	"strconv"

	"github.com/defstream/go-kickable"

	"github.com/koding/kite"
)

type Service struct {
	address string
}

func NewService(address string) *Service {
	return &Service{
		address: address,
	}
}

func (s *Service) Run() error {
	// Create a kite.
	k := kite.New("Kickable", "1.0.0")
	k.HandleFunc("CanIKick", func(r *kite.Request) (interface{}, error) {

		it := r.Args.One().MustString()
		result := kickable.CanIKick(it)

		return result, nil
	}).DisableAuthentication()
	h, p, err := net.SplitHostPort(s.address)
	if err != nil {
		return err
	}

	port, err := strconv.Atoi(p)
	if err != nil {
		return err
	}
	k.Config.IP = h
	k.Config.Port = port

	k.Run()

	return nil
}
