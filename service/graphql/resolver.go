package graphql

import (
	"github.com/defstream/go-kickable"
)

type Resolver struct{}

func (r *Resolver) CanIKick(args struct{ It *string }) *string {
	k := kickable.CanIKick(*args.It)

	return &k
}
