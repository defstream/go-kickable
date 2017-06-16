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

	// Connect to our math kite
	client := k.NewClient(address)

	client.Dial()

	return &Client{
		client: client,
	}
}

func (c *Client) CanIKick(ctx context.Context, it string) (string, error) {
	r, err := c.client.Tell("CanIKick", it) // call "square" method with argument 4
	if err != nil {
		return "", err
	}
	return r.MustString(), nil
}
