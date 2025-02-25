package receivegmailhistory

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yamamoto-tgz/autosave/packages/push"
)

func TestReceiveGmailHistory(t *testing.T) {
	h, _ := json.Marshal(History{
		EmailAddress: "autosave@yamamoto.tgz",
		HistoryId:    999999,
	})

	jsn, _ := json.Marshal(push.New([]byte(h)))

	e := event.New("1.0")
	e.SetData("application/json", jsn)

	err := receiveGmailHistory(context.Background(), e)
	if err != nil {
		t.Error(err)
	}
}
