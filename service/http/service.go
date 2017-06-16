package http

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/defstream/go-kickable"
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
	http.HandleFunc("/", handler)
	fmt.Printf("http server listening on %s\n", s.address)
	return http.ListenAndServe(s.address, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	it := r.URL.Path[1:]
	res := kickable.CanIKick(it)

	w.Write(bytes.NewBufferString(res).Bytes())
}
