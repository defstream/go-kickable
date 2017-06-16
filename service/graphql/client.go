package graphql

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"text/template"
	"time"
)

type canIKickRequest struct {
	It string
}

type Client struct {
	address string
}

func (c *Client) CanIKick(ctx context.Context, it string) (string, error) {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	body, err := buildRequest(it)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("GET", c.address, body)
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

func NewClient(address string) *Client {
	return &Client{
		address: address,
	}
}

const requestTemplate = ""

func buildRequest(it string) (*bytes.Buffer, error) {
	t := template.New("top")
	b := bytes.NewBuffer([]byte{})
	template.Must(t.Parse(requestTemplate))
	err := t.Execute(b, canIKickRequest{It: it})
	if err != nil {
		return nil, err
	}
	return b, nil
}
