package micro

import (
	proto "github.com/defstream/go-kickable/service/micro/proto"

	gomicro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Client struct {
	service gomicro.Service
	client  proto.KickableClient
}

func NewClient() *Client {
	service := gomicro.NewService(gomicro.Name("kicakble.client"))
	client := proto.NewKickableClient("Kickable", service.Client())

	return &Client{
		service: service,
		client:  client,
	}
}

func NewCanIKickRequest(it string) *proto.CanIKickRequest {
	return &proto.CanIKickRequest{
		It: it,
	}
}

func (c *Client) CanIKick(ctx context.Context, it string) (string, error) {
	req := NewCanIKickRequest(it)
	res, err := c.client.CanIKick(ctx, req)
	if err != nil {
		return "", err
	}

	return res.Response, nil
}
