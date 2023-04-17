package service

import (
	"context"

	"github.com/luandnh/go-notify/repository/model"
)

type (
	Client struct {
	}
)

func NewClient() *Client {
	return &Client{}
}

func (c *Client) GetClientByClientToken(ctx context.Context, clientToken string) (*model.Client, error) {
	return nil, nil
}
