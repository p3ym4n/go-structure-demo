package pubsub

import (
	"context"
	"time"

	gcloudpubsub "cloud.google.com/go/pubsub"
	"go-structure-demo/internal/log"
	"google.golang.org/api/option"
	pubsubtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/cloud.google.com/go/pubsub.v1"
)

var _ Client = (*GCPClient)(nil)

// MessageHandler is an alias for a function that handles an incoming message
type MessageHandler func(context.Context, []byte) (bool, error)

type Client interface {
	PublishMessage(ctx context.Context, topicID string, message []byte) (id string, err error)
	Consume(ctx context.Context, subID string, fn MessageHandler)
}

type GCPClient struct {
	logger         log.Logger
	gcpClient      *gcloudpubsub.Client
	projectID      string
	tracingEnabled bool
}

func New(logger log.Logger, ctx context.Context, projectID string, opts ...option.ClientOption) (*GCPClient, error) {
	client, err := gcloudpubsub.NewClient(ctx, projectID, opts...)

	go func(ctx context.Context) {
		defer func() { _ = client.Close() }()
		<-ctx.Done()
	}(ctx)

	return &GCPClient{logger: logger, gcpClient: client, projectID: projectID, tracingEnabled: true}, err
}

func (c *GCPClient) EnableTracing(enable bool) {
	c.tracingEnabled = enable
}

func (c *GCPClient) PublishMessage(ctx context.Context, topicID string, message []byte) (string, error) {
	topic := c.gcpClient.Topic(topicID)

	if c.tracingEnabled {
		id, err := pubsubtrace.Publish(ctx, topic, &gcloudpubsub.Message{Data: message}).Get(ctx)
		return id, err
	}

	id, err := topic.Publish(ctx, &gcloudpubsub.Message{Data: message}).Get(ctx)
	return id, err
}

func (c *GCPClient) Consume(ctx context.Context, subscriptionID string, fn MessageHandler) {
	handler := func(ctx context.Context, msg *gcloudpubsub.Message) {
		defer func() {
			if err := recover(); err != nil {
				c.logger.ErrorWithContext(ctx, "panic recovered", err)
				msg.Nack()
			}
		}()
		start := time.Now()
		ack, err := fn(ctx, msg.Data)
		if err != nil {
			c.logger.ErrorWithContext(ctx, "pubsub consumer handler error", map[string]interface{}{
				"elapsed":         time.Since(start).Milliseconds(),
				"error":           err.Error(),
				"project_id":      c.projectID,
				"subscription_id": subscriptionID,
				"ack":             ack,
			})
		}
		if ack {
			msg.Ack()
		} else {
			msg.Nack()
		}
	}

	sub := c.gcpClient.Subscription(subscriptionID)
	if c.tracingEnabled {
		handler = pubsubtrace.WrapReceiveHandler(sub, handler)
	}
	err := sub.Receive(ctx, handler)
	if err != nil {
		c.logger.ErrorWithContext(ctx, "pubsub consumer receiving error", map[string]interface{}{
			"project_id":      c.projectID,
			"subscription_id": subscriptionID,
			"error":           err.Error(),
		})
	}
}
