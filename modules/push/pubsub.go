package push

import (
	"cloud.google.com/go/pubsub"
	"github.com/cloudevents/sdk-go/v2/event"
)

func ExtractData(e event.Event) []byte {
	type Pubsub struct {
		Message pubsub.Message `json:"message"`
	}

	var p Pubsub
	e.DataAs(&p)

	return p.Message.Data
}
