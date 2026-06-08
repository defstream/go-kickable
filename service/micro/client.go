package micro

import (
	"context"

	proto "github.com/defstream/go-kickable/service/micro/proto"
	micro "go-micro.dev/v4"
)

type Client struct {
	service micro.Service
	client  proto.KickableClient
}

func NewClient() *Client {
	svc := micro.NewService(micro.Name("kickable.client"))
	c := proto.NewKickableClient("Kickable", svc.Client())

	return &Client{
		service: svc,
		client:  c,
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
