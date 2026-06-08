package micro

import (
	"context"
	"fmt"

	kickable "github.com/defstream/go-kickable"
	proto "github.com/defstream/go-kickable/service/micro/proto"
	micro "go-micro.dev/v4"
)

type Service struct {
	service micro.Service
}

func NewService() *Service {
	svc := micro.NewService(
		micro.Name("Kickable"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "Kickable",
		}),
	)

	svc.Init()

	s := &Service{service: svc}

	if err := proto.RegisterKickableHandler(svc.Server(), s); err != nil {
		panic(err)
	}

	return s
}

func (g *Service) CanIKick(ctx context.Context, req *proto.CanIKickRequest, rsp *proto.CanIKickResponse) error {
	rsp.Response = kickable.CanIKick(req.It)
	return nil
}

func (g *Service) Run() error {
	fmt.Printf("micro server listening\n")
	return g.service.Run()
}
