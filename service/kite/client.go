
package kite

import (
	"context"

	"github.com/koding/kite"
)

type Client struct {
	client *kite.Client
}

func NewClient(address string) *Client {

	k := kite.New("Kickable", "1.0.0")

	client := k.NewClient(address)

	client.Dial()

	return &Client{
		client: client,
	}
}

func (c *Client) CanIKick(ctx context.Context, it string) (string, error) {
	r, err := c.client.Tell("CanIKick", it)
	if err != nil {
		return "", err
	}
	return r.MustString(), nil
}
