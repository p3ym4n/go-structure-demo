package pubsubtest

import (
	"context"
	"go-structure-demo/internal/pubsub"
)

var _ pubsub.Client = (*Client)(nil)

type Client struct {
	PublishErr error
	ConsumeErr error
}

func NewMock() pubsub.Client {
	return new(Client)
}

func (c *Client) PublishMessage(ctx context.Context, topicID string, message []byte) (id string, err error) {
	return "mockEventID", c.PublishErr
}

func (c *Client) Consume(ctx context.Context, subscriptionID string, fn pubsub.MessageHandler) {
}
