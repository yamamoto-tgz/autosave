package push

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type LinePusher struct {
	ProjectId string
	TopicId   string
}

func (l *LinePusher) SendLineMessage(ctx context.Context, txt []byte) error {
	cl, err := pubsub.NewClient(ctx, l.ProjectId)
	if err != nil {
		return err
	}
	defer cl.Close()

	_, err = cl.Topic(l.TopicId).Publish(ctx, &pubsub.Message{Data: txt}).Get(ctx)
	if err != nil {
		return err
	}

	return nil
}

func NewDefaultLinePusher() LinePusher {
	return LinePusher{
		ProjectId: "autosave-tgz",
		TopicId:   "send-line-message",
	}
}
