package http

import (
	"bytes"
	"context"
	"io"
	nethttp "net/http"
	"time"
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
	var client = &nethttp.Client{
		Timeout: time.Second * 10,
	}

	req, err := nethttp.NewRequest("GET", c.address, nethttp.NoBody)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	buffer := bytes.NewBuffer([]byte{})
	io.Copy(buffer, res.Body)

	return string(buffer.Bytes()), nil
}
