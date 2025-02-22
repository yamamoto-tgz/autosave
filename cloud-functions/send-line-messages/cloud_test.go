package saveexpenses

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yamamoto-tgz/autosave/modules/line"
	"github.com/yamamoto-tgz/autosave/modules/pubsubdata"
)

func TestSaveGmailHistory(t *testing.T) {
	messages := []line.TextMessage{
		line.NewTextMessage("HELLO"),
		line.NewTextMessage("WORLD"),
	}
	j, _ := json.Marshal(messages)
	p := pubsubdata.New(string(j))
	e := event.New("1.0")
	e.SetData("application/json", p)

	err := sendLineMessages(context.Background(), e)
	if err != nil {
		t.Error(err)
	}
}
