package graphql

import (
	"fmt"
	"net/http"

	gql "github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
)

var schema *gql.Schema

type Service struct {
	address string
	schema  *gql.Schema
}

func NewService(address string) (*Service, error) {
	schema, err := gql.ParseSchema(Schema, &Resolver{})
	if err != nil {
		return nil, err
	}
	return &Service{
		address: address,
		schema:  schema,
	}, nil
}

func (s *Service) Run() error {
	http.Handle("/", &relay.Handler{Schema: s.schema})
	fmt.Printf("graphql server listening on %s\n", s.address)
	return http.ListenAndServe(s.address, nil)
}
