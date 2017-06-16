package json

import (
	"context"
	"net"
	"net/rpc/jsonrpc"
)

type Client struct {
	address string
}

func NewClient(address string) *Client {
	return &Client{
		address: address,
	}
}

func (c *Client) CanIKick(ctx context.Context, it string) (string, error) {
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)

	var result string

	err = client.Call("Kickable.CanIKick", it, &result)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
