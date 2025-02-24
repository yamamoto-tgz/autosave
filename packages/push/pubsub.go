package push

import (
	"cloud.google.com/go/pubsub"
	"github.com/cloudevents/sdk-go/v2/event"
)

type Pubsub struct {
	Message pubsub.Message `json:"message"`
}

func New(data []byte) Pubsub {
	return Pubsub{
		Message: pubsub.Message{
			Data: data,
		},
	}
}

func ExtractData(e event.Event) []byte {

	var p Pubsub
	e.DataAs(&p)

	return p.Message.Data
}
