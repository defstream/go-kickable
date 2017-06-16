package grpc

import (
	proto "github.com/defstream/go-kickable/service/grpc/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client struct {
	connection *grpc.ClientConn
	client     proto.KickableClient
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := proto.NewKickableClient(conn)

	return &Client{
		connection: conn,
		client:     client,
	}, nil
}

func NewKickItRequest(it string) *proto.CanIKickRequest {
	return &proto.CanIKickRequest{
		It: it,
	}
}

func (c *Client) KickIt(ctx context.Context, it string) (string, error) {
	req := NewKickItRequest(it)

	res, err := c.client.CanIKick(ctx, req)
	if err != nil {
		return "", err
	}

	return res.Response, nil
}

func (c *Client) Close() error {
	return c.connection.Close()
}
