package savegmailhistory

import (
	"context"
	"testing"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yamamoto-tgz/autosave/modules/pubsubdata"
)

func TestSaveGmailHistory(t *testing.T) {
	p := pubsubdata.New(`{"historyId":12345,"emailAddress":"autosave@yamamoto.tgz"}`)

	e := event.New("1.0")
	e.SetData("application/json", p)

	err := saveGmailHistory(context.Background(), e)
	if err != nil {
		t.Error(err)
	}
}
