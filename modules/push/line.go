package push

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type LinePusher struct {
	ProjectId string
	TopicId   string
}

func (l *LinePusher) SendLineMessage(ctx context.Context, txt string) error {
	cl, err := pubsub.NewClient(ctx, l.ProjectId)
	if err != nil {
		return err
	}
	defer cl.Close()

	cl.Topic(l.TopicId).Publish(ctx, &pubsub.Message{
		Data: []byte(txt),
	})

	return nil
}
