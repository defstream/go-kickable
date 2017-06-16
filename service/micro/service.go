package micro

import (
	"fmt"

	"github.com/defstream/go-kickable"
	proto "github.com/defstream/go-kickable/service/micro/proto"

	gomicro "github.com/micro/go-micro"

	"golang.org/x/net/context"
)

type Service struct {
	service gomicro.Service
}

func NewService() *Service {
	// Define the Micro Service
	service := gomicro.NewService(
		gomicro.Name("Kickable"),
		gomicro.Version("latest"),
		gomicro.Metadata(map[string]string{
			"type": "Kickable",
		}),
	)

	// Initialize the servie
	service.Init()

	s := &Service{
		service: service,
	}
	// Register handler
	proto.RegisterKickableHandler(service.Server(), s)

	// Return the micro service
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
