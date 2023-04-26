package alist

import (
	"net/url"
)

type MakeClientOptions struct {
	Endpoint string `json:"endpoint"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
}

type Client struct {
	MakeClientOptions
}

func NewClient(opt MakeClientOptions) Client {
	return Client{
		MakeClientOptions: opt,
	}
}

func (c *Client) ApiUrl(path string) (result string) {
	result, _ = url.JoinPath(c.Endpoint, "api", path)
	return
}
