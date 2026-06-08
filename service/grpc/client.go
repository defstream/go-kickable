package grpc

import (
	"context"

	proto "github.com/defstream/go-kickable/service/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	connection *grpc.ClientConn
	client     proto.KickableClient
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
